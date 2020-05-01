package typdocker

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/urfave/cli/v2"
)

func (m *DockerUtility) cmdWipe(c *typbuildtool.Context) *cli.Command {
	return &cli.Command{
		Name:   "wipe",
		Usage:  "Kill all running docker container",
		Action: c.ActionFunc(m.dockerWipe),
	}
}

func (m *DockerUtility) dockerWipe(c *typbuildtool.CliContext) (err error) {
	var (
		ids []string
	)

	if ids, err = dockerIDs(c.Context); err != nil {
		return fmt.Errorf("Docker-ID: %w", err)
	}

	c.Info("Wipe all docker container")
	for _, id := range ids {
		if err = kill(c.Context, id); err != nil {
			c.Warnf("Fail to kill #%s: %s", id, err.Error())
		}
	}
	return nil

}

func dockerIDs(ctx context.Context) (ids []string, err error) {
	var (
		out strings.Builder
	)

	cmd := exec.CommandContext(ctx, "docker", "ps", "-q")
	cmd.Stderr = os.Stderr
	cmd.Stdout = &out

	if err = cmd.Run(); err != nil {
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
	cmd := exec.CommandContext(ctx, "docker", "kill", id)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
