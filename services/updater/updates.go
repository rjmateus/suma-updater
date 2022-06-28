package updater

import (
	"encoding/xml"
	"errors"
	"fmt"
	"os/exec"
)

const (
	ListUpdates = "lu"
	ListPatches = "lp"
)

func getXmlUpdates(cmd string) []byte {
	command := exec.Command("zypper", "--xmlout", cmd)

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

func GetAvailableUpdates() (ZypperUpdatesResult[Update], error) {
	xmlBytes := getXmlUpdates(ListUpdates)
	var outProcessed ZypperUpdatesResult[Update]
	xml.Unmarshal(xmlBytes, &outProcessed)
	return outProcessed, nil
}

func GetAvailablePatches() (ZypperUpdatesResult[Patch], error) {

	xmlBytes := getXmlUpdates(ListPatches)
	var outProcessed ZypperUpdatesResult[Patch]
	xml.Unmarshal(xmlBytes, &outProcessed)
	return outProcessed, nil
}
