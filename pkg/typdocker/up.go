package typdocker

import (
	"os"
	"os/exec"

	"github.com/typical-go/typical-go/pkg/typbuild"
	"github.com/urfave/cli/v2"
)

func (m *DockerUtility) cmdUp(c *typbuild.Context) *cli.Command {
	return &cli.Command{
		Name:    "up",
		Aliases: []string{"start"},
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "wipe"},
		},
		Usage:  "Spin up docker containers according docker-compose",
		Action: c.ActionFunc(LogName, m.dockerUp),
	}
}

func (m *DockerUtility) dockerUp(c *typbuild.CliContext) (err error) {

	if c.Cli.Bool("wipe") {
		m.dockerWipe(c)
	}

	if _, err = os.Stat(DockerComposeFile); os.IsNotExist(err) {
		if err = m.dockerCompose(c); err != nil {
			return
		}
	}

	c.Info("Docker up")
	cmd := exec.CommandContext(c.Context, "docker-compose", "up", "--remove-orphans", "-d")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()

}
