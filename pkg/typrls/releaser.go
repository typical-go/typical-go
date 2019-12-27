package typrls

import (
	"errors"
	"fmt"
)

// Releaser responsible to release distruction
type Releaser struct {
	name       string
	targets    []Target
	publishers []Publisher
	filter     Filter
	Tagging
}

// Tagging is setting how to make tag
type Tagging struct {
	IncludeBranch   bool
	IncludeCommitID bool
}

// New return new instance of releaser
func New(targets ...Target) *Releaser {
	return &Releaser{
		targets: targets,
		filter: &StandardFilter{
			Ignorings: []string{
				"merge",
				"bump",
				"revision",
				"generate",
				"wip",
			}},
	}
}

// WithName to set name and return its instance
func (r *Releaser) WithName(name string) *Releaser {
	r.name = name
	return r
}

// WithPublisher to set the publisher and return its instance
func (r *Releaser) WithPublisher(publishers ...Publisher) *Releaser {
	r.publishers = publishers
	return r
}

// WithFilter to set filter and return its instance
func (r *Releaser) WithFilter(filter Filter) *Releaser {
	r.filter = filter
	return r
}

// WithStandardFilter to set filter and return its instance
func (r *Releaser) WithStandardFilter(ignorings ...string) *Releaser {
	r.filter = &StandardFilter{ignorings}
	return r
}

// Validate the releaser
func (r *Releaser) Validate() (err error) {
	if len(r.targets) < 1 {
		return errors.New("Missing 'Targets'")
	}
	for _, target := range r.targets {
		if err = target.Validate(); err != nil {
			return fmt.Errorf("Target: %w", err)
		}
	}
	return
}
