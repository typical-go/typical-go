package typdocker

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"

	"github.com/urfave/cli/v2"
)

var _ typgo.Utility = (*DockerUtility)(nil)

const (
	// DockerComposeFile contain full path of docker-compose.yml file
	DockerComposeFile = "docker-compose.yml"

	// LogName of docker utility
	LogName = "docker"

	// V3 is version 3
	V3 = "3"
)

// DockerUtility for docker
type DockerUtility struct {
	version   string
	composers []Composer
}

// Compose new docker module
func Compose(composers ...Composer) *DockerUtility {
	return &DockerUtility{
		version:   V3,
		composers: composers,
	}
}

// WithVersion to set the version
func (m *DockerUtility) WithVersion(version string) *DockerUtility {
	m.version = version
	return m
}

// Commands of docker
func (m *DockerUtility) Commands(c *typgo.BuildCli) []*cli.Command {
	return []*cli.Command{
		{
			Name:  "docker",
			Usage: "Docker utility",
			Subcommands: []*cli.Command{
				m.cmdCompose(c),
				m.cmdUp(c),
				m.cmdDown(c),
				m.cmdWipe(c),
			},
		},
	}
}

func (m *DockerUtility) cmdCompose(c *typgo.BuildCli) *cli.Command {
	return &cli.Command{
		Name:   "compose",
		Usage:  "Generate docker-compose.yaml",
		Action: c.ActionFn(LogName, m.dockerCompose),
	}
}

func (m *DockerUtility) cmdWipe(c *typgo.BuildCli) *cli.Command {
	return &cli.Command{
		Name:   "wipe",
		Usage:  "Kill all running docker container",
		Action: c.ActionFn(LogName, m.dockerWipe),
	}
}

func (m *DockerUtility) dockerWipe(c *typgo.Context) (err error) {
	var ids []string
	ctx := c.Ctx()
	if ids, err = dockerIDs(ctx); err != nil {
		return fmt.Errorf("Docker-ID: %w", err)
	}
	for _, id := range ids {
		if err = kill(ctx, id); err != nil {
			c.Warnf("Fail to kill #%s: %s", id, err.Error())
		}
	}
	return nil
}

func (m *DockerUtility) dockerCompose(c *typgo.Context) (err error) {
	var (
		out []byte
	)

	if len(m.composers) < 1 {
		return errors.New("Nothing to compose")
	}

	if out, err = ComposeRecipe(m.version, m.composers); err != nil {
		return
	}

	c.Info("Generate docker-compose.yml")
	return ioutil.WriteFile(DockerComposeFile, out, 0644)
}

func (m *DockerUtility) cmdUp(c *typgo.BuildCli) *cli.Command {
	return &cli.Command{
		Name:    "up",
		Aliases: []string{"start"},
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "wipe"},
		},
		Usage:  "Spin up docker containers according docker-compose",
		Action: c.ActionFn(LogName, m.dockerUp),
	}
}

func (m *DockerUtility) dockerUp(c *typgo.Context) (err error) {
	if c.Bool("wipe") {
		m.dockerWipe(c)
	}

	if _, err = os.Stat(DockerComposeFile); os.IsNotExist(err) {
		if err = m.dockerCompose(c); err != nil {
			return
		}
	}

	cmd := &execkit.Command{
		Name:   "docker-compose",
		Args:   []string{"up", "--remove-orphans", "-d"},
		Stderr: os.Stderr,
		Stdout: os.Stdout,
	}

	cmd.Print(os.Stdout)

	return cmd.Run(c.Ctx())
}

func (m *DockerUtility) cmdDown(c *typgo.BuildCli) *cli.Command {
	return &cli.Command{
		Name:    "down",
		Aliases: []string{"stop"},
		Usage:   "Take down all docker containers according docker-compose",
		Action:  c.ActionFn(LogName, dockerDown),
	}
}

func dockerDown(c *typgo.Context) error {
	cmd := &execkit.Command{
		Name:   "docker-compose",
		Args:   []string{"down"},
		Stderr: os.Stderr,
		Stdout: os.Stdout,
	}
	cmd.Print(os.Stdout)
	return cmd.Run(c.Ctx())
}
