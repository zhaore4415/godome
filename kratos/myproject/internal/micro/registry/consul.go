package registry

import (
	"net/url"

	"myproject/internal/proto/hello"
	"bsi/kratos/micro/registry/consul"

	"github.com/go-kratos/kratos/v2/registry"
	"github.com/hashicorp/consul/api"
)

func NewConsulClient(bs *hello.Bootstrap) (*api.Client, error) {
	if bs.Registry == nil {
		panic("registry config is nil")
	}
	if bs.Registry.Consul == nil {
		panic("consul registry config is nil")
	}
	uri, err := url.Parse(bs.Registry.Consul.Addr)
	if err != nil {
		return nil, err
	}
	return api.NewClient(&api.Config{
		Address:    uri.Host,
		Scheme:     uri.Scheme,
		Token:      bs.Registry.Consul.Token,
		Datacenter: bs.Registry.Consul.Dc,
	})
}

func NewConsulRegistry(client *api.Client) registry.Registrar {
	return consul.New(client)
}

func NewConsulDiscovery(client *api.Client) registry.Discovery {
	return consul.New(client)
}
