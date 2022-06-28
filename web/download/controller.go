package download

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/rjmateus/suma-updater/util"
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

func handleGetRepodata(c *gin.Context) {
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

func downloadProcessor(c *gin.Context, filePath string) {
	if !PathExists(filePath) {
		c.String(http.StatusNotFound, fmt.Sprintf("File not found:%s", c.Request.URL.Path))
	} else {
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", path.Base(filePath)))
		c.File(filePath)
	}
}

func handleGetMediaFiles(c *gin.Context) {
	fileName := c.Param("file")
	/*channel := c.Param("channel")

	if fileName == "products" {

	} else {*/
	c.String(http.StatusNotFound, fmt.Sprintf("%s not found", fileName))
	//}
}

const mountPoint = "/var/spacewalk/"

func handleGetPackage(c *gin.Context) {

	channel := c.Param("channel")

	filePath := c.Request.URL.Path

	basename := path.Base(filePath)

	pkinfo := parsePackageFileName(filePath)
	fmt.Println(pkinfo)

	packageDb, error := getPackageFromDb(channel, pkinfo)

	if error != nil {
		c.String(http.StatusNotFound, fmt.Sprintf("%s not found", basename))
	}

	downloadProcessor(c, path.Join(mountPoint, packageDb.path))
}

type Package struct {
	id   string
	path string
}

func getPackageFromDb(channel string, pkg pkgInfo) (Package, error) {
	db := util.GetDBconnection("/etc/rhn/rhn.conf")
	//db := util.GetDBconnection("rhn.conf")
	defer db.Close()

	sql := `select p.id, p.path, pe.epoch as epoch
                from
                  rhnPackageArch pa,
                  rhnChannelPackage cp,
                  rhnPackage p,
                  rhnChecksum cs,
                  rhnPackageEVR pe,
                  rhnPackageName pn,
                  rhnChannel c
        where 1=1
            and c.label = $1
            and pn.name = $2
            and pe.version = $3
            and pe.release = $4
            and c.id = cp.channel_id
            and pa.label = $5
            and pn.id = p.name_id
            and p.id = cp.package_id
            and p.evr_id = pe.id
            and p.package_arch_id = pa.id
            and p.checksum_id = cs.id`

	parameter := []interface{}{channel, pkg.name, pkg.version, pkg.release, pkg.arch}
	if len(pkg.checksum) > 0 {
		parameter = append(parameter, pkg.checksum)
		sql = sql + " and cs.checksum = $6"
	} else {
		sql = sql + " and cs.checksum is null"
	}

	queryResult := util.ExecuteQueryWithResults(db, sql, parameter...)

	if len(queryResult) == 0 {
		return Package{}, errors.New("no package found")
	}

	result := make([]Package, 0)

	for _, row := range queryResult {
		mapData := convertRowToMap(row)
		if len(pkg.epoch) > 0 {
			if pkg.epoch != fmt.Sprintf("%s", mapData["epoc"]) {
				continue
			}
		}
		result = append(result, Package{
			id:   mapData["id"],
			path: mapData["path"],
		})
	}
	if len(result) == 0 {
		return Package{}, errors.New("no package found")
	}
	return result[0], nil
}

func convertRowToMap(data []util.RowDataStructure) map[string]string {

	mapData := make(map[string]string, 0)
	for _, r := range data {
		mapData[r.ColumnName] = fmt.Sprintf("%s", r.Value)
	}

	return mapData
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
