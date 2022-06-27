package serverStatus

import "encoding/xml"

type ServerStatus struct {
	MajorVersion string
	Release      string
	Arch         string
}

// Zypper package Search result
type Solvable struct {
	XMLName    xml.Name `xml:"solvable"`
	Status     string   `xml:"status,attr"`
	Name       string   `xml:"name,attr"`
	Kind       string   `xml:"kind,attr"`
	Edition    string   `xml:"edition,attr"`
	Arch       string   `xml:"arch,attr"`
	Repository string   `xml:"repository,attr"`
}

type ZypperSearchResult struct {
	XMLName  xml.Name   `xml:"stream"`
	Messages []Message  `xml:"message"`
	Solvable []Solvable `xml:"search-result>solvable-list>solvable"`
}

type Message struct {
	XMLName xml.Name `xml:"message"`
	Type    string   `xml:"type,attr"`
}
