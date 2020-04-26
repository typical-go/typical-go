package typbuildtool

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/typical-go/typical-go/pkg/git"
	"github.com/urfave/cli/v2"
)

func (b *BuildTool) cmdPublish(c *Context) *cli.Command {

	return &cli.Command{
		Name:    "publish",
		Usage:   "Publish the project",
		Aliases: []string{"p"},
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "no-test", Usage: "skip the test"},
			&cli.BoolFlag{Name: "force", Usage: "Release by passed all validation"},
			&cli.BoolFlag{Name: "alpha", Usage: "Release for alpha version"},
		},
		Action: c.ActionFunc(b.release),
	}
}

func (b *BuildTool) release(c *CliContext) (err error) {

	var (
		rc           *ReleaseContext
		releaseFiles []string
	)
	if err = git.Fetch(c.Context); err != nil {
		return fmt.Errorf("Failed git fetch: %w", err)
	}
	defer git.Fetch(c.Context)

	if !c.Bool("no-test") {
		if err = test(c); err != nil {
			return
		}
	}

	if rc, err = b.releaseContext(c); err != nil {
		return
	}

	if releaseFiles, err = b.Release(rc); err != nil {
		return
	}

	pc := &PublishContext{
		ReleaseContext: rc,
		ReleaseFiles:   releaseFiles,
	}

	if err = b.Publish(pc); err != nil {
		return
	}

	return

}

// Publish the project
func (b *BuildTool) Publish(pc *PublishContext) (err error) {
	for _, module := range b.buildSequences {
		if publisher, ok := module.(Publisher); ok {
			if err = publisher.Publish(pc); err != nil {
				return
			}
		}
	}
	return
}

// Release the project
func (b *BuildTool) Release(rc *ReleaseContext) (files []string, err error) {
	for _, module := range b.buildSequences {
		if releaser, ok := module.(Releaser); ok {
			var files1 []string
			if files1, err = releaser.Release(rc); err != nil {
				return
			}
			files = append(files, files1...)
		}
	}
	return
}

func (b *BuildTool) releaseContext(c *CliContext) (*ReleaseContext, error) {

	force := c.Bool("force")
	alpha := c.Bool("alpha")
	tag := b.releaseTag(c.Context, c.Core.Version, alpha)

	if status := git.Status(c.Context); status != "" && !force {
		return nil, fmt.Errorf("Please commit changes first:\n%s", status)
	}

	var latest string
	if latest = git.LatestTag(c.Context); latest == tag && !force {
		return nil, fmt.Errorf("%s already released", latest)
	}

	var gitLogs []*git.Log
	if gitLogs = git.RetrieveLogs(c.Context, latest); len(gitLogs) < 1 && !force {
		return nil, errors.New("No change to be released")
	}

	return &ReleaseContext{
		CliContext: c,
		Alpha:      alpha,
		Tag:        tag,
		GitLogs:    gitLogs,
	}, nil
}

// Tag return relase tag
func (b *BuildTool) releaseTag(ctx context.Context, version string, alpha bool) string {
	var builder strings.Builder
	builder.WriteString("v")
	builder.WriteString(version)
	if b.includeBranch {
		builder.WriteString("_")
		builder.WriteString(git.Branch(ctx))
	}
	if b.includeCommitID {
		builder.WriteString("_")
		builder.WriteString(git.LatestCommit(ctx))
	}
	if alpha {
		builder.WriteString("_alpha")
	}
	return builder.String()
}
