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
		Action: m.wipeAction(c),
	}
}

func (m *DockerUtility) wipeAction(c *typbuildtool.Context) cli.ActionFunc {
	return func(cc *cli.Context) (err error) {
		var (
			ids []string
		)

		if ids, err = dockerIDs(cc.Context); err != nil {
			return fmt.Errorf("Docker-ID: %w", err)
		}

		c.Info("Wipe all docker container")
		for _, id := range ids {
			if err = kill(cc.Context, id); err != nil {
				c.Warnf("Fail to kill #%s: %w", id, err)
			}
		}
		return nil
	}
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
