package grpc

import (
	"myproject/internal/micro/server"
	"myproject/internal/proto/hello"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewGRPCServer,

	NewGreeterHandler,
	NewUserHandler,
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(logger log.Logger, bs *hello.Bootstrap,
	greeter *GreeterHandler,
	user *UserHandler,
	// todo inject other handler
) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.Middleware(server.Middlewares(logger)...),
	}
	if bs.Server.Grpc.Network != "" {
		opts = append(opts, grpc.Network(bs.Server.Grpc.Network))
	}
	if bs.Server.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(bs.Server.Grpc.Addr))
	}
	if bs.Server.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(bs.Server.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	hello.RegisterGreeterServer(srv, greeter)
	hello.RegisterUserServer(srv, user)
	// todo register other service...

	return srv
}
