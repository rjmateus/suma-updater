package api

import (
	"fmt"
	"github.com/rjmateus/suma-updater/services/serverVersion"
	"github.com/rjmateus/suma-updater/services/updater"
)

func getNewPatternVersion(updatesList updater.ZypperUpdatesResult[updater.Update], pkgName string) *updater.Update {

	fmt.Printf("available updates: %d\n\n", len(updatesList.Updates))
	for _, update := range updatesList.Updates {
		fmt.Printf("%s: %s -> %s\n", update.Name, update.EditionOld, update.Edition)
		if update.Name == pkgName {
			return &update
		}
	}
	return nil
}

func GetServerStatus() (ServerStatus, error) {
	serverVersion, error := serverVersion.GetServerStatus()
	if error != nil {
		fmt.Println("Error getting server status")
		return ServerStatus{}, error
	}
	updates, error := updater.GetAvailableUpdates()
	if error != nil {
		fmt.Println("Error getting updates")
		return ServerStatus{}, error
	}
	patches, error := updater.GetAvailablePatches()
	if error != nil {
		fmt.Println("Error getting patches")
		return ServerStatus{}, error
	}
	//newPkgVersion := getNewPatternVersion(updates, serverVersion.ControlPkgName)

	return ServerStatus{
		Version:    serverVersion.Version,
		Release:    serverVersion.Release,
		Arch:       serverVersion.Arch,
		NewVersion: serverVersion.NewVersion,
		NewRelease: serverVersion.NewRelease,
		Updates:    len(updates.Updates),
		Patches:    len(patches.Updates),
	}, nil
}
