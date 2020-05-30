package typgo

import (
	"errors"
	"fmt"
	"strings"

	"github.com/typical-go/typical-go/pkg/git"
	"github.com/typical-go/typical-go/pkg/typvar"
	"github.com/urfave/cli/v2"
)

type (
	// Releaser responsible to release
	Releaser interface {
		Release(*Context) error
	}

	// Releases for composite release
	Releases []Releaser

	// ReleaseFn release function
	ReleaseFn func(*Context) error

	releaserImpl struct {
		fn ReleaseFn
	}
)

var _ Releaser = (Releases)(nil)

//
// releaserImpl
//

// NewRelease return new instance of Releaser
func NewRelease(fn ReleaseFn) Releaser {
	return &releaserImpl{fn: fn}
}

func (r *releaserImpl) Release(c *Context) error {
	return r.fn(c)
}

//
// Releaser
//

// Release the releasers
func (r Releases) Release(c *Context) (err error) {
	for _, releaser := range r {
		if err = releaser.Release(c); err != nil {
			return
		}
	}
	return
}

//
// Command
//

func cmdRelease(c *BuildCli) *cli.Command {
	return &cli.Command{
		Name:  "release",
		Usage: "Release the project",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "no-test", Usage: "skip the test"},
			&cli.BoolFlag{Name: "force", Usage: "Release by passed all validation"},
			&cli.BoolFlag{Name: "alpha", Usage: "Release for alpha version"},
		},
		Action: c.ActionFn("RELEASE", release),
	}
}

func release(c *Context) (err error) {
	if c.Release == nil {
		return errors.New("No Releaser")
	}

	if !c.Bool("no-test") {
		if err = test(c); err != nil {
			return
		}
	}

	ctx := c.Ctx()

	if err = git.Fetch(ctx); err != nil {
		return fmt.Errorf("Failed git fetch: %w", err)
	}
	defer git.Fetch(ctx)

	force := c.Bool("force")
	alpha := c.Bool("alpha")
	tag := releaseTag(c, alpha)

	if status := git.Status(ctx); status != "" && !force {
		return fmt.Errorf("Please commit changes first:\n%s", status)
	}

	latest := git.LatestTag(ctx)
	if latest == tag && !force {
		return fmt.Errorf("%s already released", latest)
	}

	gitLogs := git.RetrieveLogs(ctx, latest)
	if len(gitLogs) < 1 && !force {
		return errors.New("No change to be released")
	}

	typvar.Rls.Alpha = alpha
	typvar.Rls.Tag = tag
	typvar.Rls.GitLogs = gitLogs

	if err = c.Release.Release(c); err != nil {
		return
	}

	return

}

func releaseTag(c *Context, alpha bool) string {
	var builder strings.Builder
	builder.WriteString("v")
	builder.WriteString(c.Descriptor.Version)
	// if c.BuildTool.IncludeBranch {
	// 	builder.WriteString("_")
	// 	builder.WriteString(git.Branch(c.Context))
	// }
	// if c.BuildTool.IncludeCommitID {
	// 	builder.WriteString("_")
	// 	builder.WriteString(git.LatestCommit(c.Context))
	// }
	if alpha {
		builder.WriteString("_alpha")
	}
	return builder.String()
}
