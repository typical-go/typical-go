package typrls

import (
	"fmt"
	"os"
	"strings"

	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

type (
	// ReleaseProject release command
	ReleaseProject struct {
		Before            typgo.Action
		GenerateTagFn     func(c *typgo.Context, alpha bool) string
		GenerateSummaryFn func(c *typgo.Context) string
		Releaser          Releaser
		Publisher         Publisher
	}
)

const (
	// AlphaFlag alpha flag
	AlphaFlag = "alpha"
	// TagNameFlag tag name flag
	TagNameFlag = "tag-name"
	// SkipPublishFlag skip publish flag
	SkipPublishFlag = "skip-publish"
	// ReleaseFolderFlag release folder flag
	ReleaseFolderFlag    = "release-folder"
	defaultReleaseFolder = "release"
)

// DefaultPrefixes ...
var DefaultPrefixes = []string{"merge", "bump", "revision", "generate", "wip"}

var _ typgo.Tasker = (*ReleaseProject)(nil)

// Task to release
func (r *ReleaseProject) Task() *typgo.Task {
	return &typgo.Task{
		Name:  "release",
		Usage: "Release the project",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: AlphaFlag, Usage: "Release for alpha version"},
			&cli.StringFlag{Name: TagNameFlag, Usage: "Override the release-tag"},
			&cli.BoolFlag{Name: SkipPublishFlag, Usage: "Skip publish"},
			&cli.StringFlag{Name: ReleaseFolderFlag, Usage: "release folder", Value: defaultReleaseFolder},
		},
		Before: r.Before,
		Action: r,
	}
}

//
// ReleaseTool
//

var _ typgo.Action = (*ReleaseProject)(nil)

// Execute release
func (r *ReleaseProject) Execute(c *typgo.Context) error {
	r.setDefault()

	GitFetch(c)
	defer GitFetch(c)

	alpha := c.Bool(AlphaFlag)
	tagName := c.String(TagNameFlag)
	if tagName == "" {
		tagName = r.GenerateTagFn(c, alpha)
	}

	context := &Context{
		Context:       c,
		Alpha:         alpha,
		TagName:       tagName,
		ReleaseFolder: c.String(ReleaseFolderFlag),
		Summary:       r.GenerateSummaryFn(c),
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

func (r *ReleaseProject) setDefault() {
	if r.GenerateSummaryFn == nil {
		r.GenerateSummaryFn = DefaultGenerateSummary
	}
	if r.GenerateTagFn == nil {
		r.GenerateTagFn = DefaultGenerateTag
	}
}

// DefaultGenerateTag create release tag
func DefaultGenerateTag(c *typgo.Context, alpha bool) string {
	tagName := "v0.0.1"
	if c.Descriptor.ProjectVersion != "" {
		tagName = fmt.Sprintf("v%s", c.Descriptor.ProjectVersion)
	}
	if alpha {
		tagName = tagName + "_alpha"
	}
	return tagName
}

// DefaultGenerateSummary default generate summary
func DefaultGenerateSummary(c *typgo.Context) string {
	var changes []string
	currentTag := GitCurrentTag(c)
	for _, log := range GitLogs(c, currentTag) {
		if !HasPrefix(log.Message, DefaultPrefixes) {
			changes = append(changes, fmt.Sprintf("%s %s", log.ShortCode, log.Message))
		}
	}
	return strings.Join(changes, "\n")
}

// HasPrefix return true if eligible to excluded by prefix
func HasPrefix(msg string, prefixes []string) bool {
	msg = strings.ToLower(msg)
	for _, prefix := range prefixes {
		if strings.HasPrefix(msg, strings.ToLower(prefix)) {
			return true
		}
	}
	return false
}
