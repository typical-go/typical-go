package typbuild

import (
	"context"
	"errors"
	"fmt"

	"github.com/typical-go/typical-go/pkg/git"
	"github.com/typical-go/typical-go/pkg/typenv"
)

// ReleaseOption is option for release
type ReleaseOption struct {
	Alpha     bool
	Force     bool
	NoTest    bool
	NoBuild   bool
	NoPublish bool
}

func (b *Build) release(ctx context.Context, c *Context, opt *ReleaseOption) (err error) {
	if b.releaser == nil {
		return errors.New("Releaser is missing")
	}

	var (
		tag        string
		latest     string
		changeLogs []string
		binaries   []string
		name       = typenv.ProjectName
	)

	if !opt.NoBuild {
		if err = b.buildProject(ctx, c); err != nil {
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
	if tag, err = b.releaser.Tag(ctx, c.Version, opt.Alpha); err != nil {
		return fmt.Errorf("Failed generate tag: %w", err)
	}
	if status := git.Status(ctx); status != "" && !opt.Force {
		return fmt.Errorf("Please commit changes first:\n%s", status)
	}
	if latest = git.LatestTag(ctx); latest == tag && !opt.Force {
		return fmt.Errorf("%s already released", latest)
	}
	if changeLogs = git.Logs(ctx, latest); len(changeLogs) < 1 && !opt.Force {
		return errors.New("No change to be released")
	}
	if binaries, err = b.releaser.BuildRelease(ctx, name, tag, changeLogs, opt.Alpha); err != nil {
		return fmt.Errorf("Failed build release: %w", err)
	}
	if !opt.NoPublish {
		if err = b.releaser.Publish(ctx, name, tag, changeLogs, binaries, opt.Alpha); err != nil {
			return fmt.Errorf("Failed publish: %w", err)
		}
	}
	return
}
