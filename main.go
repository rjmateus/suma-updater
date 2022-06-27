package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rjmateus/suma-updater/commands"
	"github.com/rjmateus/suma-updater/serverStatus"
	"github.com/rjmateus/suma-updater/updater"
	"net/http"
)

func main() {
	setupRouter()

	// https://hoohoo.top/blog/20210530112304-golang-tutorial-introduction-gin-html-template-and-how-integration-with-bootstrap/

}

func setupRouter() *gin.Engine {
	r := gin.Default()
	apiGroup := r.Group("/api")
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong\n")
	})

	apiGroup.GET("/status", func(c *gin.Context) {
		status, _ := serverStatus.GetServerStatus()
		c.JSON(http.StatusOK, status)
	})

	apiGroup.GET("/serverUpdatable", func(c *gin.Context) {
		updates, _ := updater.GetAvailableUpdates()
		c.JSON(http.StatusOK, updater.IsServerUpdateAvailable(updates))
	})

	apiGroup.GET("/updates", func(c *gin.Context) {
		updates, _ := updater.GetAvailableUpdates()
		c.JSON(http.StatusOK, updates.Updates)
	})

	apiGroup.GET("/patches", func(c *gin.Context) {
		updates, _ := updater.GetAvailablepatches()
		c.JSON(http.StatusOK, updates.Updates)
	})

	apiGroup.POST("/refresh", func(c *gin.Context) {
		output := commands.ZypperRef()
		c.JSON(http.StatusOK, output)
	})

	apiGroup.POST("/update", func(c *gin.Context) {
		var json struct {
			Packages []string `json:"packages" binding:"required"`
		}

		if c.Bind(&json) == nil {
			c.JSON(http.StatusOK, gin.H{"value": json.Packages})
		} else {
			c.String(http.StatusOK, "no value")
		}
	})

	r.Run(":8088")
	return r
}
