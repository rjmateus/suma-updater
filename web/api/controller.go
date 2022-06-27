package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rjmateus/suma-updater/services/updater"
	"github.com/rjmateus/suma-updater/services/zypper"
	"net/http"
)

func handleGetStatus(c *gin.Context) {
	status, error := GetServerStatus()
	if error != nil {
		c.JSON(http.StatusOK, ApiErro{error.Error()})
	} else {
		c.JSON(http.StatusOK, status)
	}
}

func handleRefresh(c *gin.Context) {
	output, _ := zypper.ZypperRef()
	c.JSON(http.StatusOK, output)
}

func handleGetUpdates(c *gin.Context) {
	updates, _ := updater.GetAvailableUpdates()
	c.JSON(http.StatusOK, updates.Updates)
}

func handleInstallUpdates(c *gin.Context) {
	var json struct {
		Packages []string `json:"packages" binding:"required"`
	}

	if c.Bind(&json) == nil {
		c.JSON(http.StatusOK, gin.H{"value": json.Packages})
	} else {
		c.String(http.StatusOK, "no value")
	}
}

func handleGetPatches(c *gin.Context) {
	updates, _ := updater.GetAvailablePatches()
	c.JSON(http.StatusOK, updates.Updates)
}
