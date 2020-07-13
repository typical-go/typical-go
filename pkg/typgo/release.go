package typgo

import (
	"errors"
	"fmt"
	"strings"

	"github.com/typical-go/typical-go/pkg/git"
	"github.com/urfave/cli/v2"
)

type (
	// ReleaseCmd release command
	ReleaseCmd struct {
		Releaser
	}
	// Releaser responsible to release
	Releaser interface {
		Release(*ReleaseContext) error
	}
	// Releasers for composite release
	Releasers []Releaser
	// ReleaseContext contain data for release
	ReleaseContext struct {
		*Context
		Alpha   bool
		Tag     string
		GitLogs []*git.Log
	}
	// ReleaseFn release function
	ReleaseFn    func(*ReleaseContext) error
	releaserImpl struct {
		fn ReleaseFn
	}
)

//
// ReleaseCmd
//

var _ Cmd = (*ReleaseCmd)(nil)
var _ Action = (*ReleaseCmd)(nil)

// Command release
func (r *ReleaseCmd) Command(c *BuildCli) *cli.Command {
	return &cli.Command{
		Name:  "release",
		Usage: "Release the project",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "force", Usage: "Release by passed all validation"},
			&cli.BoolFlag{Name: "alpha", Usage: "Release for alpha version"},
		},
		Action: c.ActionFn(r.Execute),
	}
}

// Execute release
func (r *ReleaseCmd) Execute(c *Context) (err error) {

	ctx := c.Ctx()

	if err = git.Fetch(ctx); err != nil {
		return fmt.Errorf("Failed git fetch: %w", err)
	}
	defer git.Fetch(ctx)

	force := c.Bool("force")
	alpha := c.Bool("alpha")
	tag := releaseTag(c, alpha)

	status := git.Status(ctx)
	if status != "" && !force {
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

	return r.Release(&ReleaseContext{
		Context: c,
		Alpha:   alpha,
		Tag:     tag,
		GitLogs: gitLogs,
	})
}

//
// releaserImpl
//

// NewReleaser return new instance of Releaser
func NewReleaser(fn ReleaseFn) Releaser {
	return &releaserImpl{fn: fn}
}

func (r *releaserImpl) Release(c *ReleaseContext) error {
	return r.fn(c)
}

//
// Releaser
//

var _ Releaser = (Releasers)(nil)

// Release the releasers
func (r Releasers) Release(c *ReleaseContext) (err error) {
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

func releaseTag(c *Context, alpha bool) string {
	version := "0.0.1"
	if c.Descriptor.Version != "" {
		version = c.Descriptor.Version
	}

	var builder strings.Builder
	builder.WriteString("v")
	builder.WriteString(version)
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
