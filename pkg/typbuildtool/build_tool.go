package typbuildtool

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typbuild"
	"github.com/typical-go/typical-go/pkg/typbuild/prebld"
	"github.com/typical-go/typical-go/pkg/typrls"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// BuildTool is typical Build Tool for golang
type BuildTool struct {
	commanders []BuildCommander
	builder    typbuild.Builder
	releaser   typrls.Releaser

	declStore *prebld.DeclStore
}

// New return new instance of build
func New() *BuildTool {
	return &BuildTool{
		builder:  typbuild.New(),
		releaser: typrls.New(),
	}
}

// AppendCommander to return build with appended commander
func (b *BuildTool) AppendCommander(commanders ...BuildCommander) *BuildTool {
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

// Validate build
func (b *BuildTool) Validate() (err error) {
	if b.releaser != nil {
		if err = common.Validate(b.releaser); err != nil {
			return fmt.Errorf("Build: Releaser: %w", err)
		}
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
	app.Commands = b.BuildCommands(&Context{
		TypicalContext: t,
	})

	return app.Run(os.Args)
}

// BuildCommands to return command
func (b *BuildTool) BuildCommands(c *Context) []*cli.Command {
	buildCtx := &typbuild.Context{
		TypicalContext: c.TypicalContext,
		DeclStore:      b.declStore,
	}
	cmds := []*cli.Command{
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "Build the binary",
			Action: func(cliCtx *cli.Context) (err error) {
				if b.builder == nil {
					panic("Builder can't nil")
				}
				_, err = b.builder.Build(cliCtx.Context, buildCtx)
				return
			},
		},
		{
			Name:    "clean",
			Aliases: []string{"c"},
			Usage:   "Clean the project from generated file during build time",
			Action: func(cliCtx *cli.Context) error {
				return b.clean(cliCtx.Context, c)
			},
		},
		{
			Name:            "run",
			Aliases:         []string{"r"},
			Usage:           "Run the binary",
			SkipFlagParsing: true,
			Action: func(cliCtx *cli.Context) (err error) {
				return b.run(cliCtx.Context, buildCtx, cliCtx.Args().Slice())
			},
		},
		{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   "Run the testing",
			Action: func(cliCtx *cli.Context) error {
				return b.test(cliCtx.Context, c.TypicalContext)
			},
		},
		{
			Name:  "mock",
			Usage: "Generate mock class",
			Flags: []cli.Flag{
				// &cli.BoolFlag{Name: "no-delete", Usage: "Generate mock class with delete previous generation"},
			},
			Action: func(cliCtx *cli.Context) (err error) {
				return b.mock(cliCtx.Context, buildCtx, &MockOption{})
			},
		},
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
	for _, commanders := range b.commanders {
		cmds = append(cmds, commanders.BuildCommands(c)...)
	}
	return cmds
}

func (b *BuildTool) run(ctx context.Context, c *typbuild.Context, args []string) (err error) {
	var binary string
	if binary, err = b.builder.Build(ctx, c); err != nil {
		return
	}
	log.Info("Run the application")
	cmd := exec.CommandContext(ctx, binary, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
