package typgo

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/filekit"
	"github.com/urfave/cli/v2"
)

type (
	// GoTest command test
	GoTest struct {
		Args     []string
		Includes []string
		Excludes []string
	}
)

var _ CliCommander = (*GoTest)(nil)
var _ Action = (*GoTest)(nil)

// Cli command test
func (t *GoTest) Cli(b *BuildSys) *cli.Command {
	return &cli.Command{
		Name:    "test",
		Aliases: []string{"t"},
		Usage:   "Test the project",
		Action:  b.Action(t),
	}
}

// Execute standard test
func (t *GoTest) Execute(c *Context) error {
	packages, err := t.walk()
	if err != nil {
		return err
	}

	if len(packages) < 1 {
		fmt.Fprintln(Stdout, "Nothing to test")
		return nil
	}

	args := []string{"test", "-cover"}
	args = append(args, t.Args...)
	args = append(args, packages...)

	return c.Execute(&execkit.Command{
		Name:   "go",
		Args:   args,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})
}

func (t *GoTest) walk() (packages []string, err error) {
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if !filekit.MatchMulti(t.Excludes, path) && filekit.MatchMulti(t.Includes, path) && info.IsDir() {
			packages = append(packages, "./"+path)
		}
		return nil
	})
	return
}
