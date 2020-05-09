package typdocker

import (
	"errors"
	"io/ioutil"

	"github.com/typical-go/typical-go/pkg/typcore"

	"github.com/urfave/cli/v2"
)

var _ typcore.Utility = (*DockerUtility)(nil)

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
func (m *DockerUtility) Commands(c *typcore.Context) []*cli.Command {
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

func (m *DockerUtility) cmdCompose(c *typcore.Context) *cli.Command {
	return &cli.Command{
		Name:   "compose",
		Usage:  "Generate docker-compose.yaml",
		Action: c.ActionFunc(LogName, m.dockerCompose),
	}
}

func (m *DockerUtility) dockerCompose(c *typcore.CliContext) (err error) {
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
