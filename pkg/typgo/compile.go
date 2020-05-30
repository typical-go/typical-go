package typgo

import (
	"errors"
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typtmpl"
	"github.com/typical-go/typical-go/pkg/typvar"
	"github.com/urfave/cli/v2"
)

type (
	// Compiler responsible to compile
	Compiler interface {
		Compile(*Context) error
	}

	// StdCompile is standard compile
	StdCompile struct{}
)

var _ (Compiler) = (*StdCompile)(nil)

//
// StdCompile
//

// Compile standard go project
func (*StdCompile) Compile(c *Context) (err error) {
	binary := fmt.Sprintf("%s/%s", typvar.BinFolder, c.Descriptor.Name)
	src := fmt.Sprintf("%s/%s", typvar.CmdFolder, c.Descriptor.Name)

	// NOTE: create main.go if not exist
	if _, err = os.Stat(src + "/main.go"); os.IsNotExist(err) {
		os.MkdirAll(src, 0777)

		if err = typtmpl.WriteFile(src+"/main.go", 0777, &typtmpl.AppMain{
			DescPkg: typvar.ProjectPkg + "/typical",
		}); err != nil {
			return
		}
	}

	return execute(c, &buildkit.GoBuild{
		Out:    binary,
		Source: "./" + src,
		Stderr: os.Stderr,
		Stdout: os.Stderr,
	})
}

//
// Command
//

func cmdCompile(c *BuildCli) *cli.Command {
	return &cli.Command{
		Name:   "compile",
		Usage:  "Compile the project",
		Action: c.ActionFn("COMPILE", compile),
	}
}

func compile(c *Context) error {
	if c.Compile == nil {
		return errors.New("compile is missing")
	}
	return c.Compile.Compile(c)
}
