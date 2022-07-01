package serverVersion

import (
	"fmt"
	"os/exec"
	"strings"
)

func getServicesName() ([]string, error) {
	command := exec.Command("systemctl", "list-dependencies", "spacewalk.target", "--plain")
	out, error := command.Output()
	result := make([]string, 0)
	if error != nil {
		return result, error
	}

	for _, service := range strings.Split(string(out), "\n") {
		sTrim := strings.TrimSpace(service)
		if len(sTrim) > 0 {
			result = append(result, sTrim)
		}
	}
	result = append(result, "jabberd.service", "osa-dispatcher.service")
	return result, nil
}

type UnitStatus struct {
	Unit       string
	Status     string
	StatusData string
}

type SystemUnitStatus struct {
	Services []UnitStatus
}

func getUnitStatus(unit string) UnitStatus {
	command := exec.Command("systemctl", "show", unit, "--no-pager", "--property=ActiveState,StateChangeTimestamp")
	out, error := command.Output()
	result := UnitStatus{Unit: unit, Status: "unknown", StatusData: "unknown"}
	if error != nil {
		fmt.Println(error)
		return result
	}

	for _, prop := range strings.Split(string(out), "\n") {
		pTrim := strings.TrimSpace(prop)
		if len(pTrim) > 0 {
			index := strings.Index(pTrim, "=")
			prop := strings.TrimSpace(pTrim[:index])
			if prop == "ActiveState" {
				result.Status = strings.TrimSpace(pTrim[index+1:])
			} else if prop == "StateChangeTimestamp" {
				result.StatusData = strings.TrimSpace(pTrim[index+1:])
			}
		}
	}

	return result
}

func GetServiceStatus() (SystemUnitStatus, error) {
	units, error := getServicesName()
	if error != nil {
		return SystemUnitStatus{}, error
	}

	result := make([]UnitStatus, 0)
	for _, unit := range units {
		result = append(result, getUnitStatus(unit))
	}
	return SystemUnitStatus{result}, nil
}
