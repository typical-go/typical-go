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
	// BuildCli detail
	BuildCli struct {
		*Descriptor
		ASTStore *typast.ASTStore
		Imports  []string
	}
)

func createBuildCli(d *Descriptor) *BuildCli {
	appDirs, appFiles := WalkLayout(d.Layouts)
	astStore, _ := typast.CreateASTStore(appFiles...)
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

func (b *BuildCli) commands() []*cli.Command {
	var cmds []*cli.Command
	for _, command := range b.Commands {
		cmds = append(cmds, command.Command(b))
	}
	return cmds
}

// Context of build-cli
func (b *BuildCli) Context(c *cli.Context) *Context {
	return &Context{
		Context:    c,
		Descriptor: b.Descriptor,
		ASTStore:   b.ASTStore,
		Imports:    b.Imports,
	}
}

// ActionFn to return related action func
func (b *BuildCli) ActionFn(fn ExecuteFn) func(*cli.Context) error {
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
