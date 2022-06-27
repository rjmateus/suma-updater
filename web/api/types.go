package api

type ServerStatus struct {
	Version    string
	Release    string
	Arch       string
	NewVersion string
	NewRelease string
	Updates    int
	Patches    int
}

type ApiErro struct {
	Message string
}
