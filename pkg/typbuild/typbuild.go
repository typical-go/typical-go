package typbuild

import (
	"context"
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typbuild/prebld"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// Build tool
type Build struct {
	commanders  []BuildCommander
	prebuilders []Prebuilder
	releaser    Releaser
}

// Prebuilder responsible to prebuild task
type Prebuilder interface {
	Prebuild(ctx context.Context, c *Context) error
}

// BuildCommander responsible to return commands for Build-Tool
type BuildCommander interface {
	BuildCommands(c *Context) []*cli.Command
}

// New return new instance of build
func New() *Build {
	return &Build{
		prebuilders: []Prebuilder{&standardPrebuilder{}},
	}
}

// AppendCommander to return build with appended commander
func (b *Build) AppendCommander(commanders ...BuildCommander) *Build {
	b.commanders = append(b.commanders, commanders...)
	return b
}

// WithRelease to set releaser
func (b *Build) WithRelease(releaser Releaser) *Build {
	b.releaser = releaser
	return b
}

// WithPrebuild to set prebuilder
func (b *Build) WithPrebuild(prebuilders ...Prebuilder) *Build {
	b.prebuilders = append(b.prebuilders, prebuilders...)
	return b
}

// Validate build
func (b *Build) Validate() (err error) {
	if b.releaser != nil {
		if err = common.Validate(b.releaser); err != nil {
			return fmt.Errorf("Build: Releaser: %w", err)
		}
	}
	return
}

// Run build tool
func (b *Build) Run(bctx *typcore.BuildContext) (err error) {
	var decls []*prebld.Declaration
	if decls, err = prebld.Walk(bctx.Files); err != nil {
		return
	}

	c := &Context{
		BuildContext: bctx,
		Declarations: decls,
	}

	app := cli.NewApp()
	app.Name = bctx.Name
	app.Usage = "" // NOTE: intentionally blank
	app.Description = bctx.Description
	app.Version = bctx.Version
	app.Commands = b.BuildCommands(c)

	return app.Run(os.Args)
}

// BuildCommands to return command
func (b *Build) BuildCommands(c *Context) []*cli.Command {
	cmds := []*cli.Command{
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "Build the binary",
			Action: func(cliCtx *cli.Context) (err error) {
				return b.buildProject(cliCtx.Context, c)
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
				return b.run(cliCtx.Context, c, cliCtx.Args().Slice())
			},
		},
		{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   "Run the testing",
			Action: func(cliCtx *cli.Context) error {
				return b.test(cliCtx.Context, c)
			},
		},
		{
			Name:  "mock",
			Usage: "Generate mock class",
			Flags: []cli.Flag{
				&cli.BoolFlag{Name: "no-delete", Usage: "Generate mock class with delete previous generation"},
			},
			Action: func(cliCtx *cli.Context) (err error) {
				return b.mock(cliCtx.Context, c, &MockOption{
					NoDelete: cliCtx.Bool("no-delete"),
				})
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
				return b.release(cliCtx.Context, c, &ReleaseOption{
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
