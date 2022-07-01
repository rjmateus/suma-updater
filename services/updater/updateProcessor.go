package updater

import (
	"encoding/json"
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

	fmt.Println("start zypper process")
	command := exec.Command("zypper", cmd...)

	out, errorZypper := command.Output()
	fmt.Println("finish zypper process")

	errorsReturn := make([]any, 0)

	if errorZypper != nil {
		var exerr *exec.ExitError
		if errors.As(errorZypper, &exerr) {
			fmt.Printf("exit code error: %d \n", exerr.ExitCode())
			if exerr.ExitCode() >= 100 {
				fmt.Println("worning error returned", exerr)
			} else {
				errorsReturn = append(errorsReturn, errorZypper)
			}
		}
	}
	var outProcessed ZypperRunUpdateResult
	errorParse := xml.Unmarshal(out, &outProcessed)
	if errorParse != nil {
		errorsReturn = append(errorsReturn, errorParse)
	}

	_, errStart := startServives()
	if errStart != nil {
		errorsReturn = append(errorsReturn, errStart)
	}
	if len(errorsReturn) > 0 {
		errorOut, _ := json.Marshal(errorsReturn)
		return outProcessed, errors.New(string(errorOut))
	} else {
		return outProcessed, nil
	}
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
	fmt.Println("starting spacewalk-service")
	out, err := command.Output()
	fmt.Println("finish starting spacewalk-service")
	return out, err
}
