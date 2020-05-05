package typdocker

import (
	"os"
	"os/exec"

	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/urfave/cli/v2"
)

func (m *DockerUtility) cmdDown(c *typbuildtool.Context) *cli.Command {
	return &cli.Command{
		Name:    "down",
		Aliases: []string{"stop"},
		Usage:   "Take down all docker containers according docker-compose",
		Action:  c.ActionFunc(LogName, dockerDown),
	}
}

func dockerDown(c *typbuildtool.CliContext) error {
	c.Info("Docker down")
	cmd := exec.CommandContext(c.Context, "docker-compose", "down")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
