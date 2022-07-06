package download

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rjmateus/suma-updater/config"
	"github.com/rjmateus/suma-updater/repositories/download"
	"net/http"
	"os"
	"path"
	"strings"
)

func PathExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		// path/to/whatever exists
		return true
	} else {
		return false
	}

}

func HandleRepodata() gin.HandlerFunc {
	return func(c *gin.Context) {
		channel := c.Param("channel")
		fileName := c.Param("file")

		filePath := fmt.Sprintf("/var/cache/rhn/repodata/%s/%s", channel, fileName)
		if !PathExists(filePath) {
			if strings.HasSuffix(fileName, ".asc") || strings.HasSuffix(fileName, ".key") {
				c.String(http.StatusNotFound, fmt.Sprintf("Key or signature file not provided: %s", fileName))
			} else {
				c.String(http.StatusNotFound, fmt.Sprintf("File not found:%s", c.Request.URL.Path))
			}
		}
		downloadProcessor(c, filePath)
	}
}

const mountPoint = "/var/spacewalk/"

func HandlePackage(app *config.Application) gin.HandlerFunc {

	return func(c *gin.Context) {
		channel := c.Param("channel")
		pkinfo := parsePackageFileName(c.Request.URL.Path)
		packageDb, error := download.GetDownloadPackage(app.DBGorm, channel, pkinfo.name, pkinfo.version, pkinfo.release, pkinfo.arch, pkinfo.checksum, pkinfo.epoch)
		if error != nil {
			c.String(http.StatusNotFound, fmt.Sprintf("%s not found", path.Base(c.Request.URL.Path)))
		}
		downloadProcessor(c, path.Join(mountPoint, packageDb.Path))
	}
}

func HandlerMediaFiles() gin.HandlerFunc {
	return func(c *gin.Context) {
		fileName := c.Param("file")
		/*channel := c.Param("channel")

		if fileName == "products" {

		} else {*/
		c.String(http.StatusNotFound, fmt.Sprintf("%s not found", fileName))
		//}
	}
}

func downloadProcessor(c *gin.Context, filePath string) {
	if !PathExists(filePath) {
		c.String(http.StatusNotFound, fmt.Sprintf("File not found:%s", c.Request.URL.Path))
	} else {
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", path.Base(filePath)))
		c.File(filePath)
	}
}

type pkgInfo struct {
	name     string
	version  string
	release  string
	epoch    string
	arch     string
	orgId    string
	checksum string
}

func parsePackageFileName(filepath string) pkgInfo {
	parts := strings.Split(filepath, "/")

	extension := path.Ext(filepath)
	basename := strings.TrimSuffix(path.Base(filepath), extension)
	arch := basename[strings.LastIndex(basename, ".")+1:]
	rest := basename[:strings.LastIndex(basename, ".")]

	//String arch = StringUtils.substringAfterLast(basename, ".");
	//String rest = StringUtils.substringBeforeLast(basename, ".");
	release := ""
	name := ""
	version := ""
	epoch := ""
	org := ""
	checksum := ""

	// Debian packages names need spacial handling
	if "deb" == extension || "udeb" == extension {
		/*name = StringUtils.substringBeforeLast(rest, "_");
		rest = StringUtils.substringAfterLast(rest, "_");
		PackageEvr pkgEv = PackageEvr.parseDebian(rest);
		epoch = pkgEv.getEpoch();
		version = pkgEv.getVersion();
		release = pkgEv.getRelease();*/
	} else {
		release = rest[strings.LastIndex(rest, "-")+1:]
		rest = rest[:strings.LastIndex(rest, "-")]
		version = rest[strings.LastIndex(rest, "-")+1:]
		name = rest[:strings.LastIndex(rest, "-")]
	}
	// path is getPackage/<org>/<checksum>/filename
	if len(parts) == 9 && parts[5] == "getPackage" {
		org = parts[6]
		checksum = parts[7]
	}
	return pkgInfo{
		name:     name,
		version:  version,
		release:  release,
		epoch:    epoch,
		arch:     arch,
		orgId:    org,
		checksum: checksum,
	}
}
