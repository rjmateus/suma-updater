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

func handleUpdatePkg(c *gin.Context) {
	var json struct {
		Packages []string `json:"packages" binding:"required"`
	}

	if c.Bind(&json) == nil {
		if len(json.Packages) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "unable to process body"})
		}
		result, error := updater.UpdatePackages(json.Packages)
		if error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": error})
		} else {
			c.JSON(http.StatusOK, result)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unable to process body"})
	}
}

func handleInstallPatches(c *gin.Context) {
	var json struct {
		withUpdate   bool `json:"withUpdate" `
		withOptional bool `json:"withOptional" `
	}

	if c.Bind(&json) == nil {
		result, error := updater.UpdatePatches(json.withUpdate, json.withOptional)
		if error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": error})
		} else {
			c.JSON(http.StatusOK, result)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unable to process body"})
	}

}

func handleGetPatches(c *gin.Context) {
	updates, _ := updater.GetAvailablePatches()
	c.JSON(http.StatusOK, updates.Updates)
}
