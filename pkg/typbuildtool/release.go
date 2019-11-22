package typbuildtool

import (
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typrls"
	"github.com/urfave/cli"
)

func (t buildtool) cmdRelease() cli.Command {
	return cli.Command{
		Name:  "release",
		Usage: "Release the distribution",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "no-test",
				Usage: "Release without run automated test",
			},
			cli.BoolFlag{
				Name:  "no-publish",
				Usage: "Release without create github release",
			},
			cli.BoolFlag{
				Name:  "force",
				Usage: "Release by passed all validation",
			},
			cli.BoolFlag{
				Name:  "alpha",
				Usage: "Release for alpha version",
			},
		},
		Action: t.releaseDistribution,
	}
}

func (t buildtool) releaseDistribution(ctx *cli.Context) (err error) {
	var rls *typrls.Release
	log.Info("Release distribution")
	if !ctx.Bool("no-test") {
		if err = t.runTesting(ctx); err != nil {
			return
		}
	}
	log.Info("Release the distribution")
	if rls, err = t.Release(t.Version, ctx.Bool("force"), ctx.Bool("alpha")); err != nil {
		return
	}
	if !ctx.Bool("no-publish") {
		for _, publisher := range t.Publishers {
			log.Info("Publish the distribution")
			if err = publisher.Publish(rls); err != nil {
				return
			}
		}
	}
	return
}
