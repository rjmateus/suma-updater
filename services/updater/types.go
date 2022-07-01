package updater

import (
	"encoding/xml"
)

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
	Status      string  `xml:"status,attr"`
	Category    string  `xml:"category,attr"`
	Severity    string  `xml:"severity,attr"`
	Pkgmanager  string  `xml:"pkgmanager,attr"`
	Restart     string  `xml:"restart,attr"`
	Interactive string  `xml:"interactive,attr"`
	StatusSince string  `xml:"status-since"`
	IssueDate   string  `xml:"issue-date"`
	Issues      []Issue `xml:"issue-list>issue"`
}

type ZypperUpdatesResultPatch struct {
	XMLName        xml.Name  `xml:"stream" json:"-"`
	Messages       []Message `xml:"message"`
	Updates        []Patch   `xml:"update-status>update-list>update"`
	BlockedUpdates []Patch   `xml:"update-status>blocked-update-list>update"`
}
type ZypperUpdatesResultUpdates struct {
	XMLName  xml.Name  `xml:"stream" json:"-"`
	Messages []Message `xml:"message"`
	Updates  []Update  `xml:"update-status>update-list>update"`
}

type Message struct {
	XMLName xml.Name `xml:"message" json:"-"`
	Type    string   `xml:"type,attr"`
	Message string   `xml:",chardata"`
}

type SolvableUpdate struct {
	XMLName     xml.Name `xml:"solvable" json:"-"`
	Type        string   `xml:"type,attr"`
	Name        string   `xml:"name,attr"`
	Edition     string   `xml:"edition,attr"`
	Arch        string   `xml:"arch,attr"`
	Repository  string   `xml:"repository,attr"`
	EditionOld  string   `xml:"edition-old,attr"`
	ArchOld     string   `xml:"arch-old,attr"`
	Summary     string   `xml:"summary,attr"`
	Description string   `xml:"description"`
}

type InstallSumary struct {
	XMLName          xml.Name         `xml:"install-summary" json:"-"`
	DownloadSize     string           `xml:"download-size,attr"`
	SpaceUsageDiff   string           `xml:"space-usage-diff,attr"`
	PackagesToChange string           `xml:"packages-to-change,attr"`
	NeedRestart      string           `xml:"need-restart,attr"`
	NeedReboot       string           `xml:"need-reboot,attr"`
	Updates          []SolvableUpdate `xml:"to-upgrade>solvable"`
	Patches          []SolvableUpdate `xml:"to-install>solvable"`
}

type Prompt struct {
	XMLName xml.Name `xml:"prompt" json:"-"`
}

type ZypperRunUpdateResult struct {
	XMLName  xml.Name      `xml:"stream" json:"-"`
	Messages []Message     `xml:"message"`
	Summary  InstallSumary `xml:"install-summary"`
}
