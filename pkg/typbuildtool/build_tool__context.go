package typbuildtool

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/typical-go/typical-go/pkg/git"
	"github.com/urfave/cli/v2"
)

// createContext to create new instance of Context
func (b *TypicalBuildTool) createContext(tc *Context, cc *cli.Context) *BuildContext {
	return &BuildContext{
		Context: tc,
		Cli:     cc,
	}
}

func (b *TypicalBuildTool) createReleaseContext(c *BuildContext) (*ReleaseContext, error) {
	ctx := c.Cli.Context
	force := c.Cli.Bool("force")
	alpha := c.Cli.Bool("alpha")
	tag := b.releaseTag(ctx, c.Version, alpha)

	if status := git.Status(ctx); status != "" && !force {
		return nil, fmt.Errorf("Please commit changes first:\n%s", status)
	}

	var latest string
	if latest = git.LatestTag(ctx); latest == tag && !force {
		return nil, fmt.Errorf("%s already released", latest)
	}

	var gitLogs []*git.Log
	if gitLogs = git.RetrieveLogs(ctx, latest); len(gitLogs) < 1 && !force {
		return nil, errors.New("No change to be released")
	}

	return &ReleaseContext{
		BuildContext: c,
		Alpha:        alpha,
		Tag:          tag,
		GitLogs:      gitLogs,
	}, nil
}

// Tag return relase tag
func (b *TypicalBuildTool) releaseTag(ctx context.Context, version string, alpha bool) string {
	var builder strings.Builder
	builder.WriteString("v")
	builder.WriteString(version)
	if b.includeBranch {
		builder.WriteString("_")
		builder.WriteString(git.Branch(ctx))
	}
	if b.includeCommitID {
		builder.WriteString("_")
		builder.WriteString(git.LatestCommit(ctx))
	}
	if alpha {
		builder.WriteString("_alpha")
	}
	return builder.String()
}
