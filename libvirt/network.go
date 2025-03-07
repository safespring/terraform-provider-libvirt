package libvirt

import (
	"fmt"
	"log"
	"net"
	"strings"

	libvirt "github.com/digitalocean/go-libvirt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"libvirt.org/go/libvirtxml"
)

func waitForNetworkActive(virConn *libvirt.Libvirt, network libvirt.Network) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		active, err := virConn.NetworkIsActive(network)
		if err != nil {
			return nil, "", err
		}
		if active == 1 {
			return network, "ACTIVE", nil
		}
		return network, "BUILD", err
	}
}

// waitForNetworkDestroyed waits for a network to destroyed
func waitForNetworkDestroyed(virConn *libvirt.Libvirt, uuidStr string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("Waiting for network %s to be destroyed", uuidStr)

		uuid := parseUUID(uuidStr)

		_, err := virConn.NetworkLookupByUUID(uuid)
		if err.(libvirt.Error).Code == uint32(libvirt.ErrNoNetwork) {
			return virConn, "NOT-EXISTS", nil
		}
		return virConn, "ACTIVE", err
	}
}

// getNetModeFromResource returns the network mode fromm a network definition
func getNetModeFromResource(d *schema.ResourceData) string {
	return strings.ToLower(d.Get("mode").(string))
}

// getIPsFromResource gets the IPs configurations from the resource definition
func getIPsFromResource(d *schema.ResourceData) ([]libvirtxml.NetworkIP, error) {
	addresses, ok := d.GetOk("addresses")
	if !ok {
		return []libvirtxml.NetworkIP{}, nil
	}

	// check if DHCP must be enabled by default
	var dhcpEnabled bool
	netMode := getNetModeFromResource(d)
	if netMode == netModeIsolated || netMode == netModeNat || netMode == netModeRoute || netMode == netModeOpen {
		dhcpEnabled = true
	}

	ipsPtrsLst := []libvirtxml.NetworkIP{}
	for _, addressI := range addresses.([]interface{}) {
		// get the IP address entry for this subnet (with a guessed DHCP range)
		dni, dhcp, err := getNetworkIPConfig(addressI.(string))
		if err != nil {
			return nil, err
		}

		// HACK traditionally the provider simulated an easy cloud, so
		// DHCP was enabled by default.
		//
		// However, there was some weird code added later that couples
		// network device and networks.
		//
		// This requires us to have the "enabled" property to be
		// computed, which prevents us from having it as Default: true.
		//
		// This code enables DHCP by default if the setting is not
		// explicitly given, to still allow for it to be Computed.
		//
		// The same reason we have to use deprecated GetOkExists
		// because it is computed but we need to know if the user has
		// explicitly set it to false
		//
		//nolint:staticcheck
		if dhcpEnabledByUser, dhcpSetByUser := d.GetOkExists("dhcp.0.enabled"); dhcpSetByUser {
			dhcpEnabled = dhcpEnabledByUser.(bool)
		} else {
			// if not specified, default to enable it
			dhcpEnabled = true
		}

		if dhcpEnabled {
			dni.DHCP = dhcp
		} else {
			// if a network exist with enabled but an user want to disable it
			// we need to set DHCP struct to nil.
			dni.DHCP = nil
		}

		ipsPtrsLst = append(ipsPtrsLst, *dni)
	}

	return ipsPtrsLst, nil
}

func getNetworkIPConfig(address string) (*libvirtxml.NetworkIP, *libvirtxml.NetworkDHCP, error) {
	_, ipNet, err := net.ParseCIDR(address)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing addresses definition '%s': %s", address, err)
	}
	ones, bits := ipNet.Mask.Size()
	family := "ipv4"
	if bits == (net.IPv6len * 8) {
		family = "ipv6"
	}
	ipsRange := (1 << bits) - (1 << ones)
	if ipsRange < 4 {
		return nil, nil, fmt.Errorf("netmask seems to be too strict: only %d IPs available (%s)", ipsRange-3, family)
	}

	// we should calculate the range served by DHCP. For example, for
	// 192.168.121.0/24 we will serve 192.168.121.2 - 192.168.121.254
	start, end := networkRange(ipNet)

	// skip the .0, (for the network),
	start[len(start)-1]++

	// assign the .1 to the host interface
	dni := &libvirtxml.NetworkIP{
		Address: start.String(),
		Prefix:  uint(ones),
		Family:  family,
	}

	start[len(start)-1]++ // then skip the .1
	end[len(end)-1]--     // and skip the .255 (for broadcast)

	dhcp := &libvirtxml.NetworkDHCP{
		Ranges: []libvirtxml.NetworkDHCPRange{
			{
				Start: start.String(),
				End:   end.String(),
			},
		},
	}

	return dni, dhcp, nil
}

// getBridgeFromResource returns a libvirt's NetworkBridge
// from the ResourceData provided.
func getBridgeFromResource(d *schema.ResourceData) *libvirtxml.NetworkBridge {
	// use a bridge provided by the user, or create one otherwise (libvirt will assign on automatically when empty)
	bridgeName := ""
	if b, ok := d.GetOk("bridge"); ok {
		bridgeName = b.(string)
	}

	bridge := &libvirtxml.NetworkBridge{
		Name: bridgeName,
		STP:  "on",
	}

	return bridge
}

// getDomainFromResource returns a libvirt's NetworkDomain
// from the ResourceData provided.
func getDomainFromResource(d *schema.ResourceData) *libvirtxml.NetworkDomain {
	domainName, ok := d.GetOk("domain")
	if !ok {
		return nil
	}

	domain := &libvirtxml.NetworkDomain{
		Name: domainName.(string),
	}

	if dnsLocalOnly, ok := d.GetOk(dnsPrefix + ".local_only"); ok {
		if dnsLocalOnly.(bool) {
			domain.LocalOnly = "yes" // this "boolean" must be "yes"|"no"
		}
	}

	return domain
}

func getMTUFromResource(d *schema.ResourceData) *libvirtxml.NetworkMTU {
	if mtu, ok := d.GetOk("mtu"); ok {
		return &libvirtxml.NetworkMTU{Size: uint(mtu.(int))}
	}

	return nil
}

// getDNSMasqOptionFromResource returns a list of dnsmasq options
// from the network definition
func getDNSMasqOptionFromResource(d *schema.ResourceData) ([]libvirtxml.NetworkDnsmasqOption, error) {
	var dnsmasqOption []libvirtxml.NetworkDnsmasqOption
	dnsmasqOptionPrefix := "dnsmasq_options.0"
	if dnsmasqOptionCount, ok := d.GetOk(dnsmasqOptionPrefix + ".options.#"); ok {
		for i := 0; i < dnsmasqOptionCount.(int); i++ {
			dnsmasqOptionsPrefix := fmt.Sprintf(dnsmasqOptionPrefix+".options.%d", i)

			optionName := d.Get(dnsmasqOptionsPrefix + ".option_name").(string)
			optionValue := d.Get(dnsmasqOptionsPrefix + ".option_value").(string)
			dnsmasqOption = append(dnsmasqOption, libvirtxml.NetworkDnsmasqOption{
				Value: optionName + "=" + optionValue,
			})
		}
	}

	return dnsmasqOption, nil
}
