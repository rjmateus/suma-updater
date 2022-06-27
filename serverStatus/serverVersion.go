package serverStatus

import (
	"encoding/xml"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

const PatternName = "patterns-uyuni_server"

func getPatternXmlInstalled() []byte {
	command := exec.Command("zypper", "--xmlout", "search", "-v", "-i", PatternName)

	out, error := command.Output()
	if error != nil {
		var exerr *exec.ExitError
		if errors.As(error, &exerr) {
			fmt.Printf("exit code error: %d \n", exerr.ExitCode())
			if exerr.ExitCode() == 104 {
				fmt.Println("patterns-suma_server not dound ")
			}
		}
	}
	return out
}

func parseInstalledXml(out []byte) ZypperSearchResult {
	var outProcessed ZypperSearchResult
	xml.Unmarshal(out, &outProcessed)
	return outProcessed
}

func GetServerStatus() (ServerStatus, error) {

	xmlSearchBytes := getPatternXmlInstalled()
	searchResultIntalled := parseInstalledXml(xmlSearchBytes)
	if len(searchResultIntalled.Solvable) == 0 {
		return ServerStatus{}, errors.New("unable to detect server version")
	}
	version := searchResultIntalled.Solvable[0].Edition
	vSplit := strings.Split(version, "-")
	return ServerStatus{vSplit[0], vSplit[1], searchResultIntalled.Solvable[0].Arch}, nil
}
