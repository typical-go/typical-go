package typrls

import (
	"errors"
	"fmt"
	"os"

	"github.com/google/go-github/github"
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
	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		return errors.New("Environment 'GITHUB_TOKEN' is missing")
	}

	ctx := c.Ctx()
	token := &oauth2.Token{AccessToken: githubToken}
	oauth := oauth2.NewClient(ctx, oauth2.StaticTokenSource(token))
	repo := github.NewClient(oauth).Repositories

	if _, _, err = repo.GetReleaseByTag(ctx, g.Owner, g.Repo, c.TagName); err == nil {
		return fmt.Errorf("Tag '%s' already published", c.TagName)
	}
	fmt.Printf("\nCreate github release for %s/%s\n", g.Owner, g.Repo)
	githubRls := &github.RepositoryRelease{
		Name:       github.String(fmt.Sprintf("%s - %s", c.BuildSys.ProjectName, c.TagName)),
		TagName:    github.String(c.TagName),
		Body:       github.String(c.Summary),
		Draft:      github.Bool(false),
		Prerelease: github.Bool(c.Alpha),
	}
	if githubRls, _, err = repo.CreateRelease(ctx, g.Owner, g.Repo, githubRls); err != nil {
		return
	}

	return
}
