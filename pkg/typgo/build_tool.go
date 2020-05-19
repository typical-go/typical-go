package typgo

import (
	"fmt"
	"os"
	"strings"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typlog"
	"github.com/typical-go/typical-go/pkg/typtmpl"
	"github.com/typical-go/typical-go/pkg/typvar"
	"github.com/urfave/cli/v2"
)

type (
	// BuildTool detail
	BuildTool struct {
		*Descriptor
		ASTStore *typast.ASTStore
		Precond  *typtmpl.Precond
	}

	// Context of build tool
	Context struct {
		typlog.Logger

		Cli *cli.Context
		*BuildTool
	}

	// CliFunc is command line function
	CliFunc func(*Context) error

	// Preconditioner responsible to precondition
	Preconditioner interface {
		Precondition(*Context) error
	}
)

func createBuildTool(d *Descriptor) *BuildTool {
	var (
		astStore *typast.ASTStore
		err      error
	)
	appDirs, appFiles := WalkLayout(d.Layouts)

	if astStore, err = typast.CreateASTStore(appFiles...); err != nil {
		// TODO:
		// logger.Warn(err.Error())
	}

	return &BuildTool{
		Descriptor: d,
		ASTStore:   astStore,
		Precond: &typtmpl.Precond{
			Imports: retrImports(appDirs),
			Package: "main",
		},
	}
}

func launchBuildTool(d *Descriptor) error {
	app := cli.NewApp()
	app.Name = d.Name
	app.Usage = "Build-Tool"
	app.Description = d.Description
	app.Version = d.Version

	buildTool := createBuildTool(d)

	app.Before = beforeBuildTool(buildTool)
	app.Commands = buildTool.Commands()

	return app.Run(os.Args)
}

func beforeBuildTool(b *BuildTool) cli.BeforeFunc {
	return func(cli *cli.Context) (err error) {
		ctx := cli.Context
		precondFile := fmt.Sprintf("%s/%s/precond_DO_NOT_EDIT.go", typvar.CmdFolder, b.Name)

		logger := typlog.Logger{Name: "PRECOND"}

		if b.SkipPrecond {
			logger.Info("Skip the preconditon")
			return
		}

		os.Remove(precondFile)

		if err = b.Precondition(&Context{
			BuildTool: b,
			Cli:       cli,
			Logger:    logger,
		}); err != nil {
			return
		}

		if len(b.Precond.Lines) > 0 {
			logger.Infof("Write %s", precondFile)
			if err = typtmpl.WriteFile(precondFile, 0777, b.Precond); err != nil {
				return
			}
			if err = buildkit.GoImports(ctx, precondFile); err != nil {
				return
			}
		}
		return
	}
}

// Commands of build-tool
func (b *BuildTool) Commands() (cmds []*cli.Command) {
	cmds = []*cli.Command{
		cmdTest(b),
		cmdRun(b),
		cmdPublish(b),
		cmdClean(b),
	}

	if b.Utility != nil {
		for _, cmd := range b.Utility.Commands(b) {
			cmds = append(cmds, cmd)
		}
	}

	return cmds
}

// ActionFunc to return related action func
func (b *BuildTool) ActionFunc(name string, fn CliFunc) func(*cli.Context) error {
	return func(cli *cli.Context) error {
		return fn(&Context{
			Logger: typlog.Logger{
				Name: strings.ToUpper(name),
			},
			Cli:       cli,
			BuildTool: b,
		})
	}
}

func retrImports(dirs []string) []string {
	imports := []string{
		"github.com/typical-go/typical-go/pkg/typgo",
	}
	for _, dir := range dirs {
		if !strings.Contains(dir, "internal") {
			imports = append(imports, fmt.Sprintf("%s/%s", typvar.ProjectPkg, dir))
		}
	}
	return imports
}
