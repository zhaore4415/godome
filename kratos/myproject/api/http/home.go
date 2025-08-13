package http

import (
	"net/http"

	"myproject/core"

	"github.com/gin-gonic/gin"
)

type HomeController struct {
	greeterSrv *core.GreeterService
}

func NewHomeController(r *gin.Engine, greeterSrv *core.GreeterService) *HomeController {
	home := &HomeController{greeterSrv: greeterSrv}
	r.GET("/", home.index)

	return home
}

func (c *HomeController) index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Welcome Kratos Server")
}
