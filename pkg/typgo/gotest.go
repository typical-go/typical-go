package typgo

import (
	"os"
	"time"

	"github.com/typical-go/typical-go/pkg/filekit"
	"github.com/urfave/cli/v2"
)

type (
	// GoTest command test
	GoTest struct {
		Timeout  time.Duration
		NoCover  bool
		Verbose  bool
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
func (t *GoTest) Task() *Task {
	return &Task{
		Name:            "test",
		Aliases:         []string{"t"},
		Usage:           "Test the project",
		SkipFlagParsing: true,
		Action:          t,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: coverprofileFlag, Usage: "override arguments"},
		},
	}
}

// Execute standard test
func (t *GoTest) Execute(c *Context) error {
	if t.Timeout == 0 {
		t.Timeout = 30 * time.Second
	}
	packages, err := filekit.FindDir(t.Includes, t.Excludes)
	if err != nil {
		return err
	}

	if len(packages) < 1 {
		c.Info("Nothing to test")
		return nil
	}

	args := []string{"test"}
	if !t.NoCover {
		args = append(args, "-cover")
	}
	if t.Verbose {
		args = append(args, "-v")
	}
	args = append(args, "-timeout="+t.Timeout.String())
	args = append(args, c.Args().Slice()...)
	args = append(args, packages...)

	return c.Execute(&Bash{
		Name:   "go",
		Args:   args,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})
}
