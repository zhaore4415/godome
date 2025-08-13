package server

import (
	_ "bsi/kratos/micro/codec"
	"bsi/kratos/micro/errors"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
)

func Middlewares(logger log.Logger) []middleware.Middleware {
	return []middleware.Middleware{
		recovery.Recovery(),
		logging.Server(logger),
		tracing.Server(),
		metadata.Server(),
		validate.Validator(),
		errors.Errors(),
	}
}
