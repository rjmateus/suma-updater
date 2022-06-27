package serverVersion

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/rjmateus/suma-updater/services"
	"github.com/rjmateus/suma-updater/services/updater"
	"os/exec"
	"strings"
)

type zypperCmdOut struct {
	pkgName string
	data    []byte
	error   error
}

func (out zypperCmdOut) isPkgNotFoundError() bool {
	if out.error != nil {
		var exerr *exec.ExitError
		if errors.As(out.error, &exerr) {
			if exerr.ExitCode() == 104 {
				return true
			}
		}
	}
	return false
}

func getPatternXmlInstalled(packageName string) zypperCmdOut {
	command := exec.Command("zypper", "--xmlout", "search", "-v", "-i", packageName)
	out, error := command.Output()
	return zypperCmdOut{packageName, out, error}

}

func parseInstalledXml(out []byte) ZypperSearchResult {
	var outProcessed ZypperSearchResult
	xml.Unmarshal(out, &outProcessed)
	return outProcessed
}

func getNewPatternVersion(pkgName string) *updater.Update {
	updates, _ := updater.GetAvailableUpdates()
	fmt.Printf("available updates: %d\n\n", len(updates.Updates))
	for _, update := range updates.Updates {
		//fmt.Printf("%s: %s -> %s\n", update.Name, update.EditionOld, update.Edition)
		if update.Name == pkgName {
			return &update
		}
	}
	return nil
}

func GetServerStatus() (ServerStatus, error) {
	cmdOut := getPatternXmlInstalled(services.PatternSumaName)
	if cmdOut.error != nil {
		if cmdOut.isPkgNotFoundError() {
			cmdOut = getPatternXmlInstalled(services.PatternUyuniName)
			if cmdOut.error != nil {
				return ServerStatus{}, cmdOut.error
			}

		} else {
			return ServerStatus{}, cmdOut.error
		}
	}
	searchResultInstalled := parseInstalledXml(cmdOut.data)
	if len(searchResultInstalled.Solvable) == 0 {
		return ServerStatus{}, errors.New("unable to detect server version")
	}
	version := searchResultInstalled.Solvable[0].Edition
	vSplit := strings.Split(version, "-")

	newVersion := getNewPatternVersion(cmdOut.pkgName)
	vSplitNew := []string{"", ""}
	if newVersion != nil {
		vSplitNew = strings.Split(newVersion.Edition, "-")
	}

	return ServerStatus{
		cmdOut.pkgName,
		vSplit[0],
		vSplit[1],
		searchResultInstalled.Solvable[0].Arch,
		vSplitNew[0],
		vSplitNew[1]}, nil
}
