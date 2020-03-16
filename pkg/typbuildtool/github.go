package typbuildtool

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/typical-go/typical-go/pkg/git"

	"github.com/google/go-github/github"

	"golang.org/x/oauth2"
)

// Github publisher
type Github struct {
	Owner    string
	RepoName string
	Filter   ReleaseFilter
}

// NewGithub to return new instance of Github
func NewGithub(owner, repo string) *Github {
	return &Github{
		Owner:    owner,
		RepoName: repo,
		Filter:   DefaultNoPrefix(),
	}
}

// WithFilter return github with filter
func (g *Github) WithFilter(filter ReleaseFilter) *Github {
	g.Filter = filter
	return g
}

// Publish to github
func (g *Github) Publish(c *PublishContext) (err error) {
	token := os.Getenv("GITHUB_TOKEN")
	ctx := c.Cli.Context
	if token == "" {
		return errors.New("Environment 'GITHUB_TOKEN' is missing")
	}
	repo := github.NewClient(oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}))).Repositories
	if _, _, err = repo.GetReleaseByTag(ctx, g.Owner, g.RepoName, c.Tag); err == nil {
		return fmt.Errorf("Tag '%s' already published", c.Tag)
	}
	c.Infof("Create github release for %s/%s", g.Owner, g.RepoName)
	githubRls := &github.RepositoryRelease{
		Name:       github.String(fmt.Sprintf("%s - %s", c.Name, c.Tag)),
		TagName:    github.String(c.Tag),
		Body:       github.String(g.releaseNote(c.GitLogs)),
		Draft:      github.Bool(false),
		Prerelease: github.Bool(c.Alpha),
	}
	if githubRls, _, err = repo.CreateRelease(ctx, g.Owner, g.RepoName, githubRls); err != nil {
		return
	}
	for _, file := range c.ReleaseFiles {
		c.Infof("Upload asset: %s", file)
		if err = g.upload(ctx, repo, *githubRls.ID, file); err != nil {
			return
		}
	}
	return
}

func (g *Github) upload(ctx context.Context, svc *github.RepositoriesService, id int64, binary string) (err error) {
	var (
		file       *os.File
		binaryPath = fmt.Sprintf("%s/%s", DefaultReleaseFolder, binary)
	)
	if file, err = os.Open(binaryPath); err != nil {
		return
	}
	defer file.Close()
	opts := &github.UploadOptions{Name: binary}
	_, _, err = svc.UploadReleaseAsset(ctx, g.Owner, g.RepoName, id, opts, file)
	return
}

func (g *Github) releaseNote(gitLogs []*git.Log) string {
	var b strings.Builder
	for _, log := range gitLogs {
		if m := g.ReleaseFilter(log.Message); m != "" {
			b.WriteString(log.Short)
			b.WriteString(" ")
			b.WriteString(log.Message)
			b.WriteString("\n")
		}
	}
	return b.String()
}

// ReleaseFilter to filter the message
func (g *Github) ReleaseFilter(msg string) string {
	return g.Filter.ReleaseFilter(msg)
}
