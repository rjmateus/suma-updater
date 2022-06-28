package updater

import (
	"encoding/xml"
	"errors"
	"fmt"
	"os/exec"
)

func performUpdate(cmd []string) (ZypperRunUpdateResult, error) {

	stopStdout, err := stopServives()
	if err != nil {
		return ZypperRunUpdateResult{}, errors.New(string(stopStdout))
	}

	fmt.Printf("start zypper process")
	command := exec.Command("zypper", cmd...)

	out, error := command.Output()
	fmt.Printf("finish zypper process")
	if error != nil {
		return ZypperRunUpdateResult{}, error
	}

	var outProcessed ZypperRunUpdateResult
	xml.Unmarshal(out, &outProcessed)

	startStdout, err := startServives()
	if err != nil {
		return outProcessed, errors.New(string(startStdout))
	}
	return outProcessed, nil
}

func UpdatePackages(pkgs []string) (ZypperRunUpdateResult, error) {

	cmd := []string{"--xmlout", "up", "--skip-interactive", "--no-confirm"}
	cmd = append(cmd, pkgs...)
	return performUpdate(cmd)
}

func UpdatePatches(withUpdates bool, withOptional bool) (ZypperRunUpdateResult, error) {

	cmd := []string{"--xmlout", "patch", "--skip-interactive", "--no-confirm"}
	if withUpdates {
		cmd = append(cmd, "--with-update")
	}
	if withOptional {
		cmd = append(cmd, "--with-optional")
	}

	return performUpdate(cmd)
}

func stopServives() ([]byte, error) {
	command := exec.Command("spacewalk-service", "stop")
	fmt.Println("stopped spacewalk-service")
	return command.Output()
}
func startServives() ([]byte, error) {
	command := exec.Command("spacewalk-service", "start")
	fmt.Println("started spacewalk-service")
	return command.Output()
}
