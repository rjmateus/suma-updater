package zypper

import (
	"encoding/xml"
	"os/exec"
)

func ZypperRef() (ZypperResult, error) {
	command := exec.Command("zypper", "--xmlout", "ref")

	out, error := command.Output()
	if error != nil {
		return ZypperResult{}, error
	}

	var outProcessed ZypperResult
	xml.Unmarshal(out, &outProcessed)
	return outProcessed, nil
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
