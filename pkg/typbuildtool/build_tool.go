package typbuildtool

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typast"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// TypicalBuildTool is typical Build Tool for golang project
type TypicalBuildTool struct {
	commanders []Commander
	builder    Builder
	runner     Runner
	cleaner    Cleaner
	tester     Tester
	mocker     Mocker
	releaser   Releaser

	ast *typast.Ast
}

// New return new instance of build
func New() *TypicalBuildTool {
	return &TypicalBuildTool{
		builder:  NewBuilder(),
		runner:   NewRunner(),
		cleaner:  NewCleaner(),
		tester:   NewTester(),
		mocker:   NewMocker(),
		releaser: NewReleaser(),
	}
}

// AppendCommander to return BuildTool with appended commander
func (b *TypicalBuildTool) AppendCommander(commanders ...Commander) *TypicalBuildTool {
	b.commanders = append(b.commanders, commanders...)
	return b
}

// WithtBuilder return  BuildTool with new builder
func (b *TypicalBuildTool) WithtBuilder(builder Builder) *TypicalBuildTool {
	b.builder = builder
	return b
}

// WithRunner return BuildTool with appended runner
func (b *TypicalBuildTool) WithRunner(runner Runner) *TypicalBuildTool {
	b.runner = runner
	return b
}

// WithReleaser return BuildTool with new releaser
func (b *TypicalBuildTool) WithReleaser(releaser Releaser) *TypicalBuildTool {
	b.releaser = releaser
	return b
}

// WithMocker return BuildTool with new mocker
func (b *TypicalBuildTool) WithMocker(mocker Mocker) *TypicalBuildTool {
	b.mocker = mocker
	return b
}

// WithCleaner return BuildTool with new cleaner
func (b *TypicalBuildTool) WithCleaner(cleaner Cleaner) *TypicalBuildTool {
	b.cleaner = cleaner
	return b
}

// WithTester return BuildTool with new tester
func (b *TypicalBuildTool) WithTester(tester Tester) *TypicalBuildTool {
	b.tester = tester
	return b
}

// Validate build
func (b *TypicalBuildTool) Validate() (err error) {

	if err = common.Validate(b.builder); err != nil {
		return fmt.Errorf("BuildTool: Builder: %w", err)
	}

	if err = common.Validate(b.runner); err != nil {
		return fmt.Errorf("BuildTool: Runner: %w", err)
	}

	if err = common.Validate(b.mocker); err != nil {
		return fmt.Errorf("BuildTool: Mocker: %w", err)
	}

	if err = common.Validate(b.cleaner); err != nil {
		return fmt.Errorf("BuildTool: Cleaner: %w", err)
	}

	if err = common.Validate(b.tester); err != nil {
		return fmt.Errorf("BuildTool: Tester: %w", err)
	}

	if err = common.Validate(b.releaser); err != nil {
		return fmt.Errorf("BuildTool: Releaser: %w", err)
	}

	return
}

// Run build tool
func (b *TypicalBuildTool) Run(t *typcore.Context) (err error) {
	if b.ast, err = typast.Walk(t.Files); err != nil {
		return
	}

	app := cli.NewApp()
	app.Name = t.Name
	app.Usage = "" // NOTE: intentionally blank
	app.Description = t.Description
	app.Version = t.Version
	app.Commands = b.Commands(&Context{
		Context: t,
	})

	return app.Run(os.Args)
}

// Commands to return command
func (b *TypicalBuildTool) Commands(c *Context) (cmds []*cli.Command) {

	if b.builder != nil {
		cmds = append(cmds, b.buildCommand(c))
	}

	if b.runner != nil {
		cmds = append(cmds, b.runCommand(c))
	}

	if b.cleaner != nil {
		cmds = append(cmds, b.cleanCommand(c))
	}

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

func (b *TypicalBuildTool) buildCommand(c *Context) *cli.Command {
	return &cli.Command{
		Name:    "build",
		Aliases: []string{"b"},
		Usage:   "Build the binary",
		Action: func(cliCtx *cli.Context) (err error) {
			_, err = b.builder.Build(b.createBuildContext(cliCtx, c))
			return
		},
	}
}

func (b *TypicalBuildTool) runCommand(c *Context) *cli.Command {
	return &cli.Command{
		Name:            "run",
		Aliases:         []string{"r"},
		Usage:           "Run the binary",
		SkipFlagParsing: true,
		Action: func(cliCtx *cli.Context) (err error) {
			var binary string

			if binary, err = b.builder.Build(b.createBuildContext(cliCtx, c)); err != nil {
				return
			}

			log.Info("Run the application")
			return b.runner.Run(&RunContext{
				BuildContext: b.createBuildContext(cliCtx, c),
				Binary:       binary,
			})
		},
	}
}

func (b *TypicalBuildTool) releaseCommand(c *Context) *cli.Command {
	return &cli.Command{
		Name:  "release",
		Usage: "Release the distribution",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "no-test", Usage: "Release without run unit test"},
			&cli.BoolFlag{Name: "no-build", Usage: "Release without build"},
			&cli.BoolFlag{Name: "no-publish", Usage: "Release without create github release"},
			&cli.BoolFlag{Name: "force", Usage: "Release by passed all validation"},
			&cli.BoolFlag{Name: "alpha", Usage: "Release for alpha version"},
		},
		Action: func(cliCtx *cli.Context) (err error) {

			if !cliCtx.Bool("no-build") && b.builder != nil {
				if _, err = b.builder.Build(b.createBuildContext(cliCtx, c)); err != nil {
					return
				}
			}

			if !cliCtx.Bool("no-test") && b.tester != nil {
				if err = b.tester.Test(b.createBuildContext(cliCtx, c)); err != nil {
					return
				}
			}

			return b.releaser.Release(&ReleaseContext{
				BuildContext: b.createBuildContext(cliCtx, c),
				Alpha:        cliCtx.Bool("alpha"),
			})
		},
	}
}

func (b *TypicalBuildTool) mockCommand(c *Context) *cli.Command {
	return &cli.Command{
		Name:  "mock",
		Usage: "Generate mock class",
		Action: func(cliCtx *cli.Context) (err error) {
			if b.mocker == nil {
				panic("Mocker is nil")
			}
			return b.mocker.Mock(&MockContext{
				BuildContext: b.createBuildContext(cliCtx, c),
			})
		},
	}
}

func (b *TypicalBuildTool) cleanCommand(c *Context) *cli.Command {
	return &cli.Command{
		Name:    "clean",
		Aliases: []string{"c"},
		Usage:   "Clean the project from generated file during build time",
		Action: func(cliCtx *cli.Context) error {
			return b.cleaner.Clean(b.createBuildContext(cliCtx, c))
		},
	}
}

func (b *TypicalBuildTool) testCommand(c *Context) *cli.Command {
	return &cli.Command{
		Name:    "test",
		Aliases: []string{"t"},
		Usage:   "Run the testing",
		Action: func(cliCtx *cli.Context) error {
			return b.tester.Test(b.createBuildContext(cliCtx, c))
		},
	}
}

func (b *TypicalBuildTool) createBuildContext(cliCtx *cli.Context, c *Context) *BuildContext {
	return &BuildContext{
		Context: c,
		Cli:     cliCtx,
		Ast:     b.ast,
	}
}
