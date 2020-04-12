package typbuildtool

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

var (
	_ Cleaner  = (*StdBuild)(nil)
	_ Tester   = (*StdBuild)(nil)
	_ Releaser = (*StdBuild)(nil)
	_ Utility  = (*StdBuild)(nil)
	_ Runner   = (*StdBuild)(nil)
)

// StdBuild is standard build module for go project
type StdBuild struct {
	stdout         io.Writer
	stderr         io.Writer
	stdin          io.Reader
	coverProfile   string
	releaseTargets []ReleaseTarget
	releaseFolder  string

	testTimeout time.Duration
}

// StandardBuild return new instance of Module
func StandardBuild() *StdBuild {
	return &StdBuild{
		stdout:        os.Stdout,
		stderr:        os.Stderr,
		stdin:         os.Stdin,
		coverProfile:  "cover.out",
		releaseFolder: "release",
		testTimeout:   20 * time.Second,
	}
}

// WithStdout return StandardModule with new stdout
func (b *StdBuild) WithStdout(stdout io.Writer) *StdBuild {
	b.stdout = stdout
	return b
}

// WithStderr return StandardModule with new stderr
func (b *StdBuild) WithStderr(stderr io.Writer) *StdBuild {
	b.stderr = stderr
	return b
}

// WithStdin return StandardModule with new stderr
func (b *StdBuild) WithStdin(stdin io.Reader) *StdBuild {
	b.stdin = stdin
	return b
}

// WithCoverProfile return StandardModule with new cover profile
func (b *StdBuild) WithCoverProfile(coverProfile string) *StdBuild {
	b.coverProfile = coverProfile
	return b
}

// WithReleaseTargets return StandardModule with new releaseTarget
func (b *StdBuild) WithReleaseTargets(releaseTargets ...ReleaseTarget) *StdBuild {
	b.releaseTargets = releaseTargets
	return b
}

// WithReleaseFolder return StandardModule with new release folder
func (b *StdBuild) WithReleaseFolder(releaseFolder string) *StdBuild {
	b.releaseFolder = releaseFolder
	return b
}

// WithTestTimeout return StandardModule with new testTimeout
func (b *StdBuild) WithTestTimeout(testTimeout time.Duration) *StdBuild {
	b.testTimeout = testTimeout
	return b
}

// Validate the releaser
func (b *StdBuild) Validate() (err error) {
	for _, target := range b.releaseTargets {
		if err = target.Validate(); err != nil {
			return fmt.Errorf("Target: %w", err)
		}
	}
	return
}

// Commands of build-tool
func (b *StdBuild) Commands(c *Context) []*cli.Command {
	return []*cli.Command{
		{
			Name:  "mock",
			Usage: "Generate mock class",
			Action: func(cliCtx *cli.Context) (err error) {
				return b.Mock(c.BuildContext(cliCtx))
			},
		},
	}
}
