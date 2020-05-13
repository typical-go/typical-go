package typgo

import (
	"errors"
	"fmt"
	"strings"

	"github.com/typical-go/typical-go/pkg/git"
	"github.com/typical-go/typical-go/pkg/typvar"
	"github.com/urfave/cli/v2"
)

func createBuildToolCmds(d *Descriptor) (cmds []*cli.Command) {
	b := &BuildTool{
		Descriptor: d,
	}

	cmds = []*cli.Command{
		cmdTest(b),
		cmdRun(b),
		cmdPublish(b),
		cmdClean(b),
	}

	if d.Utility != nil {
		for _, cmd := range d.Utility.Commands(b) {
			cmds = append(cmds, cmd)
		}
	}

	return cmds
}

func cmdTest(c *BuildTool) *cli.Command {
	return &cli.Command{
		Name:    "test",
		Aliases: []string{"t"},
		Usage:   "Test the project",
		Action:  c.ActionFunc("TEST", test),
	}
}

func test(c *Context) (err error) {
	_, err = c.Execute(c, TestPhase)
	return
}

func cmdRun(c *BuildTool) *cli.Command {
	return &cli.Command{
		Name:            "run",
		Aliases:         []string{"r"},
		Usage:           "Run the project in local environment",
		SkipFlagParsing: true,
		Action:          c.ActionFunc("RUN", run),
	}
}

func run(c *Context) (err error) {
	_, err = c.Execute(c, RunPhase)
	return
}

func cmdPublish(c *BuildTool) *cli.Command {
	return &cli.Command{
		Name:    "publish",
		Usage:   "Publish the project",
		Aliases: []string{"p"},
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "no-test", Usage: "skip the test"},
			&cli.BoolFlag{Name: "force", Usage: "Release by passed all validation"},
			&cli.BoolFlag{Name: "alpha", Usage: "Release for alpha version"},
		},
		Action: c.ActionFunc("PUBLISH", Publish),
	}
}

func cmdClean(c *BuildTool) *cli.Command {
	return &cli.Command{
		Name:    "clean",
		Aliases: []string{"c"},
		Usage:   "Clean the project",
		Action:  c.ActionFunc("CLEAN", clean),
	}
}

func clean(c *Context) (err error) {
	if _, err = c.Execute(c, CleanPhase); err != nil {
		return
	}
	return
}

// Publish project
func Publish(c *Context) (err error) {

	var (
		latest  string
		gitLogs []*git.Log
	)

	if !c.Cli.Bool("no-test") {
		if err = test(c); err != nil {
			return
		}
	}

	ctx := c.Cli.Context

	if err = git.Fetch(ctx); err != nil {
		return fmt.Errorf("Failed git fetch: %w", err)
	}
	defer git.Fetch(ctx)

	force := c.Cli.Bool("force")
	alpha := c.Cli.Bool("alpha")
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
	builder.WriteString(c.BuildTool.Version)
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
