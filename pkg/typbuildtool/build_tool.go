package typbuildtool

import (
	"fmt"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typbuild"
	"github.com/typical-go/typical-go/pkg/typbuild/prebld"
	"github.com/typical-go/typical-go/pkg/typclean"
	"github.com/typical-go/typical-go/pkg/typmock"
	"github.com/typical-go/typical-go/pkg/typrls"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// BuildTool is typical Build Tool for golang
type BuildTool struct {
	commanders []Commander
	builder    typbuild.Builder
	mocker     typmock.Mocker
	cleaner    typclean.Cleaner
	releaser   typrls.Releaser

	declStore *prebld.DeclStore
}

// New return new instance of build
func New() *BuildTool {
	return &BuildTool{
		builder:  typbuild.New(),
		mocker:   typmock.New(),
		cleaner:  typclean.New(),
		releaser: typrls.New(),
	}
}

// AppendCommander to return build with appended commander
func (b *BuildTool) AppendCommander(commanders ...Commander) *BuildTool {
	b.commanders = append(b.commanders, commanders...)
	return b
}

// WithtBuilder return new BuildTool with new builder
func (b *BuildTool) WithtBuilder(builder typbuild.Builder) *BuildTool {
	b.builder = builder
	return b
}

// WithRelease return BuildTool with new releaser
func (b *BuildTool) WithRelease(releaser typrls.Releaser) *BuildTool {
	b.releaser = releaser
	return b
}

// WithMocker return BuildTool with mocker
func (b *BuildTool) WithMocker(mocker typmock.Mocker) *BuildTool {
	b.mocker = mocker
	return b
}

// Validate build
func (b *BuildTool) Validate() (err error) {

	if err = common.Validate(b.builder); err != nil {
		return fmt.Errorf("BuildTool: Builder: %w", err)
	}

	if err = common.Validate(b.mocker); err != nil {
		return fmt.Errorf("BuildTool: Mocker: %w", err)
	}

	if err = common.Validate(b.releaser); err != nil {
		return fmt.Errorf("BuildTool: Releaser: %w", err)
	}

	return
}

// Run build tool
func (b *BuildTool) Run(t *typcore.TypicalContext) (err error) {
	if b.declStore, err = prebld.Walk(t.Files); err != nil {
		return
	}

	app := cli.NewApp()
	app.Name = t.Name
	app.Usage = "" // NOTE: intentionally blank
	app.Description = t.Description
	app.Version = t.Version
	app.Commands = b.Commands(&Context{
		TypicalContext: t,
	})

	return app.Run(os.Args)
}

// Commands to return command
func (b *BuildTool) Commands(c *Context) (cmds []*cli.Command) {
	if b.builder != nil {
		cmds = append(cmds, b.buildCommands(c)...)
	}
	if b.cleaner != nil {
		cmds = append(cmds, b.cleanCommands(c)...)

	}
	cmds = append(cmds,
		&cli.Command{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   "Run the testing",
			Action: func(cliCtx *cli.Context) error {
				return b.test(cliCtx.Context, c.TypicalContext)
			},
		},
	)

	if b.mocker != nil {
		cmds = append(cmds, b.mockCommands(c)...)
	}

	if b.releaser != nil {
		cmds = append(cmds, b.releaseCommands(c)...)
	}
	for _, commanders := range b.commanders {
		cmds = append(cmds, commanders.Commands(c)...)
	}
	return cmds
}

func (b *BuildTool) buildCommands(c *Context) []*cli.Command {
	buildCtx := &typbuild.Context{
		TypicalContext: c.TypicalContext,
		DeclStore:      b.declStore,
	}
	return []*cli.Command{
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "Build the binary",
			Action: func(cliCtx *cli.Context) (err error) {
				_, err = b.builder.Build(cliCtx.Context, buildCtx)
				return
			},
		},
		{
			Name:            "run",
			Aliases:         []string{"r"},
			Usage:           "Run the binary",
			SkipFlagParsing: true,
			Action: func(cliCtx *cli.Context) (err error) {
				var (
					binary string
					ctx    = cliCtx.Context
					args   = cliCtx.Args().Slice()
				)
				if binary, err = b.builder.Build(ctx, buildCtx); err != nil {
					return
				}
				log.Info("Run the application")
				cmd := exec.CommandContext(ctx, binary, args...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Stdin = os.Stdin
				return cmd.Run()
			},
		},
	}
}

func (b *BuildTool) releaseCommands(c *Context) []*cli.Command {
	buildCtx := &typbuild.Context{
		TypicalContext: c.TypicalContext,
		DeclStore:      b.declStore,
	}
	return []*cli.Command{
		{
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
				return b.release(cliCtx.Context, buildCtx, &ReleaseOption{
					Alpha:     cliCtx.Bool("alpha"),
					Force:     cliCtx.Bool("force"),
					NoTest:    cliCtx.Bool("no-test"),
					NoBuild:   cliCtx.Bool("no-build"),
					NoPublish: cliCtx.Bool("no-publish"),
				})
			},
		},
	}
}

func (b *BuildTool) mockCommands(c *Context) []*cli.Command {
	return []*cli.Command{
		{
			Name:  "mock",
			Usage: "Generate mock class",
			Action: func(cliCtx *cli.Context) (err error) {
				if b.mocker == nil {
					panic("Mocker is nil")
				}
				return b.mocker.Mock(cliCtx.Context, &typmock.Context{
					TypicalContext: c.TypicalContext,
					DeclStore:      b.declStore,
				})
			},
		},
	}
}

func (b *BuildTool) cleanCommands(c *Context) []*cli.Command {
	return []*cli.Command{
		{
			Name:    "clean",
			Aliases: []string{"c"},
			Usage:   "Clean the project from generated file during build time",
			Action: func(cliCtx *cli.Context) error {

				return b.cleaner.Clean(cliCtx.Context, &typclean.Context{
					TypicalContext: c.TypicalContext,
				})
			},
		},
	}
}
