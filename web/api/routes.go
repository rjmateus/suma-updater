package api

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(engine *gin.RouterGroup) {

	engine.GET("/status", handleGetStatus)
	engine.POST("/refresh", handleRefresh)

	engine.GET("/updates", handleGetUpdates)
	engine.POST("/update", handleInstallUpdates)

	engine.GET("/patches", handleGetPatches)
}
