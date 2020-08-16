package typrls

import (
	"errors"
	"os"

	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

type (
	// ReleaseCmd release command
	ReleaseCmd struct {
		Before typgo.Action
		Action typgo.Action
	}
	// ReleaseProject action
	ReleaseProject struct {
		Validator  Validator
		Tagger     Tagger
		Summarizer Summarizer
		Releaser   Releaser
	}
)

var _ typgo.Cmd = (*ReleaseCmd)(nil)

const (
	// ForceFlag -force cli param
	ForceFlag = "force"
	// AlphaFlag -alpha cli param
	AlphaFlag = "alpha"
	// TagNameFlag -tag cli param
	TagNameFlag = "tag-name"
	// ReleaseFolderFlag -release-folder cli param
	ReleaseFolderFlag    = "release-folder"
	defaultReleaseFolder = "release"
)

// Command release
func (r *ReleaseCmd) Command(sys *typgo.BuildSys) *cli.Command {
	return &cli.Command{
		Name:  "release",
		Usage: "Release the project",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: ForceFlag, Usage: "Release by passed all validation"},
			&cli.BoolFlag{Name: AlphaFlag, Usage: "Release for alpha version"},
			&cli.StringFlag{Name: TagNameFlag, Usage: "Override the release-tag"},
			&cli.StringFlag{Name: ReleaseFolderFlag, Usage: "release folder", Value: defaultReleaseFolder},
		},
		Before: sys.Action(r.Before),
		Action: sys.Action(r.Action),
	}
}

//
// ReleaseProject
//

var _ typgo.Action = (*ReleaseProject)(nil)

func (r *ReleaseProject) validate() error {
	if r.Summarizer == nil {
		return errors.New("typrls: missing summary")
	}
	if r.Tagger == nil {
		return errors.New("typrls: missing tag")
	}
	if r.Releaser == nil {
		return errors.New("typrls: missing releaser")
	}
	return nil
}

// Execute release
func (r *ReleaseProject) Execute(c *typgo.Context) error {
	if err := r.validate(); err != nil {
		return err
	}

	ctx := c.Ctx()

	gitFetch(ctx)
	defer gitFetch(ctx)

	currentTag := gitTag(ctx)
	alpha := c.Bool(AlphaFlag)
	tagName := c.String(TagNameFlag)
	releaseFolder := c.String(ReleaseFolderFlag)

	if tagName == "" {
		tagName = r.Tagger.CreateTag(c, alpha)
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
		ReleaseFolder: releaseFolder,
	}

	if r.Validator != nil && !c.Bool(ForceFlag) {
		if err := r.Validator.Validate(rlsCtx); err != nil {
			return err
		}
	}

	summary, err := r.Summarizer.Summarize(rlsCtx)
	if err != nil {
		return err
	}
	rlsCtx.Summary = summary

	os.RemoveAll(releaseFolder)
	os.MkdirAll(releaseFolder, 0777)

	return r.Releaser.Release(rlsCtx)
}
