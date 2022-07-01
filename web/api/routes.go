package api

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(engine *gin.RouterGroup) {

	engine.GET("/status", handleGetStatus)
	engine.GET("/serviceStatus", handleGetServiceStatus)
	engine.POST("/refresh", handleRefresh)

	engine.GET("/updates", handleGetUpdates)
	engine.GET("/patches", handleGetPatches)

	engine.POST("/patch", handleInstallPatches)
	engine.POST("/updatePackage", handleUpdatePkg)

	engine.POST("/reboot", handleReboot)

}
