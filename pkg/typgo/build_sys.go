package typgo

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/urfave/cli/v2"
)

type (
	// BuildSys detail
	BuildSys struct {
		*Descriptor
		ASTStore *typast.ASTStore
		Imports  []string
		Commands []*cli.Command
	}
)

func createBuildSys(d *Descriptor) *BuildSys {
	appDirs, appFiles := WalkLayout(d.Layouts)
	astStore, _ := typast.CreateASTStore(appFiles...)
	imports := retrImports(appDirs)

	sys := &BuildSys{
		Descriptor: d,
		ASTStore:   astStore,
		Imports:    imports,
	}

	for _, cmd := range d.Cmds {
		sys.Commands = append(sys.Commands, cmd.Command(sys))
	}

	return sys
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

func (b *BuildSys) app() *cli.App {
	cli.AppHelpTemplate = appHelpTemplate
	cli.SubcommandHelpTemplate = subcommandHelpTemplate

	app := cli.NewApp()
	app.Commands = b.Commands
	return app
}

// Context of build-cli
func (b *BuildSys) Context(c *cli.Context) *Context {
	return &Context{
		Context:  c,
		BuildSys: b,
	}
}

// ActionFn to return related action func
func (b *BuildSys) ActionFn(fn ExecuteFn) func(*cli.Context) error {
	return func(cli *cli.Context) error {
		return fn(b.Context(cli))
	}
}

// WalkLayout return dirs and files
func WalkLayout(layouts []string) (dirs, files []string) {
	for _, layout := range layouts {
		filepath.Walk(layout, func(path string, info os.FileInfo, err error) error {
			if info == nil {
				return nil
			}

			if info.IsDir() {
				dirs = append(dirs, path)
				return nil
			}

			if isGoSource(path) {
				files = append(files, path)
			}
			return nil
		})
	}
	return
}

func isGoSource(path string) bool {
	return strings.HasSuffix(path, ".go") &&
		!strings.HasSuffix(path, "_test.go")
}
