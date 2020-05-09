package typdocker

import (
	"os"
	"os/exec"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

func (m *DockerUtility) cmdDown(c *typcore.Context) *cli.Command {
	return &cli.Command{
		Name:    "down",
		Aliases: []string{"stop"},
		Usage:   "Take down all docker containers according docker-compose",
		Action:  c.ActionFunc(LogName, dockerDown),
	}
}

func dockerDown(c *typcore.CliContext) error {
	c.Info("Docker down")
	cmd := exec.CommandContext(c.Cli.Context, "docker-compose", "down")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
