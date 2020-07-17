package typrls

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"github.com/typical-go/typical-go/pkg/git"
	"golang.org/x/oauth2"
)

type (
	// Github to publish
	Github struct {
		Owner string
		Repo  string
	}
)

var _ Releaser = (*Github)(nil)

// Release to github
func (g *Github) Release(c *Context) (err error) {
	token := os.Getenv("GITHUB_TOKEN")

	if token == "" {
		return errors.New("Environment 'GITHUB_TOKEN' is missing")
	}

	ctx := c.Ctx()
	oauth := oauth2.NewClient(ctx,
		oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		),
	)
	repo := github.NewClient(oauth).Repositories
	tag := c.Tag
	gitLogs := c.GitLogs
	alpha := c.Alpha

	if _, _, err = repo.GetReleaseByTag(ctx, g.Owner, g.Repo, tag); err == nil {
		return fmt.Errorf("Tag '%s' already published", tag)
	}
	fmt.Printf("\nCreate github release for %s/%s\n", g.Owner, g.Repo)
	githubRls := &github.RepositoryRelease{
		Name:       github.String(fmt.Sprintf("%s - %s", c.Descriptor.Name, tag)),
		TagName:    github.String(tag),
		Body:       github.String(g.releaseNote(gitLogs)),
		Draft:      github.Bool(false),
		Prerelease: github.Bool(alpha),
	}
	if githubRls, _, err = repo.CreateRelease(ctx, g.Owner, g.Repo, githubRls); err != nil {
		return
	}

	return
}

func (g *Github) releaseNote(gitLogs []*git.Log) string {
	var b strings.Builder
	for _, log := range gitLogs {
		if !ExcludeMessage(log.Message) {
			b.WriteString(log.Short)
			b.WriteString(" ")
			b.WriteString(log.Message)
			b.WriteString("\n")
		}
	}
	return b.String()
}

// ExcludeMessage return true is message mean to be exclude
func ExcludeMessage(msg string) bool {
	msg = strings.ToLower(msg)
	for _, prefix := range ExclMsgPrefix {
		if strings.HasPrefix(msg, strings.ToLower(prefix)) {
			return true
		}
	}
	return false
}
