package updater

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/rjmateus/suma-updater/serverStatus"
	"os/exec"
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
	xmlBytes := getXmlUpdates("lu")
	var outProcessed ZypperUpdatesResult[Update]
	xml.Unmarshal(xmlBytes, &outProcessed)
	return outProcessed, nil
}

func GetAvailablepatches() (ZypperUpdatesResult[Patch], error) {

	xmlBytes := getXmlUpdates("lp")
	var outProcessed ZypperUpdatesResult[Patch]
	xml.Unmarshal(xmlBytes, &outProcessed)
	return outProcessed, nil
}

func IsServerUpdateAvailable(updates ZypperUpdatesResult[Update]) *Update {

	fmt.Printf("available updates: %d\n\n", len(updates.Updates))
	var serverUpdate *Update = nil
	for _, update := range updates.Updates {
		fmt.Printf("%s: %s -> %s\n", update.Name, update.EditionOld, update.Edition)

		if update.Name == serverStatus.PatternName {
			serverUpdate = &update
		}
	}
	if serverUpdate != nil {
		return serverUpdate
	}
	return nil
}
