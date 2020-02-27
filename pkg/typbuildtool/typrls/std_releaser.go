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
	targets       []Target
	publishers    []Publisher
	releaseFolder string
	Tagging
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
		releaseFolder: "release",
	}
}

// WithTarget to set target and return its instance
func (r *StdReleaser) WithTarget(targets ...Target) *StdReleaser {
	r.targets = targets
	return r
}

// WithPublisher return StdReleaser with new publisher
func (r *StdReleaser) WithPublisher(publishers ...Publisher) *StdReleaser {
	r.publishers = publishers
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
func (r *StdReleaser) Release(ctx context.Context, c *Context) (err error) {

	var (
		tag      string
		latest   string
		gitLogs  []*git.Log
		binaries []string
	)

	if err = git.Fetch(ctx); err != nil {
		return fmt.Errorf("Failed git fetch: %w", err)
	}
	defer git.Fetch(ctx)

	tag = r.Tag(ctx, c.Version, c.Alpha)

	if status := git.Status(ctx); status != "" && !c.Force {
		return fmt.Errorf("Please commit changes first:\n%s", status)
	}
	if latest = git.LatestTag(ctx); latest == tag && !c.Force {
		return fmt.Errorf("%s already released", latest)
	}
	if gitLogs = git.Logs(ctx, latest); len(gitLogs) < 1 && !c.Force {
		return errors.New("No change to be released")
	}

	for _, target := range r.targets {
		var binary string
		if binary, err = r.build(ctx, c, tag, target); err != nil {
			return fmt.Errorf("Failed build release: %w", err)
		}
		binaries = append(binaries, binary)
	}

	if !c.NoPublish {
		if err = r.Publish(ctx, &PublishContext{
			Context:  c,
			Tag:      tag,
			Binaries: binaries,
			GitLogs:  gitLogs,
		}); err != nil {
			return fmt.Errorf("Failed publish: %w", err)
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
func (r *StdReleaser) Publish(ctx context.Context, p *PublishContext) (err error) {
	for _, publisher := range r.publishers {
		if err = publisher.Publish(ctx, p); err != nil {
			return
		}
	}
	return
}

func (r *StdReleaser) build(ctx context.Context, rel *Context, tag string, target Target) (binary string, err error) {
	goos := target.OS()
	goarch := target.Arch()
	binary = strings.Join([]string{rel.Name, tag, goos, goarch}, "_")
	// TODO: Support CGO
	cmd := exec.CommandContext(ctx, "go", "build",
		"-o", fmt.Sprintf("%s/%s", r.releaseFolder, binary),
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
