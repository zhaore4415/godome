package core

import (
	"myproject/core/cache"
	mgorm "bsi/kratos/micro/gorm"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	mgorm.NewTransactionScope,
	cache.NewRedisCache,

	NewGreeterService,
	NewUserService,
)
