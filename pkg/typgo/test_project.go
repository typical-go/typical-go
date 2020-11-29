package typgo

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/filekit"
	"github.com/urfave/cli/v2"
)

const (
	defaultTestTimeout      = 25 * time.Second
	defaultTestCoverProfile = "cover.out"
)

type (
	// TestProject command test
	TestProject struct {
		Timeout      time.Duration
		CoverProfile string
		Race         bool
		Patterns     []string
	}
)

var _ Cmd = (*TestProject)(nil)
var _ Action = (*TestProject)(nil)

// Command test
func (t *TestProject) Command(b *BuildSys) *cli.Command {
	return &cli.Command{
		Name:    "test",
		Aliases: []string{"t"},
		Usage:   "Test the project",
		Action:  b.Action(t),
	}
}

// Execute standard test
func (t *TestProject) Execute(c *Context) error {
	packages, err := t.walk()
	if err != nil {
		return err
	}

	if len(packages) < 1 {
		fmt.Fprintln(Stdout, "Nothing to test")
		return nil
	}

	return c.Execute(&execkit.GoTest{
		Packages:     packages,
		Timeout:      t.getTimeout(),
		CoverProfile: t.getCoverProfile(),
		Race:         t.Race,
	})
}

func (t *TestProject) walk() (packages []string, err error) {
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if filekit.MatchMulti(t.Patterns, path) && info.IsDir() {
			packages = append(packages, "./"+path)
		}
		return nil
	})
	return
}

func (t *TestProject) getTimeout() time.Duration {
	if t.Timeout <= 0 {
		return defaultTestTimeout
	}
	return t.Timeout
}

func (t *TestProject) getCoverProfile() string {
	if t.CoverProfile == "" {
		return defaultTestCoverProfile
	}
	return t.CoverProfile
}
