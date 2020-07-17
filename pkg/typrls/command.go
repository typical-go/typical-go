package typrls

import (
	"fmt"
	"strings"

	"github.com/typical-go/typical-go/pkg/git"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

var (
	// ExclMsgPrefix excluded message prefix list
	ExclMsgPrefix = []string{
		"merge", "bump", "revision", "generate", "wip",
	}
	forceParam = "force"
	alphaParam = "alpha"
)

type (
	// Command release command
	Command struct {
		Releaser
		Validator
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
	ctx := c.Ctx()

	if err := git.Fetch(ctx); err != nil {
		return fmt.Errorf("Failed git fetch: %w", err)
	}
	defer git.Fetch(ctx)

	releaseTag := r.releaseTag(c)
	currentTag := git.CurrentTag(ctx)

	rlsContext := &Context{
		Context: c,
		Git: Git{
			Status:     git.Status(ctx),
			CurrentTag: currentTag,
			Logs:       git.RetrieveLogs(ctx, currentTag),
		},
		ReleaseTag: releaseTag,
	}

	if r.Validator != nil && !c.Bool(forceParam) {
		if err := r.Validator.Validate(rlsContext); err != nil {
			return err
		}
	}

	return r.Release(rlsContext)
}

func (r *Command) releaseTag(c *typgo.Context) string {
	alpha := c.Bool(alphaParam)
	version := "0.0.1"
	if c.Descriptor.Version != "" {
		version = c.Descriptor.Version
	}

	var builder strings.Builder
	builder.WriteString("v")
	builder.WriteString(version)
	// if c.BuildTool.IncludeBranch {
	// 	builder.WriteString("_")
	// 	builder.WriteString(git.Branch(c.Context))
	// }
	// if c.BuildTool.IncludeCommitID {
	// 	builder.WriteString("_")
	// 	builder.WriteString(git.LatestCommit(c.Context))
	// }
	if alpha {
		builder.WriteString("_alpha")
	}
	return builder.String()
}
