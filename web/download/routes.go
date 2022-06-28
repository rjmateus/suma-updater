package download

import "github.com/gin-gonic/gin"

func InitRoutes(engine *gin.RouterGroup) {
	engine.GET("/download/:channel/repodata/:file", handleGetRepodata)
	engine.HEAD("/download/:channel/repodata/:file", handleGetRepodata)

	engine.GET("/download/:channel/media.1/:file", handleGetMediaFiles)
	engine.HEAD("/download/:channel/media.1/:file", handleGetMediaFiles)

	engine.GET("/download/:channel/getPackage/:org/:checksum/:file", handleGetPackage)
	engine.HEAD("/download/:channel/getPackage/:org/:checksum/:file", handleGetPackage)
	// :org doesn't represent org, in this constext it represents the file.
	// this is a known limitation of gin. See:
	// - https://github.com/gin-gonic/gin/issues/1301#issuecomment-392346179
	// - https://github.com/gin-gonic/gin/issues/1681
	engine.GET("/download/:channel/getPackage/:org", handleGetPackage)
	engine.HEAD("/download/:channel/getPackage/:org", handleGetPackage)
}
