package typdocker

import (
	"os"
	"os/exec"

	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/urfave/cli/v2"
)

func (m *DockerUtility) cmdUp(c *typbuildtool.Context) *cli.Command {
	return &cli.Command{
		Name:    "up",
		Aliases: []string{"start"},
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "wipe"},
		},
		Usage:  "Spin up docker containers according docker-compose",
		Action: m.upAction(c),
	}
}

func (m *DockerUtility) upAction(c *typbuildtool.Context) cli.ActionFunc {
	return func(cc *cli.Context) (err error) {
		if cc.Bool("wipe") {
			m.wipeAction(c)(cc)
		}

		if _, err = os.Stat(dockerComposeFile); os.IsNotExist(err) {
			if err = m.composeAction(c)(cc); err != nil {
				return
			}
		}

		c.Info("Docker up")
		cmd := exec.CommandContext(cc.Context, "docker-compose", "up", "--remove-orphans", "-d")
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		return cmd.Run()
	}
}
