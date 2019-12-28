package typrls

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/typical-go/typical-go/pkg/git"
	"github.com/typical-go/typical-go/pkg/typenv"
)

// BuildRelease the distribution
func (r *Releaser) BuildRelease(ctx context.Context, name, tag string, changeLogs []string, alpha bool) (binaries []string, err error) {
	if r.filter != nil {
		changeLogs = r.filter.Filter(changeLogs)
	}
	for _, target := range r.targets {
		var binary string
		if binary, err = r.build(ctx, name, tag, target); err != nil {
			return
		}
		binaries = append(binaries, binary)
	}
	return
}

// Publish the release
func (r *Releaser) Publish(ctx context.Context, name, tag string, changeLogs, binaries []string, alpha bool) (err error) {
	for _, publisher := range r.publishers {
		if err = publisher.Publish(ctx, name, tag, changeLogs, binaries, alpha); err != nil {
			return
		}
	}
	return
}

// Tag return relase tag
func (r *Releaser) Tag(ctx context.Context, version string, alpha bool) (tag string, err error) {
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
	tag = b.String()
	return
}

func (r *Releaser) build(ctx context.Context, name, tag string, target Target) (binary string, err error) {
	goos := target.OS()
	goarch := target.Arch()
	binary = strings.Join([]string{name, tag, goos, goarch}, "_")
	// TODO: Support CGO
	cmd := exec.CommandContext(ctx, "go", "build",
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

func (r *Releaser) releaseName() string {
	name := r.name
	if name == "" {
		dir, _ := os.Getwd()
		name = filepath.Base(dir)
	}
	return name
}
