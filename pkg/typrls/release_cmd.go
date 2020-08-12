package typrls

import (
	"errors"
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

const (
	// ForceFlag -force cli param
	ForceFlag = "force"
	// AlphaFlag -alpha cli param
	AlphaFlag = "alpha"
	// TagNameFlag -tag cli param
	TagNameFlag = "tag-name"
)

type (
	// ReleaseCmd release command
	ReleaseCmd struct {
		Releaser
		Before        typgo.Action
		Validation    Validator
		Summary       Summarizer
		ReleaseFolder string
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
			&cli.StringFlag{Name: TagNameFlag, Usage: "Override the release-tag"},
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

	currentTag := gitTag(ctx)
	alpha := c.Bool(AlphaFlag)
	tagName := c.String(TagNameFlag)

	if tagName == "" {
		tagName = createTagName(c, alpha)
	}

	if r.ReleaseFolder == "" {
		r.ReleaseFolder = "release"
	}

	rlsCtx := &Context{
		Context: c,
		Alpha:   alpha,
		Git: &Git{
			Status:     gitStatus(ctx),
			CurrentTag: currentTag,
			Logs:       gitLogs(ctx, currentTag),
		},
		TagName:       tagName,
		ReleaseFolder: r.ReleaseFolder,
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

	os.RemoveAll(r.ReleaseFolder)
	os.MkdirAll(r.ReleaseFolder, 0777)

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
