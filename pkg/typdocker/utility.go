package typdocker

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/typical-go/typical-go/pkg/typbuildtool"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

const (
	dockerComposeFile = "docker-compose.yml"
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
				{
					Name:  "compose",
					Usage: "Generate docker-compose.yaml",
					Action: func(cc *cli.Context) (err error) {
						c.Info("Generate docker-compose.yml")
						return m.compose()
					},
				},
				{
					Name:    "up",
					Aliases: []string{"start"},
					Flags: []cli.Flag{
						&cli.BoolFlag{Name: "wipe"},
					},
					Usage: "Spin up docker containers according docker-compose",
					Action: func(cc *cli.Context) (err error) {
						bc := c.BuildContext(cc)
						if cc.Bool("wipe") {
							c.Info("Wipe all docker container")
							m.wipe(bc)
						}
						c.Info("Docker up")
						return m.up(bc)
					},
				},
				{
					Name:    "down",
					Aliases: []string{"stop"},
					Usage:   "Take down all docker containers according docker-compose",
					Action: func(cc *cli.Context) (err error) {
						c.Info("Docker down")
						return m.down(c.BuildContext(cc))
					},
				},
				{
					Name:  "wipe",
					Usage: "Kill all running docker container",
					Action: func(cc *cli.Context) (err error) {
						c.Info("Wipe all docker container")
						return m.wipe(c.BuildContext(cc))
					},
				},
			},
		},
	}
}

func (m *DockerUtility) compose() (err error) {
	if len(m.composers) < 1 {
		return errors.New("Nothing to compose")
	}
	var out []byte
	if out, err = yaml.Marshal(m.recipe()); err != nil {
		return
	}
	if err = ioutil.WriteFile(dockerComposeFile, out, 0644); err != nil {
		return
	}
	return
}

func (m *DockerUtility) up(c *typbuildtool.BuildContext) (err error) {
	cmd := exec.CommandContext(c.Cli.Context, "docker-compose", "up", "--remove-orphans", "-d")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func (m *DockerUtility) down(c *typbuildtool.BuildContext) (err error) {
	cmd := exec.CommandContext(c.Cli.Context, "docker-compose", "down")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func (m *DockerUtility) wipe(c *typbuildtool.BuildContext) (err error) {
	var (
		builder strings.Builder
		ctx     = c.Cli.Context
	)
	cmd := exec.CommandContext(ctx, "docker", "ps", "-q")
	cmd.Stderr = os.Stderr
	cmd.Stdout = &builder
	if err = cmd.Run(); err != nil {
		return
	}
	dockerIDs := strings.Split(builder.String(), "\n")
	for _, id := range dockerIDs {
		if id != "" {
			cmd := exec.CommandContext(ctx, "docker", "kill", id)
			cmd.Stderr = os.Stderr
			if err = cmd.Run(); err != nil {
				return
			}
		}
	}
	return
}

func (m *DockerUtility) recipe() (root *Recipe) {
	root = &Recipe{
		Version:  m.version,
		Services: make(Services),
		Networks: make(Networks),
		Volumes:  make(Volumes),
	}
	for _, composer := range m.composers {
		if obj := composer.DockerCompose(m.version); obj != nil {
			root.Append(obj)
		}
	}
	return
}
