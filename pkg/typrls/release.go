package typrls

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/typical-go/typical-go/pkg/git"
	"github.com/typical-go/typical-go/pkg/typenv"
)

// Release the distribution
func (r *Releaser) Release(ctx context.Context, name, version string, force, alpha, noPublish bool) (err error) {
	var (
		latestTag  string
		changeLogs []string
		binaries   []string
	)
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
	if r.filter != nil {
		changeLogs = r.filter.Filter(changeLogs)
	}
	for _, target := range r.targets {
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
		for _, publisher := range r.publishers {
			if err = publisher.Publish(ctx, rls); err != nil {
				return
			}
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
	name := r.name
	if name == "" {
		dir, _ := os.Getwd()
		name = filepath.Base(dir)
	}
	return name
}
