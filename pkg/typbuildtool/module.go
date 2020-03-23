package typbuildtool

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli/v2"
)

// Module of build-tool
type Module struct {
	stdout         io.Writer
	stderr         io.Writer
	stdin          io.Reader
	coverProfile   string
	releaseTargets []ReleaseTarget
	releaseFolder  string
}

// StandardBuild return new instance of Module
func StandardBuild() *Module {
	return &Module{
		stdout:       os.Stdout,
		stderr:       os.Stderr,
		stdin:        os.Stdin,
		coverProfile: "cover.out",
		releaseTargets: []ReleaseTarget{
			"linux/amd64",
			"darwin/amd64",
		},
		releaseFolder: "release",
	}
}

// WithStdout return StdBuilder with new stdout
func (b *Module) WithStdout(stdout io.Writer) *Module {
	b.stdout = stdout
	return b
}

// WithStderr return StdBuilder with new stderr
func (b *Module) WithStderr(stderr io.Writer) *Module {
	b.stderr = stderr
	return b
}

// WithStdin return StdBuilder with new stderr
func (b *Module) WithStdin(stdin io.Reader) *Module {
	b.stdin = stdin
	return b
}

// WithCoverProfile return StdTester with new cover profile
func (b *Module) WithCoverProfile(coverProfile string) *Module {
	b.coverProfile = coverProfile
	return b
}

// WithReleaseTargets to set target and return its instance
func (b *Module) WithReleaseTargets(releaseTargets ...ReleaseTarget) *Module {
	b.releaseTargets = releaseTargets
	return b
}

// WithReleaseFolder return StdReleaser with new release folder
func (b *Module) WithReleaseFolder(releaseFolder string) *Module {
	b.releaseFolder = releaseFolder
	return b
}

// Validate the releaser
func (b *Module) Validate() (err error) {
	if len(b.releaseTargets) < 1 {
		return errors.New("Missing 'Targets'")
	}
	for _, target := range b.releaseTargets {
		if err = target.Validate(); err != nil {
			return fmt.Errorf("Target: %w", err)
		}
	}
	return
}

// Commands of build-tool
func (b *Module) Commands(c *Context) []*cli.Command {
	return []*cli.Command{
		{
			Name:  "mock",
			Usage: "Generate mock class",
			Action: func(cliCtx *cli.Context) (err error) {
				return b.Mock(&BuildContext{
					Context: c,
					Cli:     cliCtx,
				})
			},
		},
	}
}
