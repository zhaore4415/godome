package api

import (
	"myproject/api/grpc"
	"myproject/api/http"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(grpc.ProviderSet, http.ProviderSet)
