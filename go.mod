module github.com/harvester/harvester-installer

go 1.16

require (
	github.com/docker/go-units v0.4.0 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/imdario/mergo v0.3.12
	github.com/insomniacslk/dhcp v0.0.0-20210827173440-b95caade3eac
	github.com/jroimartin/gocui v0.4.0
	github.com/mattn/go-shellwords v1.0.10 // indirect
	github.com/mudler/yip v0.0.0-20211129144714-088f39125cf7
	github.com/nsf/termbox-go v1.1.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/rancher/mapper v0.0.0-20190814232720-058a8b7feb99
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
	github.com/tredoe/osutil v1.3.6
	github.com/vishvananda/netlink v1.1.0
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c
	gopkg.in/ini.v1 v1.63.2
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/apimachinery v0.18.8
)

replace (
	github.com/nsf/termbox-go => github.com/Harvester/termbox-go v1.1.1-0.20210318083914-8ab92204a400
	k8s.io/api => github.com/rancher/kubernetes/staging/src/k8s.io/api v1.19.3-k3s1
	k8s.io/apimachinery => github.com/rancher/kubernetes/staging/src/k8s.io/apimachinery v1.19.3-k3s1
	k8s.io/client-go => github.com/rancher/kubernetes/staging/src/k8s.io/client-go v1.19.3-k3s1
)
