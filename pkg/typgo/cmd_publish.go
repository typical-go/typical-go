package typgo

import (
	"errors"
	"fmt"
	"strings"

	"github.com/typical-go/typical-go/pkg/git"
	"github.com/typical-go/typical-go/pkg/typvar"
	"github.com/urfave/cli/v2"
)

func cmdPublish(c *BuildCli) *cli.Command {
	return &cli.Command{
		Name:    "publish",
		Usage:   "Publish the project",
		Aliases: []string{"p"},
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "no-test", Usage: "skip the test"},
			&cli.BoolFlag{Name: "force", Usage: "Release by passed all validation"},
			&cli.BoolFlag{Name: "alpha", Usage: "Release for alpha version"},
		},
		Action: c.ActionFn("PUBLISH", Publish),
	}
}

// Publish project
func Publish(c *Context) (err error) {

	var (
		latest  string
		gitLogs []*git.Log
	)

	if !c.Bool("no-test") {
		if err = test(c); err != nil {
			return
		}
	}

	ctx := c.Ctx()

	if err = git.Fetch(ctx); err != nil {
		return fmt.Errorf("Failed git fetch: %w", err)
	}
	defer git.Fetch(ctx)

	force := c.Bool("force")
	alpha := c.Bool("alpha")
	tag := releaseTag(c, alpha)

	if status := git.Status(ctx); status != "" && !force {
		return fmt.Errorf("Please commit changes first:\n%s", status)
	}

	if latest = git.LatestTag(ctx); latest == tag && !force {
		return fmt.Errorf("%s already released", latest)
	}

	if gitLogs = git.RetrieveLogs(ctx, latest); len(gitLogs) < 1 && !force {
		return errors.New("No change to be released")
	}

	typvar.Rls.Alpha = alpha
	typvar.Rls.Tag = tag
	typvar.Rls.GitLogs = gitLogs

	if _, err = c.Execute(c, ReleasePhase); err != nil {
		return
	}

	if _, err = c.Execute(c, PublishPhase); err != nil {
		return
	}

	return

}

func releaseTag(c *Context, alpha bool) string {
	var builder strings.Builder
	builder.WriteString("v")
	builder.WriteString(c.Descriptor.Version)
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
