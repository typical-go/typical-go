package typdocker

import (
	"errors"
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
func (m *DockerUtility) Commands(c *typgo.BuildTool) []*cli.Command {
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

func (m *DockerUtility) cmdCompose(c *typgo.BuildTool) *cli.Command {
	return &cli.Command{
		Name:   "compose",
		Usage:  "Generate docker-compose.yaml",
		Action: c.ActionFunc(LogName, m.dockerCompose),
	}
}

func (m *DockerUtility) dockerCompose(c *typgo.CliContext) (err error) {
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

func (m *DockerUtility) cmdUp(c *typgo.BuildTool) *cli.Command {
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

func (m *DockerUtility) dockerUp(c *typgo.CliContext) (err error) {
	if c.Cli.Bool("wipe") {
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

	return cmd.Run(c.Cli.Context)
}

func (m *DockerUtility) cmdDown(c *typgo.BuildTool) *cli.Command {
	return &cli.Command{
		Name:    "down",
		Aliases: []string{"stop"},
		Usage:   "Take down all docker containers according docker-compose",
		Action:  c.ActionFunc(LogName, dockerDown),
	}
}

func dockerDown(c *typgo.CliContext) error {
	cmd := &execkit.Command{
		Name:   "docker-compose",
		Args:   []string{"down"},
		Stderr: os.Stderr,
		Stdout: os.Stdout,
	}
	cmd.Print(os.Stdout)
	return cmd.Run(c.Cli.Context)
}
