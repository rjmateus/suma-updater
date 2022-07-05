package download

import (
	"github.com/rjmateus/suma-updater/config"
	"github.com/rjmateus/suma-updater/handlers/download"
)

func InitDownloadRoutes(app *config.Application) {

	router := app.Engine.Group("/rhn/manager/download")
	router.GET("/:channel/repodata/:file", download.GetHandlerRepodata(app))
	router.HEAD("/:channel/repodata/:file", download.GetHandlerRepodata(app))

	router.GET("/:channel/media.1/:file", download.GetHandlerMediaFiles(app))
	router.HEAD("/:channel/media.1/:file", download.GetHandlerMediaFiles(app))

	router.GET("/:channel/getPackage/:org/:checksum/:file", download.GetHandlePackage(app))
	router.HEAD("/:channel/getPackage/:org/:checksum/:file", download.GetHandlePackage(app))
	// :org doesn't represent org, in this constext it represents the file.
	// this is a known limitation of gin. See:
	// - https://github.com/gin-gonic/gin/issues/1301#issuecomment-392346179
	// - https://github.com/gin-gonic/gin/issues/1681
	router.GET("/:channel/getPackage/:org", download.GetHandlePackage(app))
	router.HEAD("/:channel/getPackage/:org", download.GetHandlePackage(app))
}
