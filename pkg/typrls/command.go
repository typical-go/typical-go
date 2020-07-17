package typrls

import (
	"errors"
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
)

type (
	// Command release command
	Command struct {
		Releaser
	}
)

//
// ReleaseCmd
//

var _ typgo.Cmd = (*Command)(nil)
var _ typgo.Action = (*Command)(nil)

// Command release
func (r *Command) Command(c *typgo.BuildCli) *cli.Command {
	return &cli.Command{
		Name:  "release",
		Usage: "Release the project",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "force", Usage: "Release by passed all validation"},
			&cli.BoolFlag{Name: "alpha", Usage: "Release for alpha version"},
		},
		Action: c.ActionFn(r.Execute),
	}
}

// Execute release
func (r *Command) Execute(c *typgo.Context) (err error) {
	ctx := c.Ctx()

	if err = git.Fetch(ctx); err != nil {
		return fmt.Errorf("Failed git fetch: %w", err)
	}
	defer git.Fetch(ctx)

	force := c.Bool("force")
	alpha := c.Bool("alpha")
	tag := releaseTag(c, alpha)

	status := git.Status(ctx)
	if status != "" && !force {
		return fmt.Errorf("Please commit changes first:\n%s", status)
	}

	latest := git.LatestTag(ctx)
	if latest == tag && !force {
		return fmt.Errorf("%s already released", latest)
	}

	gitLogs := git.RetrieveLogs(ctx, latest)
	if len(gitLogs) < 1 && !force {
		return errors.New("No change to be released")
	}

	return r.Release(&Context{
		Context: c,
		Alpha:   alpha,
		Tag:     tag,
		GitLogs: gitLogs,
	})
}

//
// Command
//

func releaseTag(c *typgo.Context, alpha bool) string {
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
