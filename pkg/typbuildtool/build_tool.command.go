package typbuildtool

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/typical-go/typical-go/pkg/git"
	"github.com/urfave/cli/v2"
)

// Commands to return command
func (b *BuildTool) Commands(c *Context) (cmds []*cli.Command) {

	cmds = []*cli.Command{
		{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   "Test the project",
			Action: func(cliCtx *cli.Context) error {
				return b.Test(c.BuildContext(cliCtx))
			},
		},
		{
			Name:            "run",
			Aliases:         []string{"r"},
			Usage:           "Run the project in local environment",
			SkipFlagParsing: true,
			Action: func(cliCtx *cli.Context) (err error) {
				return b.Run(c.BuildContext(cliCtx))
			},
		},
		{
			Name:    "clean",
			Aliases: []string{"c"},
			Usage:   "Clean the project",
			Action: func(cliCtx *cli.Context) (err error) {
				return b.Clean(c.BuildContext(cliCtx))
			},
		},
		{
			Name:  "release",
			Usage: "Create project release",
			Flags: []cli.Flag{
				&cli.BoolFlag{Name: "no-test", Usage: "Release without run unit test"},
				&cli.BoolFlag{Name: "no-publish", Usage: "Release without create github release"},
				&cli.BoolFlag{Name: "force", Usage: "Release by passed all validation"},
				&cli.BoolFlag{Name: "alpha", Usage: "Release for alpha version"},
			},
			Action: func(cliCtx *cli.Context) (err error) {
				if err = git.Fetch(cliCtx.Context); err != nil {
					return fmt.Errorf("Failed git fetch: %w", err)
				}
				defer git.Fetch(cliCtx.Context)

				var (
					rc           *ReleaseContext
					releaseFiles []string
					bc           = c.BuildContext(cliCtx)
				)

				if !cliCtx.Bool("no-test") {
					if err = b.Test(bc); err != nil {
						return
					}
				}

				if rc, err = b.createReleaseContext(bc); err != nil {
					return
				}

				if releaseFiles, err = b.Release(rc); err != nil {
					return
				}

				if !cliCtx.Bool("no-publish") {
					if err = b.Publish(&PublishContext{
						ReleaseContext: rc,
						ReleaseFiles:   releaseFiles,
					}); err != nil {
						err = fmt.Errorf("Failed to publish: %w", err)
						return
					}
				}

				return
			},
		},
	}

	for _, module := range b.buildSequences {
		if utility, ok := module.(Utility); ok {
			cmds = append(cmds, utility.Commands(c)...)
		}
	}

	for _, task := range b.utilities {
		cmds = append(cmds, task.Commands(c)...)
	}
	return cmds
}

func (b *BuildTool) createReleaseContext(c *BuildContext) (*ReleaseContext, error) {
	ctx := c.Cli.Context
	force := c.Cli.Bool("force")
	alpha := c.Cli.Bool("alpha")
	tag := b.releaseTag(ctx, c.Version, alpha)

	if status := git.Status(ctx); status != "" && !force {
		return nil, fmt.Errorf("Please commit changes first:\n%s", status)
	}

	var latest string
	if latest = git.LatestTag(ctx); latest == tag && !force {
		return nil, fmt.Errorf("%s already released", latest)
	}

	var gitLogs []*git.Log
	if gitLogs = git.RetrieveLogs(ctx, latest); len(gitLogs) < 1 && !force {
		return nil, errors.New("No change to be released")
	}

	return &ReleaseContext{
		BuildContext: c,
		Alpha:        alpha,
		Tag:          tag,
		GitLogs:      gitLogs,
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
