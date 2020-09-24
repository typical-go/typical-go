package typrls

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

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

type (
	// ReleaseProject release command
	ReleaseProject struct {
		Before     typgo.Action
		Validator  Validator
		Tagger     Tagger
		Summarizer Summarizer
		Releaser   Releaser
		Publisher  Publisher
	}
)

var _ typgo.Cmd = (*ReleaseProject)(nil)

// Command release
func (r *ReleaseProject) Command(sys *typgo.BuildSys) *cli.Command {
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
		Action: sys.Action(r),
	}
}

//
// ReleaseProject
//

var _ typgo.Action = (*ReleaseProject)(nil)

// Execute release
func (r *ReleaseProject) Execute(c *typgo.Context) error {
	r.setDefault()

	gitFetch(c.Ctx())
	defer gitFetch(c.Ctx())

	context, err := r.context(c)
	if err != nil {
		return err
	}

	if r.Validator != nil && !c.Bool(ForceFlag) {
		if err := r.Validator.Validate(context); err != nil {
			return err
		}
	}

	os.RemoveAll(context.ReleaseFolder)
	os.MkdirAll(context.ReleaseFolder, 0777)

	if r.Releaser != nil && !c.Bool(SkipReleaseFlag) {
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

func (r *ReleaseProject) context(c *typgo.Context) (*Context, error) {
	ctx := c.Ctx()
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

	context.Summary = r.Summarizer.Summarize(context)
	return context, nil
}

func (r *ReleaseProject) setDefault() {
	if r.Summarizer == nil {
		r.Summarizer = DefaultSummarizer
	}
	if r.Tagger == nil {
		r.Tagger = DefaultTagger
	}
}
