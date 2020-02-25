package typbuildtool

import (
	"context"
	"errors"
	"fmt"

	"github.com/typical-go/typical-go/pkg/git"
	"github.com/typical-go/typical-go/pkg/typbuild"
	"github.com/typical-go/typical-go/pkg/typrls"
)

// ReleaseOption is option for release
type ReleaseOption struct {
	Alpha     bool
	Force     bool
	NoTest    bool
	NoBuild   bool
	NoPublish bool
}

func (b *BuildTool) release(ctx context.Context, c *typbuild.Context, opt *ReleaseOption) (err error) {
	if b.releaser == nil {
		return errors.New("Releaser is missing")
	}

	var (
		tag      string
		latest   string
		gitLogs  []*git.Log
		binaries []string
	)

	if !opt.NoBuild && b.builder != nil {
		if _, err = b.builder.Build(ctx, c); err != nil {
			return
		}
	}

	if !opt.NoTest {
		if err = b.test(ctx, c); err != nil {
			return
		}
	}

	if err = git.Fetch(ctx); err != nil {
		return fmt.Errorf("Failed git fetch: %w", err)
	}
	defer git.Fetch(ctx)

	if status := git.Status(ctx); status != "" && !opt.Force {
		return fmt.Errorf("Please commit changes first:\n%s", status)
	}
	if latest = git.LatestTag(ctx); latest == tag && !opt.Force {
		return fmt.Errorf("%s already released", latest)
	}
	if gitLogs = git.Logs(ctx, latest); len(gitLogs) < 1 && !opt.Force {
		return errors.New("No change to be released")
	}

	rls := &typrls.Context{
		Context: c,
		Name:    c.Name,
		Tag:     b.releaser.Tag(ctx, c.Version, opt.Alpha),
		GitLogs: gitLogs,
		Alpha:   opt.Alpha,
	}
	if binaries, err = b.releaser.Build(ctx, rls); err != nil {
		return fmt.Errorf("Failed build release: %w", err)
	}
	if !opt.NoPublish {
		if err = b.releaser.Publish(ctx, rls, binaries); err != nil {
			return fmt.Errorf("Failed publish: %w", err)
		}
	}
	return
}
