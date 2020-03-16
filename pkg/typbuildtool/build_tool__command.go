package typbuildtool

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/git"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// Commands to return command
func (b *TypicalBuildTool) Commands(c *typcore.Context) (cmds []*cli.Command) {

	if b.builder != nil {
		cmds = append(cmds,
			b.buildCommand(c),
			b.runCommand(c),
		)
	}

	cmds = append(cmds, b.cleanCommand(c))

	if b.tester != nil {
		cmds = append(cmds, b.testCommand(c))
	}

	if b.mocker != nil {
		cmds = append(cmds, b.mockCommand(c))
	}

	if b.releaser != nil {
		cmds = append(cmds, b.releaseCommand(c))
	}
	for _, commanders := range b.commanders {
		cmds = append(cmds, commanders.Commands(c)...)
	}
	return cmds
}

func (b *TypicalBuildTool) buildCommand(c *typcore.Context) *cli.Command {
	return &cli.Command{
		Name:    "build",
		Aliases: []string{"b"},
		Usage:   "Build the binary",
		Action: func(cliCtx *cli.Context) (err error) {
			_, err = b.Build(b.createContext(c, cliCtx))
			return
		},
	}
}

func (b *TypicalBuildTool) testCommand(c *typcore.Context) *cli.Command {
	return &cli.Command{
		Name:    "test",
		Aliases: []string{"t"},
		Usage:   "Run the testing",
		Action: func(cliCtx *cli.Context) error {
			return b.Test(b.createContext(c, cliCtx))
		},
	}
}

func (b *TypicalBuildTool) runCommand(c *typcore.Context) *cli.Command {
	return &cli.Command{
		Name:            "run",
		Aliases:         []string{"r"},
		Usage:           "Run the binary",
		SkipFlagParsing: true,
		Action: func(cliCtx *cli.Context) (err error) {
			var dist BuildDistribution
			bc := b.createContext(c, cliCtx)

			if dist, err = b.Build(bc); err != nil {
				return
			}

			if dist != nil {
				if err = dist.Run(bc); err != nil {
					return
				}
			}
			return
		},
	}
}

func (b *TypicalBuildTool) mockCommand(c *typcore.Context) *cli.Command {
	return &cli.Command{
		Name:  "mock",
		Usage: "Generate mock class",
		Action: func(cliCtx *cli.Context) (err error) {
			if b.mocker == nil {
				panic("Mocker is nil")
			}
			return b.mocker.Mock(b.createContext(c, cliCtx))
		},
	}
}

func (b *TypicalBuildTool) cleanCommand(c *typcore.Context) *cli.Command {
	return &cli.Command{
		Name:    "clean",
		Aliases: []string{"c"},
		Usage:   "Clean the project from generated file during build time",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "short", Aliases: []string{"s"}, Usage: "Short version of clean only clean build-tool"},
		},
		Action: func(cliCtx *cli.Context) (err error) {
			return b.Clean(b.createContext(c, cliCtx))
		},
	}
}

func (b *TypicalBuildTool) releaseCommand(c *typcore.Context) *cli.Command {
	return &cli.Command{
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
	}
}
