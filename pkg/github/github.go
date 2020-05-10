package github

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"github.com/typical-go/typical-go/pkg/git"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typvar"
	"golang.org/x/oauth2"
)

var (
	_ MessageFilter = (*GithubModule)(nil)
)

// GithubModule publisher
type GithubModule struct {
	owner    string
	repoName string
	filter   MessageFilter
}

// Github module
func Github(owner, repo string) *GithubModule {
	return &GithubModule{
		owner:    owner,
		repoName: repo,
		filter:   DefaultNoPrefix(),
	}
}

// Filter define filter message
func (g *GithubModule) Filter(filter MessageFilter) *GithubModule {
	g.filter = filter
	return g
}

// Publish to github
func (g *GithubModule) Publish(c *typgo.PublishContext) (err error) {
	token := os.Getenv("GITHUB_TOKEN")

	if token == "" {
		return errors.New("Environment 'GITHUB_TOKEN' is missing")
	}

	ctx := c.Cli.Context
	oauth := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}))
	repo := github.NewClient(oauth).Repositories
	if _, _, err = repo.GetReleaseByTag(ctx, g.owner, g.repoName, c.Tag); err == nil {
		return fmt.Errorf("Tag '%s' already published", c.Tag)
	}
	c.Infof("Create github release for %s/%s", g.owner, g.repoName)
	githubRls := &github.RepositoryRelease{
		Name:       github.String(fmt.Sprintf("%s - %s", c.BuildTool.Name, c.Tag)),
		TagName:    github.String(c.Tag),
		Body:       github.String(g.releaseNote(c.GitLogs)),
		Draft:      github.Bool(false),
		Prerelease: github.Bool(c.Alpha),
	}
	if githubRls, _, err = repo.CreateRelease(ctx, g.owner, g.repoName, githubRls); err != nil {
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

func (g *GithubModule) upload(ctx context.Context, svc *github.RepositoriesService, id int64, binary string) (err error) {
	var (
		file       *os.File
		binaryPath = fmt.Sprintf("%s/%s", typvar.ReleaseFolder, binary)
	)
	if file, err = os.Open(binaryPath); err != nil {
		return
	}
	defer file.Close()
	opts := &github.UploadOptions{Name: binary}
	_, _, err = svc.UploadReleaseAsset(ctx, g.owner, g.repoName, id, opts, file)
	return
}

func (g *GithubModule) releaseNote(gitLogs []*git.Log) string {
	var b strings.Builder
	for _, log := range gitLogs {
		if m := g.MessageFilter(log.Message); m != "" {
			b.WriteString(log.Short)
			b.WriteString(" ")
			b.WriteString(log.Message)
			b.WriteString("\n")
		}
	}
	return b.String()
}

// MessageFilter to filter the message
func (g *GithubModule) MessageFilter(msg string) string {
	if g.filter != nil {
		return g.filter.MessageFilter(msg)
	}
	return msg
}
