package updater

import "encoding/xml"

type SearchResuls struct {
	XMLName xml.Name `xml:"search-result"`
	Version string   `xml:"version,attr"`
}

type Source struct {
	XMLName xml.Name `xml:"source" json:"-"`
	Url     string   `xml:"url,attr"`
	Alias   string   `xml:"alias,attr"`
}

type Update struct {
	XMLName     xml.Name `xml:"update" json:"-"`
	Kind        string   `xml:"kind,attr"`
	Name        string   `xml:"name,attr"`
	Edition     string   `xml:"edition,attr"`
	Arch        string   `xml:"arch,attr"`
	EditionOld  string   `xml:"edition-old,attr"`
	Summary     string   `xml:"summary"`
	Description string   `xml:"description"`
	License     string   `xml:"license"`
	Source      Source   `xml:"source"`
}

type Issue struct {
	XMLName xml.Name `xml:"issue" json:"-"`
	Type    string   `xml:"type,attr"`
	Id      string   `xml:"id,attr"`
	Title   string   `xml:"title"`
	Href    string   `xml:"href"`
}

type Patch struct {
	Update
	Category    string  `xml:"category,attr"`
	Severity    string  `xml:"severity,attr"`
	Pkgmanager  string  `xml:"pkgmanager,attr"`
	Restart     string  `xml:"restart,attr"`
	Interactive string  `xml:"interactive,attr"`
	StatusSince string  `xml:"status-since"`
	IssueDate   string  `xml:"issue-date"`
	Issues      []Issue `xml:"issue-list>issue"`
}

type ZypperUpdatesResult[T any] struct {
	XMLName  xml.Name  `xml:"stream" json:"-"`
	Messages []Message `xml:"message"`
	Updates  []T       `xml:"update-status>update-list>update"`
}

type Message struct {
	XMLName xml.Name `xml:"message" json:"-"`
	Type    string   `xml:"type,attr"`
}
