package commands

import (
	"encoding/xml"
	"errors"
	"fmt"
	"os/exec"
)

func ZypperRef() ZypperResult {
	command := exec.Command("zypper", "--xmlout", "ref")

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

	var outProcessed ZypperResult
	xml.Unmarshal(out, &outProcessed)
	return outProcessed
}

type ZypperResult struct {
	XMLName  xml.Name  `xml:"stream" json:"-"`
	Messages []Message `xml:"message"`
}

type Message struct {
	XMLName xml.Name `xml:"message" json:"-"`
	Type    string   `xml:"type,attr"`
	Message string   `xml:",chardata"`
}
