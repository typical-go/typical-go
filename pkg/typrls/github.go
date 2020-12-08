package typrls

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/go-github/v33/github"
	"github.com/typical-go/typical-go/pkg/oskit"
	"golang.org/x/oauth2"
)

type (
	// Github to publish
	Github struct {
		Owner string
		Repo  string
		Draft bool
	}
)

var _ Publisher = (*Github)(nil)

// Publish to github
func (g *Github) Publish(c *Context) (err error) {
	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		return errors.New("github-release: missing $GITHUB_TOKEN")
	}

	token := &oauth2.Token{AccessToken: githubToken}
	oauth := oauth2.NewClient(c.Ctx(), oauth2.StaticTokenSource(token))
	repo := github.NewClient(oauth).Repositories

	if _, _, err = g.getReleaseByTag(c, repo); err == nil {
		return fmt.Errorf("github-release: Can't publish to existing tag '%s'", c.TagName)
	}

	fmt.Printf("\nCreate github release for %s/%s\n", g.Owner, g.Repo)
	rls := &github.RepositoryRelease{
		Name:       github.String(fmt.Sprintf("%s - %s", c.BuildSys.ProjectName, c.TagName)),
		TagName:    github.String(c.TagName),
		Body:       github.String(c.Summary),
		Draft:      github.Bool(g.Draft),
		Prerelease: github.Bool(c.Alpha),
	}
	if rls, _, err = g.createRelease(c, repo, rls); err != nil {
		return
	}

	files, _ := ioutil.ReadDir(c.ReleaseFolder)
	for _, fileInfo := range files {
		path := c.ReleaseFolder + "/" + fileInfo.Name()
		fmt.Fprintf(oskit.Stdout, "Upload '%s'\n", path)
		var file *os.File
		if file, err = os.Open(path); err != nil {
			return
		}
		defer file.Close()

		opt := &github.UploadOptions{Name: filepath.Base(path)}
		if _, _, err := g.uploadReleaseAsset(c, repo, *rls.ID, opt, file); err != nil {
			fmt.Fprintf(oskit.Stdout, "WARN: %s\n", err.Error())
		}
	}

	return
}

func (g *Github) getReleaseByTag(c *Context, repo *github.RepositoriesService) (*github.RepositoryRelease, *github.Response, error) {
	return repo.GetReleaseByTag(c.Ctx(), g.Owner, g.Repo, c.TagName)
}

func (g *Github) createRelease(c *Context, repo *github.RepositoriesService, rls *github.RepositoryRelease) (*github.RepositoryRelease, *github.Response, error) {
	return repo.CreateRelease(c.Ctx(), g.Owner, g.Repo, rls)
}

func (g *Github) uploadReleaseAsset(c *Context, repo *github.RepositoriesService, id int64, opt *github.UploadOptions, file *os.File) (*github.ReleaseAsset, *github.Response, error) {
	return repo.UploadReleaseAsset(c.Ctx(), g.Owner, g.Repo, id, opt, file)
}
