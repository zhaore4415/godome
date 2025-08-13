package config

import (
	registryInternal "myproject/internal/micro/registry"
	"myproject/internal/proto/hello"

	consulConfig "github.com/go-kratos/kratos/contrib/config/consul/v2"
	"github.com/go-kratos/kratos/v2/config"
	"gopkg.in/yaml.v3"
)

func LoadConfig(path string, registryDC, registryAddress, registryToken string) (*hello.Bootstrap, error) {
	rc := &hello.Registry{Consul: &hello.Registry_Consul{
		Dc:    registryDC,
		Addr:  registryAddress,
		Token: registryToken,
	}}
	cc, err := registryInternal.NewConsulClient(&hello.Bootstrap{Registry: rc})
	if err != nil {
		return nil, err
	}
	cs, err := consulConfig.New(cc, consulConfig.WithPath(path))
	if err != nil {
		return nil, err
	}
	c := config.New(config.WithSource(cs), config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
		return yaml.Unmarshal(kv.Value, v)
	}))
	defer c.Close()

	if err := c.Load(); err != nil {
		return nil, err
	}

	var bs hello.Bootstrap
	if err := c.Scan(&bs); err != nil {
		return nil, err
	} else {
		bs.Registry = rc
	}

	return &bs, nil
}
