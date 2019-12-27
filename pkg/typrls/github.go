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
func (g *Github) Publish(r *Release) (err error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return errors.New("Environment 'GITHUB_TOKEN' is missing")
	}

	ctx0 := context.Background()
	repo := github.NewClient(oauth2.NewClient(ctx0, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}))).Repositories
	if _, _, err = repo.GetReleaseByTag(ctx0, g.Owner, g.RepoName, r.Tag); err == nil {
		return fmt.Errorf("Tag '%s' already published", r.Tag)
	}
	log.Infof("Create github release for %s/%s", g.Owner, g.RepoName)
	githubRls := &github.RepositoryRelease{
		Name:       github.String(fmt.Sprintf("%s - %s", r.Name, r.Tag)),
		TagName:    github.String(r.Tag),
		Body:       github.String(releaseNote(r.ChangeLogs)),
		Draft:      github.Bool(false),
		Prerelease: github.Bool(r.Alpha),
	}
	if githubRls, _, err = repo.CreateRelease(ctx0, g.Owner, g.RepoName, githubRls); err != nil {
		return
	}
	for _, binary := range r.Binaries {
		log.Infof("Upload asset: %s", binary)
		if err = g.upload(ctx0, repo, *githubRls.ID, binary); err != nil {
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
