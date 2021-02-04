package typgo

import (
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/filekit"
	"github.com/typical-go/typical-go/pkg/oskit"
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

const (
	coverprofileFlag = "coverprofile"
)

var _ Tasker = (*GoTest)(nil)
var _ Action = (*GoTest)(nil)

// Task for gotest
func (t *GoTest) Task(b *Descriptor) *cli.Command {
	return &cli.Command{
		Name:    "test",
		Aliases: []string{"t"},
		Usage:   "Test the project",
		Action:  b.Action(t),
		Flags: []cli.Flag{
			&cli.StringFlag{Name: coverprofileFlag, Usage: "override arguments"},
		},
	}
}

// Execute standard test
func (t *GoTest) Execute(c *Context) error {
	packages, err := filekit.FindDir(t.Includes, t.Excludes)
	if err != nil {
		return err
	}

	if len(packages) < 1 {
		fmt.Fprintln(oskit.Stdout, "Nothing to test")
		return nil
	}

	args := []string{"test"}
	if coverprofile := c.String(coverprofileFlag); coverprofile != "" {
		args = append(args, "-coverprofile="+coverprofile)
	} else {
		args = append(args, "-cover")
	}
	args = append(args, t.Args...)
	args = append(args, packages...)

	return c.Execute(&Bash{
		Name:   "go",
		Args:   args,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})
}
