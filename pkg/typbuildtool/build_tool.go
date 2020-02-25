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
	commanders []typbuild.BuildCommander
	builder    typbuild.Builder
	releaser   typrls.Releaser
}

// New return new instance of build
func New() *BuildTool {
	return &BuildTool{
		builder:  typbuild.New(),
		releaser: typrls.New(),
	}
}

// AppendCommander to return build with appended commander
func (b *BuildTool) AppendCommander(commanders ...typbuild.BuildCommander) *BuildTool {
	b.commanders = append(b.commanders, commanders...)
	return b
}

// WithRelease to set releaser
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
func (b *BuildTool) Run(typCtx *typcore.TypicalContext) (err error) {
	var decls []*prebld.Declaration
	if decls, err = prebld.Walk(typCtx.Files); err != nil {
		return
	}

	c := &typbuild.Context{
		TypicalContext: typCtx,
		Declarations:   decls,
	}

	app := cli.NewApp()
	app.Name = c.Name
	app.Usage = "" // NOTE: intentionally blank
	app.Description = c.Description
	app.Version = c.Version
	app.Commands = b.BuildCommands(c)

	return app.Run(os.Args)
}

// BuildCommands to return command
func (b *BuildTool) BuildCommands(c *typbuild.Context) []*cli.Command {
	cmds := []*cli.Command{
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "Build the binary",
			Action: func(cliCtx *cli.Context) (err error) {
				if b.builder == nil {
					panic("Builder can't nil")
				}
				_, err = b.builder.Build(cliCtx.Context, c)
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
				// &cli.BoolFlag{Name: "no-delete", Usage: "Generate mock class with delete previous generation"},
			},
			Action: func(cliCtx *cli.Context) (err error) {
				return b.mock(cliCtx.Context, c, &MockOption{})
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
