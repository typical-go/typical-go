package typrls

import "context"

// Publisher reponsible to publish the release to external source
type Publisher interface {
	Publish(context.Context, *Release) error
}

// Release information
type Release struct {
	Name       string
	Tag        string
	Binaries   []string
	ChangeLogs []string
	Alpha      bool
}
