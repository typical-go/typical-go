package typrls

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/github"
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typenv"

	"golang.org/x/oauth2"
)

// Github publisher
type Github struct {
	Owner    string
	RepoName string
}

// GithubPublisher to return new instance of Github
func GithubPublisher(owner, repo string) *Github {
	return &Github{
		Owner:    owner,
		RepoName: repo,
	}
}

// Publish to github
func (g *Github) Publish(ctx context.Context, name, tag string, changeLogs, binaries []string, alpha bool) (err error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return errors.New("Environment 'GITHUB_TOKEN' is missing")
	}
	repo := github.NewClient(oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}))).Repositories
	if _, _, err = repo.GetReleaseByTag(ctx, g.Owner, g.RepoName, tag); err == nil {
		return fmt.Errorf("Tag '%s' already published", tag)
	}
	log.Infof("Create github release for %s/%s", g.Owner, g.RepoName)
	githubRls := &github.RepositoryRelease{
		Name:       github.String(fmt.Sprintf("%s - %s", name, tag)),
		TagName:    github.String(tag),
		Body:       github.String(releaseNote(changeLogs)),
		Draft:      github.Bool(false),
		Prerelease: github.Bool(alpha),
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

func (g *Github) upload(ctx0 context.Context, service *github.RepositoriesService, id int64, binary string) (err error) {
	binaryPath := fmt.Sprintf("%s/%s", typenv.Layout.Release, binary)
	var file *os.File
	if file, err = os.Open(binaryPath); err != nil {
		return
	}
	defer file.Close()
	opts := &github.UploadOptions{Name: binary}
	_, _, err = service.UploadReleaseAsset(ctx0, g.Owner, g.RepoName, id, opts, file)
	return
}

func releaseNote(changeLogs []string) string {
	var b strings.Builder
	for _, changelog := range changeLogs {
		b.WriteString(changelog)
		b.WriteString("\n")
	}
	return b.String()
}
