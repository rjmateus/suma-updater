package download

import (
	"errors"
	"fmt"
	"github.com/rjmateus/suma-updater/util"
)

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
