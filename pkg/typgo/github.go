package typgo

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"github.com/typical-go/typical-go/pkg/git"
	"github.com/typical-go/typical-go/pkg/typvar"
	"golang.org/x/oauth2"
)

var (
	_ Excluder = (*Github)(nil)
)

type (
	// Github to publish
	Github struct {
		Owner    string
		RepoName string
		PublishSetting
	}

	// PublishSetting contain setting for setting
	PublishSetting struct {
		ExcludeMessage Excluder
	}
)

// Publish to github
func (g *Github) Publish(c *PublishContext) (err error) {
	token := os.Getenv("GITHUB_TOKEN")

	if token == "" {
		return errors.New("Environment 'GITHUB_TOKEN' is missing")
	}

	ctx := c.Cli.Context
	oauth := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}))
	repo := github.NewClient(oauth).Repositories
	if _, _, err = repo.GetReleaseByTag(ctx, g.Owner, g.RepoName, c.Tag); err == nil {
		return fmt.Errorf("Tag '%s' already published", c.Tag)
	}
	c.Infof("Create github release for %s/%s", g.Owner, g.RepoName)
	githubRls := &github.RepositoryRelease{
		Name:       github.String(fmt.Sprintf("%s - %s", c.BuildTool.Name, c.Tag)),
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
		binaryPath = fmt.Sprintf("%s/%s", typvar.ReleaseFolder, binary)
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
		if !g.Exclude(log.Message) {
			b.WriteString(log.Short)
			b.WriteString(" ")
			b.WriteString(log.Message)
			b.WriteString("\n")
		}
	}
	return b.String()
}

// Exclude message
func (g *PublishSetting) Exclude(msg string) bool {
	if g.ExcludeMessage != nil {
		return g.ExcludeMessage.Exclude(msg)
	}
	return false
}
