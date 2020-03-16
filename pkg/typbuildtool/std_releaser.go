package typbuildtool

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// StdReleaser responsible to release distruction
type StdReleaser struct {
	targets       []ReleaseTarget
	releaseFolder string
}

// NewReleaser return new instance of releaser
func NewReleaser() *StdReleaser {
	return &StdReleaser{
		targets: []ReleaseTarget{
			"linux/amd64",
			"darwin/amd64",
		},
		releaseFolder: "release",
	}
}

// WithTargets to set target and return its instance
func (r *StdReleaser) WithTargets(targets ...ReleaseTarget) *StdReleaser {
	r.targets = targets
	return r
}

// WithReleaseFolder return StdReleaser with new release folder
func (r *StdReleaser) WithReleaseFolder(releaseFolder string) *StdReleaser {
	r.releaseFolder = releaseFolder
	return r
}

// Validate the releaser
func (r *StdReleaser) Validate() (err error) {
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

// Release this project
func (r *StdReleaser) Release(c *ReleaseContext) (files []string, err error) {

	for _, target := range r.targets {
		var binary string
		if binary, err = r.build(c.Context, c.Tag, target); err != nil {
			err = fmt.Errorf("Failed build release: %w", err)
			return
		}
		files = append(files, binary)
	}

	return
}

func (r *StdReleaser) build(c *Context, tag string, target ReleaseTarget) (binary string, err error) {
	ctx := c.Cli.Context
	goos := target.OS()
	goarch := target.Arch()
	binary = strings.Join([]string{c.Name, tag, goos, goarch}, "_")
	// TODO: Support CGO
	cmd := exec.CommandContext(ctx, "go", "build",
		"-o", fmt.Sprintf("%s/%s", r.releaseFolder, binary),
		"-ldflags", "-w -s",
		fmt.Sprintf("./%s/%s", c.CmdFolder(), c.Name),
	)
	cmd.Env = append(os.Environ(), "GOOS="+goos, "GOARCH="+goarch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	if err = cmd.Run(); err != nil {
		return
	}
	return
}
