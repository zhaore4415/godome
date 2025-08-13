package http

import (
	"context"

	"myproject/internal/micro/server"
	"myproject/internal/proto/hello"

	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewHTTPServer,
	NewHttpRouter,

	NewHomeController,
)

func NewHttpRouter() *gin.Engine {
	r := gin.Default()
	r.Use(kgin.Middlewares(recovery.Recovery(), customMiddleware))
	return r
}

func NewHTTPServer(logger log.Logger, bs *hello.Bootstrap, router *gin.Engine,
	_ *HomeController,
	// todo inject other controller
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(server.Middlewares(logger)...),
	}
	if bs.Server.Http == nil || len(bs.Server.Http.Addr) == 0 {
		return nil
	}
	if bs.Server.Http.Network != "" {
		opts = append(opts, http.Network(bs.Server.Http.Network))
	}
	if bs.Server.Http.Addr != "" {
		opts = append(opts, http.Address(bs.Server.Http.Addr))
	}
	if bs.Server.Http.Timeout != nil {
		opts = append(opts, http.Timeout(bs.Server.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	srv.HandlePrefix("/", router)

	return srv
}

func customMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		if /*tr*/ _, ok := transport.FromServerContext(ctx); ok {
			//fmt.Println("operation:", tr.Operation())
		}
		reply, err = handler(ctx, req)
		return
	}
}
