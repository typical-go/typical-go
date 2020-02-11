package typbuild

import (
	"errors"
	"fmt"

	"github.com/typical-go/typical-go/pkg/git"
	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/urfave/cli/v2"
)

func (b *Build) cmdRelease(c *Context) *cli.Command {
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
		Action: func(cliCtx *cli.Context) (err error) {
			if b.releaser == nil {
				return errors.New("Releaser is missing")
			}
			var (
				tag        string
				latest     string
				changeLogs []string
				binaries   []string
				name       = typenv.ProjectName
				alpha      = cliCtx.Bool("alpha")
				force      = cliCtx.Bool("force")
				noTest     = cliCtx.Bool("no-test")
				noBuild    = cliCtx.Bool("no-build")
				noPublish  = cliCtx.Bool("no-publish")
				ctx        = cliCtx.Context
			)
			if !noBuild {
				if err = b.buildProject(ctx, c); err != nil {
					return
				}
			}
			if !noTest {
				if err = runTesting(cliCtx); err != nil {
					return
				}
			}
			if err = git.Fetch(ctx); err != nil {
				return fmt.Errorf("Failed git fetch: %w", err)
			}
			defer git.Fetch(ctx)
			if tag, err = b.releaser.Tag(ctx, c.Version, alpha); err != nil {
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
