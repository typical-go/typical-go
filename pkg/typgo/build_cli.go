package typgo

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typlog"
	"github.com/typical-go/typical-go/pkg/typtmpl"
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
		Precond  *typtmpl.Precond
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

	return &BuildCli{
		Descriptor: d,
		ASTStore:   astStore,
		Precond: &typtmpl.Precond{
			Imports: retrImports(appDirs),
			Package: "main",
		},
	}
}

func retrImports(dirs []string) []string {
	imports := []string{
		"github.com/typical-go/typical-go/pkg/typgo",
	}
	for _, dir := range dirs {
		if !strings.Contains(dir, "internal") {
			imports = append(imports, fmt.Sprintf("%s/%s", ProjectPkg, dir))
		}
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

// Prebuild process
func (b *BuildCli) Prebuild() (err error) {
	c := &PrebuildContext{
		BuildCli: b,
		ctx:      context.Background(),
	}
	if c.Descriptor.Prebuild != nil {
		if err := c.Descriptor.Prebuild.Prebuild(c); err != nil {
			return err
		}
	}
	if err := savePrecond(c); err != nil {
		return fmt.Errorf("save-precond: %w", err)
	}

	if envs, _ := LoadConfig(ConfigFile); len(envs) > 0 {
		printEnv(os.Stdout, envs)
	}

	return
}

func printEnv(w io.Writer, envs map[string]string) {
	color.New(color.FgGreen).Fprint(w, "ENV")
	fmt.Fprint(w, ": ")

	for key := range envs {
		fmt.Fprintf(w, "+%s ", key)
	}
	fmt.Fprintln(w)
}

func savePrecond(c *PrebuildContext) error {
	path := Precond(c.Descriptor.Name)
	os.Remove(path)
	if c.Precond.NotEmpty() {
		if err := writeGoSource(c.Precond, path); err != nil {
			return err
		}
	}
	return nil
}

// ActionFn to return related action func
func (b *BuildCli) ActionFn(name string, fn CliFunc) func(*cli.Context) error {
	return func(cli *cli.Context) error {
		c := b.Context(strings.ToUpper(name), cli)
		return fn(c)
	}
}
