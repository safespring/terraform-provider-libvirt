module github.com/dmacvicar/terraform-provider-libvirt

require (
	github.com/community-terraform-providers/terraform-provider-ignition/v2 v2.1.2
	github.com/davecgh/go-spew v1.1.1
	github.com/digitalocean/go-libvirt v0.0.0-20220616141158-7ed4ed4decd9
	github.com/google/uuid v1.1.2
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.21.0
	github.com/hooklift/iso9660 v1.0.0
	github.com/mattn/goveralls v0.0.2
	github.com/stretchr/testify v1.7.2
	golang.org/x/crypto v0.0.0-20220517005047-85d78b3ac167
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b
	libvirt.org/go/libvirtxml v1.8003.0
)

require (
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/apparentlymart/go-cidr v1.1.0 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/c4milo/gotoolkit v0.0.0-20170704181456-e37eeabad07e // indirect
	github.com/coreos/go-json v0.0.0-20200220154158-5ae607161559 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.0.0 // indirect
	github.com/coreos/ignition/v2 v2.3.0 // indirect
	github.com/coreos/vcontext v0.0.0-20200225161404-ee043618d38d // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-checkpoint v0.5.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-cty v1.4.1-0.20200414143053-d3edf31b6320 // indirect
	github.com/hashicorp/go-hclog v1.2.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-plugin v1.4.4 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/hc-install v0.4.0 // indirect
	github.com/hashicorp/hcl/v2 v2.13.0 // indirect
	github.com/hashicorp/logutils v1.0.0 // indirect
	github.com/hashicorp/terraform-exec v0.17.2 // indirect
	github.com/hashicorp/terraform-json v0.14.0 // indirect
	github.com/hashicorp/terraform-plugin-go v0.14.0 // indirect
	github.com/hashicorp/terraform-plugin-log v0.7.0 // indirect
	github.com/hashicorp/terraform-registry-address v0.0.0-20220623143253-7d51757b572c // indirect
	github.com/hashicorp/terraform-svchost v0.0.0-20200729002733-f050f53b9734 // indirect
	github.com/hashicorp/yamux v0.0.0-20190923154419-df201c70410d // indirect
	github.com/hooklift/assert v0.0.0-20170704181755-9d1defd6d214 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/vincent-petithory/dataurl v0.0.0-20191104211930-d1553a71de50 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	github.com/vmihailenco/msgpack/v4 v4.3.12 // indirect
	github.com/vmihailenco/tagparser v0.1.1 // indirect
	github.com/zclconf/go-cty v1.10.0 // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20220517195934-5e4e11fc645e // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.1 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/genproto v0.0.0-20200711021454-869866162049 // indirect
	google.golang.org/grpc v1.48.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999

replace golang.org/x/crypto v0.0.0-20220112180741-5e0467b6c7ce => github.com/dmacvicar/golang-x-crypto v0.0.0-20220126233154-a96af8f07497

go 1.18

replace github.com/community-terraform-providers/terraform-provider-ignition/v2 => github.com/dmacvicar/terraform-provider-ignition/v2 v2.1.3-0.20210701165004-13acf61ca184
