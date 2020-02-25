package typrls

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/typical-go/typical-go/pkg/git"
)

// StdReleaser responsible to release distruction
type StdReleaser struct {
	name       string
	targets    []Target
	publishers []Publisher
	Tagging
}

// Publisher reponsible to publish the release to external source
type Publisher interface {
	Publish(ctx context.Context, rel *Context, binaries []string) error
}

// Tagging is setting how to make tag
type Tagging struct {
	IncludeBranch   bool
	IncludeCommitID bool
}

// New return new instance of releaser
func New() *StdReleaser {
	return &StdReleaser{
		targets: []Target{
			"linux/amd64",
			"darwin/amd64",
		},
	}
}

// WithTarget to set target and return its instance
func (r *StdReleaser) WithTarget(targets ...Target) *StdReleaser {
	r.targets = targets
	return r
}

// WithName to set name and return its instance
func (r *StdReleaser) WithName(name string) *StdReleaser {
	r.name = name
	return r
}

// WithPublisher to set the publisher and return its instance
func (r *StdReleaser) WithPublisher(publishers ...Publisher) *StdReleaser {
	r.publishers = publishers
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

// Tag return relase tag
func (r *StdReleaser) Tag(ctx context.Context, version string, alpha bool) string {
	var b strings.Builder
	b.WriteString("v")
	b.WriteString(version)
	if r.IncludeBranch {
		b.WriteString("_")
		b.WriteString(git.Branch(ctx))
	}
	if r.IncludeCommitID {
		b.WriteString("_")
		b.WriteString(git.LatestCommit(ctx))
	}
	if alpha {
		b.WriteString("_alpha")
	}
	return b.String()
}

// Publish the release
func (r *StdReleaser) Publish(ctx context.Context, rel *Context, binaries []string) (err error) {
	for _, publisher := range r.publishers {
		if err = publisher.Publish(ctx, rel, binaries); err != nil {
			return
		}
	}
	return
}

// Build the distribution
func (r *StdReleaser) Build(ctx context.Context, rel *Context) (binaries []string, err error) {
	for _, target := range r.targets {
		var binary string
		if binary, err = build(ctx, rel, target); err != nil {
			return
		}
		binaries = append(binaries, binary)
	}
	return
}

func build(ctx context.Context, rel *Context, target Target) (binary string, err error) {
	goos := target.OS()
	goarch := target.Arch()
	binary = strings.Join([]string{rel.Name, rel.Tag, goos, goarch}, "_")
	// TODO: Support CGO
	cmd := exec.CommandContext(ctx, "go", "build",
		"-o", fmt.Sprintf("%s/%s", rel.ReleaseFolder, binary),
		"-ldflags", "-w -s",
		fmt.Sprintf("./%s/%s", rel.CmdFolder, rel.Name),
	)
	cmd.Env = append(os.Environ(), "GOOS="+goos, "GOARCH="+goarch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	if err = cmd.Run(); err != nil {
		return
	}
	return
}
