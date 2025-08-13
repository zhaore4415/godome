package client

import (
	"context"

	"myproject/internal/proto/hello"
	"bsi/kratos/micro/client"

	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewClient)

type Client struct {
	GreeterClient hello.GreeterClient
}

func NewClient(discovery registry.Discovery) (*Client, error) {
	greeterClient, err := newGreeterClient(discovery)
	if err != nil {
		return nil, err
	}

	return &Client{
		GreeterClient: greeterClient,
	}, nil
}

func newGreeterClient(discovery registry.Discovery) (hello.GreeterClient, error) {
	const srv = "bsi.hello"
	conn, err := client.NewGrpcConn(context.Background(), srv, discovery)
	if err != nil {
		return nil, err
	}
	return hello.NewGreeterClient(conn), nil
}
