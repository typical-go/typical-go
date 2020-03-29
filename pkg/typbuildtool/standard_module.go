package typbuildtool

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

var (
	_ Builder  = (*StandardModule)(nil)
	_ Cleaner  = (*StandardModule)(nil)
	_ Tester   = (*StandardModule)(nil)
	_ Releaser = (*StandardModule)(nil)
	_ Utility  = (*StandardModule)(nil)
)

// StandardModule is standard build module for go project
type StandardModule struct {
	stdout         io.Writer
	stderr         io.Writer
	stdin          io.Reader
	coverProfile   string
	releaseTargets []ReleaseTarget
	releaseFolder  string

	testTimeout time.Duration
}

// StandardBuild return new instance of Module
func StandardBuild() *StandardModule {
	return &StandardModule{
		stdout:        os.Stdout,
		stderr:        os.Stderr,
		stdin:         os.Stdin,
		coverProfile:  "cover.out",
		releaseFolder: "release",
		testTimeout:   20 * time.Second,
	}
}

// WithStdout return StandardModule with new stdout
func (b *StandardModule) WithStdout(stdout io.Writer) *StandardModule {
	b.stdout = stdout
	return b
}

// WithStderr return StandardModule with new stderr
func (b *StandardModule) WithStderr(stderr io.Writer) *StandardModule {
	b.stderr = stderr
	return b
}

// WithStdin return StandardModule with new stderr
func (b *StandardModule) WithStdin(stdin io.Reader) *StandardModule {
	b.stdin = stdin
	return b
}

// WithCoverProfile return StandardModule with new cover profile
func (b *StandardModule) WithCoverProfile(coverProfile string) *StandardModule {
	b.coverProfile = coverProfile
	return b
}

// WithReleaseTargets return StandardModule with new releaseTarget
func (b *StandardModule) WithReleaseTargets(releaseTargets ...ReleaseTarget) *StandardModule {
	b.releaseTargets = releaseTargets
	return b
}

// WithReleaseFolder return StandardModule with new release folder
func (b *StandardModule) WithReleaseFolder(releaseFolder string) *StandardModule {
	b.releaseFolder = releaseFolder
	return b
}

// WithTestTimeout return StandardModule with new testTimeout
func (b *StandardModule) WithTestTimeout(testTimeout time.Duration) *StandardModule {
	b.testTimeout = testTimeout
	return b
}

// Validate the releaser
func (b *StandardModule) Validate() (err error) {
	for _, target := range b.releaseTargets {
		if err = target.Validate(); err != nil {
			return fmt.Errorf("Target: %w", err)
		}
	}
	return
}

// Commands of build-tool
func (b *StandardModule) Commands(c *Context) []*cli.Command {
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
