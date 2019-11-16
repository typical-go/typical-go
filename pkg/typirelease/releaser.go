package typirelease

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/typical-go/typical-go/pkg/typienv"
	"github.com/typical-go/typical-go/pkg/utility/bash"
	"github.com/typical-go/typical-go/pkg/utility/git"
)

// Releaser responsible to release distruction
type Releaser struct {
	Name                string
	Targets             []ReleaseTarget
	Version             string
	WithGitBranch       bool
	WithLatestGitCommit bool

	Publisher
}

// Release the distribution
func (r *Releaser) Release(force, alpha bool) (rls *Release, err error) {
	var latestTag string
	var changeLogs []string
	var binaries []string
	git.Fetch()
	defer git.Fetch()
	name := r.releaseName()
	tag := r.releaseTag(alpha)
	if status := git.Status(); status != "" && !force {
		err = fmt.Errorf("Please commit changes first:\n%s", status)
		return
	}
	if latestTag = git.LatestTag(); latestTag == tag && !force {
		err = fmt.Errorf("%s already released", latestTag)
		return
	}
	if changeLogs = git.Logs(latestTag); len(changeLogs) < 1 && !force {
		err = errors.New("No change to be released")
		return
	}
	changeLogs = r.filter(changeLogs)
	for _, target := range r.Targets {
		var binary string
		if binary, err = r.build(name, tag, target); err != nil {
			return
		}
		binaries = append(binaries, binary)
	}
	rls = &Release{
		Name:       name,
		Tag:        tag,
		Alpha:      alpha,
		ChangeLogs: changeLogs,
		Binaries:   binaries,
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
			return fmt.Errorf("Target: %s", err.Error())
		}
	}
	return
}

func (r *Releaser) build(name, tag string, target ReleaseTarget) (binary string, err error) {
	goos := target.OS()
	goarch := target.Arch()
	binary = strings.Join([]string{name, tag, goos, goarch}, "_")
	binaryPath := fmt.Sprintf("%s/%s", typienv.Release, binary)
	// TODO: Support CGO
	envs := []string{"GOOS=" + goos, "GOARCH=" + goarch}
	if err = bash.GoBuild(binaryPath, typienv.App.SrcPath, envs...); err != nil {
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

func (r *Releaser) releaseTag(alpha bool) string {
	var b strings.Builder
	b.WriteString("v")
	b.WriteString(r.Version)
	if r.WithGitBranch {
		b.WriteString("_")
		b.WriteString(git.Branch())
	}
	if r.WithLatestGitCommit {
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
