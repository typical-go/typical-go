package typrls

import (
	"errors"
	"fmt"

	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

const (
	// ForceFlag -force cli param
	ForceFlag = "force"
	// AlphaFlag -alpha cli param
	AlphaFlag = "alpha"
	// TagFlag -tag cli param
	TagFlag = "tag"
)

type (
	// Command release command
	Command struct {
		Releaser
		ReleaseTag string
		Alpha      bool
		Validation Validator
		Summary    Summarizer
	}
)

var _ typgo.Cmd = (*Command)(nil)
var _ typgo.Action = (*Command)(nil)

// Command release
func (r *Command) Command(c *typgo.BuildSys) *cli.Command {
	return &cli.Command{
		Name:  "release",
		Usage: "Release the project",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: ForceFlag, Usage: "Release by passed all validation"},
			&cli.BoolFlag{Name: AlphaFlag, Usage: "Release for alpha version"},
			&cli.StringFlag{Name: TagFlag, Usage: "Override the release-tag"},
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

	gitFetch(ctx)
	defer gitFetch(ctx)

	r.ReleaseTag = c.String(TagFlag)
	r.Alpha = r.Alpha || c.Bool(AlphaFlag)
	currentTag := gitTag(ctx)

	rlsCtx := &Context{
		Context: c,
		Alpha:   r.Alpha,
		Git: &Git{
			Status:     gitStatus(ctx),
			CurrentTag: currentTag,
			Logs:       gitLogs(ctx, currentTag),
		},
		ReleaseTag: r.GetReleaseTag(c),
	}

	if r.Validation != nil && !c.Bool(ForceFlag) {
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

		if r.Alpha {
			r.ReleaseTag = r.ReleaseTag + "_alpha"
		}
	}

	return r.ReleaseTag
}
