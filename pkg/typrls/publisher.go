package typrls

// Publisher reponsible to publish the release to external source
type Publisher interface {
	Publish(*Release) error
}

// Release information
type Release struct {
	Name       string
	Tag        string
	Binaries   []string
	ChangeLogs []string
	Alpha      bool
}
