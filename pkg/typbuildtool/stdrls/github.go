package stdrls

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/typical-go/typical-go/pkg/git"
	"github.com/typical-go/typical-go/pkg/typbuild"

	"github.com/google/go-github/github"
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typcore"

	"golang.org/x/oauth2"
)

// Github publisher
type Github struct {
	Owner    string
	RepoName string
	Filter   Filter
}

// GithubPublisher to return new instance of Github
func GithubPublisher(owner, repo string) *Github {
	return &Github{
		Owner:    owner,
		RepoName: repo,
		Filter:   DefaultNoPrefix(),
	}
}

// WithFilter return github with filter
func (g *Github) WithFilter(filter Filter) *Github {
	g.Filter = filter
	return g
}

// Publish to github
func (g *Github) Publish(ctx context.Context, rel *typbuild.ReleaseContext, binaries []string) (err error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return errors.New("Environment 'GITHUB_TOKEN' is missing")
	}
	repo := github.NewClient(oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}))).Repositories
	if _, _, err = repo.GetReleaseByTag(ctx, g.Owner, g.RepoName, rel.Tag); err == nil {
		return fmt.Errorf("Tag '%s' already published", rel.Tag)
	}
	log.Infof("Create github release for %s/%s", g.Owner, g.RepoName)
	githubRls := &github.RepositoryRelease{
		Name:       github.String(fmt.Sprintf("%s - %s", rel.Name, rel.Tag)),
		TagName:    github.String(rel.Tag),
		Body:       github.String(g.releaseNote(rel.GitLogs)),
		Draft:      github.Bool(false),
		Prerelease: github.Bool(rel.Alpha),
	}
	if githubRls, _, err = repo.CreateRelease(ctx, g.Owner, g.RepoName, githubRls); err != nil {
		return
	}
	for _, binary := range binaries {
		log.Infof("Upload asset: %s", binary)
		if err = g.upload(ctx, repo, *githubRls.ID, binary); err != nil {
			return
		}
	}
	return
}

func (g *Github) upload(ctx context.Context, svc *github.RepositoriesService, id int64, binary string) (err error) {
	var (
		file       *os.File
		binaryPath = fmt.Sprintf("%s/%s", typcore.DefaultReleaseFolder, binary)
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
		if m := g.Filter.Filter(log.Message); m != "" {
			b.WriteString(log.Short)
			b.WriteString(" ")
			b.WriteString(log.Message)
			b.WriteString("\n")
		}
	}
	return b.String()
}
