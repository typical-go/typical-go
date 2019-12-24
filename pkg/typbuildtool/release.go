package typbuildtool

import (
	log "github.com/sirupsen/logrus"
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

func (t buildtool) releaseDistribution(ctx *cli.Context) (err error) {
	log.Info("Release distribution")
	if !ctx.Bool("no-test") {
		if err = t.runTesting(ctx); err != nil {
			return
		}
	}
	log.Info("Release the distribution")
	if err = t.Releaser.Release(
		typenv.ProjectName,
		t.Version,
		ctx.Bool("force"),
		ctx.Bool("alpha"),
		ctx.Bool("no-publish"),
	); err != nil {
		return
	}

	return
}
