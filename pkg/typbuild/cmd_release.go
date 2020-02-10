package typbuild

import (
	"errors"
	"fmt"

	"github.com/typical-go/typical-go/pkg/git"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/urfave/cli/v2"
)

func (b *Build) cmdRelease(bctx *typcore.BuildContext) *cli.Command {
	return &cli.Command{
		Name:  "release",
		Usage: "Release the distribution",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "no-test", Usage: "Release without run unit test"},
			&cli.BoolFlag{Name: "no-build", Usage: "Release without build"},
			&cli.BoolFlag{Name: "no-publish", Usage: "Release without create github release"},
			&cli.BoolFlag{Name: "force", Usage: "Release by passed all validation"},
			&cli.BoolFlag{Name: "alpha", Usage: "Release for alpha version"},
		},
		Action: func(c *cli.Context) (err error) {
			if b.releaser == nil {
				return errors.New("Releaser is missing")
			}
			var (
				tag        string
				latest     string
				changeLogs []string
				binaries   []string
				name       = typenv.ProjectName
				alpha      = c.Bool("alpha")
				force      = c.Bool("force")
				noTest     = c.Bool("no-test")
				noBuild    = c.Bool("no-build")
				noPublish  = c.Bool("no-publish")
				ctx        = c.Context
			)
			if !noBuild {
				if err = b.buildProject(ctx, bctx); err != nil {
					return
				}
			}
			if !noTest {
				if err = runTesting(c); err != nil {
					return
				}
			}
			if err = git.Fetch(ctx); err != nil {
				return fmt.Errorf("Failed git fetch: %w", err)
			}
			defer git.Fetch(ctx)
			if tag, err = b.releaser.Tag(ctx, bctx.Version, alpha); err != nil {
				return fmt.Errorf("Failed generate tag: %w", err)
			}
			if status := git.Status(ctx); status != "" && !force {
				return fmt.Errorf("Please commit changes first:\n%s", status)
			}
			if latest = git.LatestTag(ctx); latest == tag && !force {
				return fmt.Errorf("%s already released", latest)
			}
			if changeLogs = git.Logs(ctx, latest); len(changeLogs) < 1 && !force {
				return errors.New("No change to be released")
			}
			if binaries, err = b.releaser.BuildRelease(ctx, name, tag, changeLogs, alpha); err != nil {
				return fmt.Errorf("Failed build release: %w", err)
			}
			if !noPublish {
				if err = b.releaser.Publish(ctx, name, tag, changeLogs, binaries, alpha); err != nil {
					return fmt.Errorf("Failed publish: %w", err)
				}
			}
			return
		},
	}
}
