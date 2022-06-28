package web

import (
	"github.com/gin-gonic/gin"
	"github.com/rjmateus/suma-updater/web/api"
	"github.com/rjmateus/suma-updater/web/download"
	"net/http"
)

func initLocal(engine *gin.Engine) {
	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong\n")
	})
}

func InitRoutes(engine *gin.Engine) {
	initLocal(engine)
	api.InitRoutes(engine.Group("/api"))
	download.InitRoutes(engine.Group("/rhn/manager"))
}
