package typrls

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

const (
	// AlphaFlag ...
	AlphaFlag = "alpha"
	// TagNameFlag ...
	TagNameFlag = "tag-name"
	// SkipPublishFlag ...
	SkipPublishFlag = "skip-publish"
	// ReleaseFolderFlag ...
	ReleaseFolderFlag    = "release-folder"
	defaultReleaseFolder = "release"
)

type (
	// ReleaseTool release command
	ReleaseTool struct {
		Before     typgo.Action
		Tagger     Tagger
		Summarizer Summarizer
		Releaser   Releaser
		Publisher  Publisher
	}
)

var _ typgo.Tasker = (*ReleaseTool)(nil)

// Task to release
func (r *ReleaseTool) Task(sys *typgo.BuildSys) *cli.Command {
	return &cli.Command{
		Name:  "release",
		Usage: "Release the project",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: AlphaFlag, Usage: "Release for alpha version"},
			&cli.StringFlag{Name: TagNameFlag, Usage: "Override the release-tag"},
			&cli.BoolFlag{Name: SkipPublishFlag, Usage: "Skip publish"},
			&cli.StringFlag{Name: ReleaseFolderFlag, Usage: "release folder", Value: defaultReleaseFolder},
		},
		Before: sys.Action(r.Before),
		Action: sys.Action(r),
	}
}

//
// ReleaseTool
//

var _ typgo.Action = (*ReleaseTool)(nil)

// Execute release
func (r *ReleaseTool) Execute(c *typgo.Context) error {
	r.setDefault()

	gitFetch(c)
	defer gitFetch(c)

	alpha := c.Bool(AlphaFlag)
	tagName := c.String(TagNameFlag)
	if tagName == "" {
		tagName = r.Tagger.CreateTag(c, alpha)
	}

	context := &Context{
		Context:       c,
		Alpha:         alpha,
		TagName:       tagName,
		ReleaseFolder: c.String(ReleaseFolderFlag),
		Summary:       r.Summarizer.Summarize(c),
	}

	os.RemoveAll(context.ReleaseFolder)
	os.MkdirAll(context.ReleaseFolder, 0777)

	if r.Releaser != nil {
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

func (r *ReleaseTool) setDefault() {
	if r.Summarizer == nil {
		r.Summarizer = DefaultSummarizer
	}
	if r.Tagger == nil {
		r.Tagger = DefaultTagger
	}
}
