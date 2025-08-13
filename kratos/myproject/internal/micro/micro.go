package micro

import (
	"myproject/internal/micro/client"
	"myproject/internal/micro/registry"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(registry.ProviderSet, client.ProviderSet)
