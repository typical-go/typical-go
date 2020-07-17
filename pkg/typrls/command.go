package typrls

import (
	"errors"
	"fmt"

	"github.com/typical-go/typical-go/pkg/git"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

const (
	forceParam = "force"
	alphaParam = "alpha"
)

type (
	// Command release command
	Command struct {
		Releaser
		ReleaseTag string
		Validation Validator
		Summary    Summarizer
	}
)

var _ typgo.Cmd = (*Command)(nil)
var _ typgo.Action = (*Command)(nil)

// Command release
func (r *Command) Command(c *typgo.BuildCli) *cli.Command {
	return &cli.Command{
		Name:  "release",
		Usage: "Release the project",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: forceParam, Usage: "Release by passed all validation"},
			&cli.BoolFlag{Name: alphaParam, Usage: "Release for alpha version"},
		},
		Action: c.ActionFn(r.Execute),
	}
}

// Execute release
func (r *Command) Execute(c *typgo.Context) error {
	if r.Summary == nil {
		return errors.New("typrls: missing summary")
	}

	ctx := c.Ctx()

	git.Fetch(ctx)
	defer git.Fetch(ctx)

	currentTag := git.CurrentTag(ctx)

	rlsCtx := &Context{
		Context: c,
		Git: &Git{
			Status:     git.Status(ctx),
			CurrentTag: currentTag,
			Logs:       git.RetrieveLogs(ctx, currentTag),
		},
		ReleaseTag: r.GetReleaseTag(c),
	}

	if r.Validation != nil && !c.Bool(forceParam) {
		if err := r.Validation.Validate(rlsCtx); err != nil {
			return err
		}
	}

	summary, err := r.Summary.Summarize(rlsCtx)
	if err != nil {
		return err
	}
	rlsCtx.Summary = summary

	return r.Release(rlsCtx)
}

// GetReleaseTag return release tag
func (r *Command) GetReleaseTag(c *typgo.Context) string {
	if r.ReleaseTag == "" {
		if c.Descriptor.Version != "" {
			r.ReleaseTag = fmt.Sprintf("v%s", c.Descriptor.Version)
		} else {
			r.ReleaseTag = "v0.0.1"
		}

		if c.Bool(alphaParam) {
			r.ReleaseTag = r.ReleaseTag + "_alpha"
		}
	}

	return r.ReleaseTag
}
