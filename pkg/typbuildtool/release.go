package typbuildtool

import (
	"errors"
	"fmt"

	"github.com/typical-go/typical-go/pkg/git"
	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/urfave/cli/v2"
)

func (t buildtool) cmdRelease() *cli.Command {
	return &cli.Command{
		Name:  "release",
		Usage: "Release the distribution",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "no-test", Usage: "Release without run automated test"},
			&cli.BoolFlag{Name: "no-publish", Usage: "Release without create github release"},
			&cli.BoolFlag{Name: "force", Usage: "Release by passed all validation"},
			&cli.BoolFlag{Name: "alpha", Usage: "Release for alpha version"},
		},
		Action: t.releaseDistribution,
	}
}

func (t buildtool) releaseDistribution(c *cli.Context) (err error) {
	var (
		tag        string
		latestTag  string
		changeLogs []string
		binaries   []string
		name       = typenv.ProjectName
		alpha      = c.Bool("alpha")
		force      = c.Bool("force")
		noTest     = c.Bool("no-test")
		noPublish  = c.Bool("no-publish")
		ctx        = c.Context
	)
	if !noTest {
		if err = t.runTesting(c); err != nil {
			return
		}
	}
	if err = git.Fetch(); err != nil {
		return fmt.Errorf("Failed git fetch: %w", err)
	}
	defer git.Fetch()
	if tag, err = t.Releaser.Tag(t.Version, alpha); err != nil {
		return fmt.Errorf("Failed generate tag: %w", err)
	}
	if status := git.Status(); status != "" && !force {
		return fmt.Errorf("Please commit changes first:\n%s", status)
	}
	if latestTag = git.LatestTag(); latestTag == tag && !force {
		return fmt.Errorf("%s already released", latestTag)
	}
	if changeLogs = git.Logs(latestTag); len(changeLogs) < 1 && !force {
		return errors.New("No change to be released")
	}
	if binaries, err = t.Releaser.BuildRelease(ctx, name, tag, changeLogs, alpha); err != nil {
		return fmt.Errorf("Failed build release: %w", err)
	}
	if !noPublish {
		if err = t.Releaser.Publish(ctx, name, tag, changeLogs, binaries, alpha); err != nil {
			return fmt.Errorf("Failed publish: %w", err)
		}
	}
	return
}
