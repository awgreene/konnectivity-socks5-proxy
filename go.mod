module github.com/awgreene/konnectivity-socks5-proxy

go 1.16

replace sigs.k8s.io/apiserver-network-proxy/konnectivity-client => sigs.k8s.io/apiserver-network-proxy/konnectivity-client v0.0.24

require (
	github.com/armon/go-socks5 v0.0.0-20160902184237-e75332964ef5
	github.com/spf13/cobra v1.1.3
	k8s.io/api v0.22.2
	k8s.io/apimachinery v0.22.2
	sigs.k8s.io/apiserver-network-proxy v0.0.24
	sigs.k8s.io/controller-runtime v0.10.1
)
