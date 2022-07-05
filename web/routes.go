package web

import (
	"github.com/gin-gonic/gin"
	"github.com/rjmateus/suma-updater/config"
	"github.com/rjmateus/suma-updater/web/api"
	"github.com/rjmateus/suma-updater/web/download"
	"net/http"
)

func initLocal(engine *gin.Engine) {
	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong\n")
	})
}

func InitRoutes(app *config.Application) {
	initLocal(app.Engine)
	api.InitRoutes(app.Engine.Group("/api"))
	download.InitDownloadRoutes(app)
}
