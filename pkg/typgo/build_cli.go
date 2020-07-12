package typgo

import (
	"errors"
	"fmt"
	"strings"

	"github.com/typical-go/typical-go/pkg/git"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typlog"
	"github.com/urfave/cli/v2"
)

var (
	// ConfigFile location
	ConfigFile = ".env"
)

type (
	// BuildCli detail
	BuildCli struct {
		*Descriptor
		ASTStore *typast.ASTStore
		Imports  []string
	}
	// CliFunc is command line function
	CliFunc func(*Context) error
)

func createBuildCli(d *Descriptor) *BuildCli {
	var (
		astStore *typast.ASTStore
		err      error
	)
	appDirs, appFiles := WalkLayout(d.Layouts)

	if astStore, err = typast.CreateASTStore(appFiles...); err != nil {
		// TODO:
		// logger.Warn(err.Error())
	}

	imports := retrImports(appDirs)
	return &BuildCli{
		Descriptor: d,
		ASTStore:   astStore,
		Imports:    imports,
	}
}

func retrImports(dirs []string) []string {
	imports := []string{
		"github.com/typical-go/typical-go/pkg/typapp",
	}

	for _, dir := range dirs {
		imports = append(imports, fmt.Sprintf("%s/%s", ProjectPkg, dir))
	}
	return imports
}

func (b *BuildCli) commands() ([]*cli.Command, error) {
	var cmds []*cli.Command
	if b.Test != nil {
		cmds = append(cmds, cmdTest(b))
	}
	if b.Compile != nil {
		cmds = append(cmds, cmdCompile(b))
	}
	if b.Run != nil {
		cmds = append(cmds, cmdRun(b))
	}
	if b.Release != nil {
		cmds = append(cmds, cmdRelease(b))
	}
	if b.Clean != nil {
		cmds = append(cmds, cmdClean(b))
	}

	if b.Utility != nil {
		cmds0, err := b.Utility.Commands(b)
		if err != nil {
			return nil, err
		}
		cmds = append(cmds, cmds0...)
	}

	return cmds, nil
}

// Context of build-cli
func (b *BuildCli) Context(name string, c *cli.Context) *Context {
	return &Context{
		Logger: typlog.Logger{
			Name: name,
		},
		Context:  c,
		BuildCli: b,
	}
}

// ActionFn to return related action func
func (b *BuildCli) ActionFn(name string, fn CliFunc) func(*cli.Context) error {
	return func(cli *cli.Context) error {
		c := b.Context(strings.ToUpper(name), cli)
		return fn(c)
	}
}

func cmdTest(c *BuildCli) *cli.Command {
	return &cli.Command{
		Name:    "test",
		Aliases: []string{"t"},
		Usage:   "Test the project",
		Action:  c.ActionFn("TEST", test),
	}
}

func test(c *Context) error {
	if c.Test == nil {
		return errors.New("test is missing")
	}
	return c.Test.Test(c)
}

func cmdCompile(c *BuildCli) *cli.Command {
	return &cli.Command{
		Name:    "compile",
		Aliases: []string{"c"},
		Usage:   "Compile the project",
		Action:  c.ActionFn("COMPILE", compile),
	}
}

func compile(c *Context) error {
	if c.Compile == nil {
		return errors.New("compile is missing")
	}
	return c.Compile.Compile(c)
}

func cmdRun(c *BuildCli) *cli.Command {
	return &cli.Command{
		Name:            "run",
		Aliases:         []string{"r"},
		Usage:           "Run the project in local environment",
		SkipFlagParsing: true,
		Action:          c.ActionFn("RUN", run),
	}
}

func run(c *Context) error {
	if c.Run == nil {
		return errors.New("run is missing")
	}
	if err := compile(c); err != nil {
		return err
	}
	return c.Run.Run(c)
}

func cmdRelease(c *BuildCli) *cli.Command {
	return &cli.Command{
		Name:  "release",
		Usage: "Release the project",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "skip-test", Usage: "skip test"},
			&cli.BoolFlag{Name: "skip-compile", Usage: "skip compile"},
			&cli.BoolFlag{Name: "force", Usage: "Release by passed all validation"},
			&cli.BoolFlag{Name: "alpha", Usage: "Release for alpha version"},
		},
		Action: c.ActionFn("RELEASE", release),
	}
}

func release(c *Context) (err error) {
	if c.Release == nil {
		return errors.New("No Releaser")
	}

	if !c.Bool("skip-test") {
		if err = test(c); err != nil {
			return
		}
	}

	if !c.Bool("skip-compile") {
		if err = compile(c); err != nil {
			return
		}
	}

	ctx := c.Ctx()

	if err = git.Fetch(ctx); err != nil {
		return fmt.Errorf("Failed git fetch: %w", err)
	}
	defer git.Fetch(ctx)

	force := c.Bool("force")
	alpha := c.Bool("alpha")
	tag := releaseTag(c, alpha)

	status := git.Status(ctx)
	if status != "" && !force {
		return fmt.Errorf("Please commit changes first:\n%s", status)
	}

	latest := git.LatestTag(ctx)
	if latest == tag && !force {
		return fmt.Errorf("%s already released", latest)
	}

	gitLogs := git.RetrieveLogs(ctx, latest)
	if len(gitLogs) < 1 && !force {
		return errors.New("No change to be released")
	}

	return c.Release.Release(&ReleaseContext{
		Context: c,
		Alpha:   alpha,
		Tag:     tag,
		GitLogs: gitLogs,
	})
}

func cmdClean(c *BuildCli) *cli.Command {
	return &cli.Command{
		Name:   "clean",
		Usage:  "Clean the project",
		Action: c.ActionFn("CLEAN", c.Clean.Clean),
	}
}
