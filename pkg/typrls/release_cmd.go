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

	alpha := c.Bool(AlphaFlag)
	tagName := c.String(TagFlag)
	if tagName == "" {
		tagName = createTagName(c, alpha)
	}

	currentTag := gitTag(ctx)

	rlsCtx := &Context{
		Context: c,
		Alpha:   alpha,
		Git: &Git{
			Status:     gitStatus(ctx),
			CurrentTag: currentTag,
			Logs:       gitLogs(ctx, currentTag),
		},
		TagName: tagName,
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

func createTagName(c *typgo.Context, alpha bool) string {
	tagName := "v0.0.1"
	if c.BuildSys.ProjectVersion != "" {
		tagName = fmt.Sprintf("v%s", c.BuildSys.ProjectVersion)
	}

	if alpha {
		tagName = tagName + "_alpha"
	}

	return tagName
}
