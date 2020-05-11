package typdocker

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func (m *DockerUtility) cmdWipe(c *typgo.BuildTool) *cli.Command {
	return &cli.Command{
		Name:   "wipe",
		Usage:  "Kill all running docker container",
		Action: c.ActionFunc(LogName, m.dockerWipe),
	}
}

func (m *DockerUtility) dockerWipe(c *typgo.Context) (err error) {
	var ids []string
	if ids, err = dockerIDs(c.Cli.Context); err != nil {
		return fmt.Errorf("Docker-ID: %w", err)
	}
	for _, id := range ids {
		if err = kill(c.Cli.Context, id); err != nil {
			c.Warnf("Fail to kill #%s: %s", id, err.Error())
		}
	}
	return nil
}

func dockerIDs(ctx context.Context) (ids []string, err error) {
	var out strings.Builder
	cmd := execkit.Command{
		Name:   "docker",
		Args:   []string{"ps", "-q"},
		Stderr: os.Stderr,
		Stdout: &out,
	}

	cmd.Print(os.Stdout)
	if err = cmd.Run(ctx); err != nil {
		return
	}

	for _, id := range strings.Split(out.String(), "\n") {
		if id != "" {
			ids = append(ids, id)
		}
	}
	return
}

func kill(ctx context.Context, id string) (err error) {
	cmd := execkit.Command{
		Name:   "docker",
		Args:   []string{"kill", id},
		Stderr: os.Stderr,
	}
	cmd.Print(os.Stdout)
	return cmd.Run(ctx)
}
