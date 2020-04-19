package typdocker

import (
	"github.com/typical-go/typical-go/pkg/typbuildtool"

	"github.com/urfave/cli/v2"
)

var (
	_ typbuildtool.Utility = (*DockerUtility)(nil)
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
func (m *DockerUtility) Commands(c *typbuildtool.Context) []*cli.Command {
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
