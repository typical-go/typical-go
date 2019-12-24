package typrls

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/typical-go/typical-go/pkg/git"
	"github.com/typical-go/typical-go/pkg/typenv"
)

// Releaser responsible to release distruction
type Releaser struct {
	Name       string
	Targets    []Target
	Publishers []Publisher
	Tagging
}

// Tagging release settings
type Tagging struct {
	IncludeBranch   bool
	IncludeCommitID bool
}

// Release the distribution
func (r *Releaser) Release(name, version string, force, alpha, noPublish bool) (err error) {
	var latestTag string
	var changeLogs []string
	var binaries []string
	git.Fetch()
	defer git.Fetch()
	tag := r.releaseTag(version, alpha)
	if status := git.Status(); status != "" && !force {
		return fmt.Errorf("Please commit changes first:\n%s", status)
	}
	if latestTag = git.LatestTag(); latestTag == tag && !force {
		return fmt.Errorf("%s already released", latestTag)
	}
	if changeLogs = git.Logs(latestTag); len(changeLogs) < 1 && !force {
		return errors.New("No change to be released")
	}
	changeLogs = r.filter(changeLogs)
	for _, target := range r.Targets {
		var binary string
		if binary, err = r.build(name, tag, target); err != nil {
			return
		}
		binaries = append(binaries, binary)
	}
	rls := &Release{
		Name:       name,
		Tag:        tag,
		Alpha:      alpha,
		ChangeLogs: changeLogs,
		Binaries:   binaries,
	}
	if !noPublish {
		for _, publisher := range r.Publishers {
			if err = publisher.Publish(rls); err != nil {
				return
			}
		}
	}
	return
}

// Validate the releaser
func (r *Releaser) Validate() (err error) {
	if len(r.Targets) < 1 {
		return errors.New("Missing 'Targets'")
	}
	for _, target := range r.Targets {
		if err = target.Validate(); err != nil {
			return fmt.Errorf("Target: %w", err)
		}
	}
	return
}

func (r *Releaser) build(name, tag string, target Target) (binary string, err error) {
	goos := target.OS()
	goarch := target.Arch()
	binary = strings.Join([]string{name, tag, goos, goarch}, "_")
	// TODO: Support CGO
	cmd := exec.Command("go", "build",
		"-o", fmt.Sprintf("%s/%s", typenv.Layout.Release, binary),
		"-ldflags", "-w -s",
		"./"+typenv.AppMainPath,
	)
	cmd.Env = append(os.Environ(), "GOOS="+goos, "GOARCH="+goarch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	if err = cmd.Run(); err != nil {
		return
	}
	return
}

func (r *Releaser) filter(changeLogs []string) (filtered []string) {
	for _, log := range changeLogs {
		if !ignoring(log) {
			filtered = append(filtered, log)
		}
	}
	return
}

func (r *Releaser) releaseTag(version string, alpha bool) string {
	var b strings.Builder
	b.WriteString("v")
	b.WriteString(version)
	if r.IncludeBranch {
		b.WriteString("_")
		b.WriteString(git.Branch())
	}
	if r.IncludeCommitID {
		b.WriteString("_")
		b.WriteString(git.LatestCommit())
	}
	if alpha {
		b.WriteString("-alpha")
	}
	return b.String()
}

func (r *Releaser) releaseName() string {
	name := r.Name
	if name == "" {
		dir, _ := os.Getwd()
		name = filepath.Base(dir)
	}
	return name
}
