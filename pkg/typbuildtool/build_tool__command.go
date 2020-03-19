package typbuildtool

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/git"
	"github.com/urfave/cli/v2"
)

// Commands to return command
func (b *TypicalBuildTool) Commands(c *Context) (cmds []*cli.Command) {

	cmds = []*cli.Command{
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "Build the binary",
			Action: func(cliCtx *cli.Context) (err error) {
				_, err = b.Build(b.createContext(c, cliCtx))
				return
			},
		},
		{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   "Run the testing",
			Action: func(cliCtx *cli.Context) error {
				return b.Test(b.createContext(c, cliCtx))
			},
		},
		{
			Name:            "run",
			Aliases:         []string{"r"},
			Usage:           "Run the binary",
			SkipFlagParsing: true,
			Action: func(cliCtx *cli.Context) (err error) {
				bc := b.createContext(c, cliCtx)

				var dists []BuildDistribution
				if dists, err = b.Build(bc); err != nil {
					return
				}

				for _, dist := range dists {
					if err = dist.Run(bc); err != nil {
						return
					}
				}
				return
			},
		},
		{
			Name:    "clean",
			Aliases: []string{"c"},
			Usage:   "Clean the project from generated file during build time",
			Flags: []cli.Flag{
				&cli.BoolFlag{Name: "short", Aliases: []string{"s"}, Usage: "Short version of clean only clean build-tool"},
			},
			Action: func(cliCtx *cli.Context) (err error) {
				return b.Clean(b.createContext(c, cliCtx))
			},
		},
		{
			Name:  "release",
			Usage: "Release the distribution",
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
					bc           = b.createContext(c, cliCtx)
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
					publishCtx := &PublishContext{
						ReleaseContext: rc,
						ReleaseFiles:   releaseFiles,
					}
					if err = b.Publish(publishCtx); err != nil {
						err = fmt.Errorf("Failed to publish: %w", err)
						return
					}
				}

				return
			},
		},
	}

	for _, module := range b.modules {
		if commander, ok := module.(Commander); ok {
			cmds = append(cmds, commander.Commands(c)...)
		}
	}

	for _, commander := range b.commanders {
		cmds = append(cmds, commander.Commands(c)...)
	}
	return cmds
}
