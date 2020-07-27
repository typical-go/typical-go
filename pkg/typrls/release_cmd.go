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
	// ReleaseCmd release command
	ReleaseCmd struct {
		Releaser
		Before     typgo.Action
		ReleaseTag string
		Alpha      bool
		Validation Validator
		Summary    Summarizer
	}
)

var _ typgo.Cmd = (*ReleaseCmd)(nil)
var _ typgo.Action = (*ReleaseCmd)(nil)

// Command release
func (r *ReleaseCmd) Command(sys *typgo.BuildSys) *cli.Command {
	return &cli.Command{
		Name:  "release",
		Usage: "Release the project",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: ForceFlag, Usage: "Release by passed all validation"},
			&cli.BoolFlag{Name: AlphaFlag, Usage: "Release for alpha version"},
			&cli.StringFlag{Name: TagFlag, Usage: "Override the release-tag"},
		},
		Before: sys.ActionFn(r.Before),
		Action: sys.ActionFn(r),
	}
}

// Execute release
func (r *ReleaseCmd) Execute(c *typgo.Context) error {
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
func (r *ReleaseCmd) GetReleaseTag(c *typgo.Context) string {
	if r.ReleaseTag == "" {
		if c.BuildSys.Version != "" {
			r.ReleaseTag = fmt.Sprintf("v%s", c.BuildSys.Version)
		} else {
			r.ReleaseTag = "v0.0.1"
		}

		if r.Alpha {
			r.ReleaseTag = r.ReleaseTag + "_alpha"
		}
	}

	return r.ReleaseTag
}
