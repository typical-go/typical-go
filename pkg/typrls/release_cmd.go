package typrls

import (
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
		Publisher  Publisher
	}
)

var _ typgo.Cmd = (*ReleaseCmd)(nil)

const (
	// ForceFlag ..
	ForceFlag = "force"
	// AlphaFlag ...
	AlphaFlag = "alpha"
	// TagNameFlag ...
	TagNameFlag = "tag-name"
	// SkipPublishFlag ...
	SkipPublishFlag = "skip-publish"
	// SkipReleaseFlag ...
	SkipReleaseFlag = "skip-release"
	// ReleaseFolderFlag ...
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
			&cli.BoolFlag{Name: SkipPublishFlag, Usage: "Skip publish"},
			&cli.BoolFlag{Name: SkipReleaseFlag, Usage: "Skip release"},
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
		r.Summarizer = DefaultSummarizer
	}
	if r.Tagger == nil {
		r.Tagger = DefaultTagger
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

	context := &Context{
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
		if err := r.Validator.Validate(context); err != nil {
			return err
		}
	}

	summary, err := r.Summarizer.Summarize(context)
	if err != nil {
		return err
	}
	context.Summary = summary

	if r.Releaser != nil && !c.Bool(SkipReleaseFlag) {
		os.RemoveAll(releaseFolder)
		os.MkdirAll(releaseFolder, 0777)
		if err := r.Releaser.Release(context); err != nil {
			return err
		}
	}

	if r.Publisher != nil && !c.Bool(SkipPublishFlag) {
		if err := r.Publisher.Publish(context); err != nil {
			return err
		}
	}
	return nil
}
