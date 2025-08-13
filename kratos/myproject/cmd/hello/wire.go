//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"myproject/api"
	"myproject/core"
	"myproject/data"
	"myproject/internal/micro"
	"myproject/internal/proto/hello"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(log.Logger, *hello.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(
		micro.ProviderSet,
		api.ProviderSet,
		core.ProviderSet,
		data.ProviderSet,
		newApp,
	))
}
