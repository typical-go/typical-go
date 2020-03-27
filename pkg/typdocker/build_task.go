package typdocker

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typbuildtool"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

// BuildTask for docker
type BuildTask struct {
	version   string
	composers []Composer
}

// Compose new docker module
func Compose(composers ...Composer) *BuildTask {
	return &BuildTask{
		version:   V3,
		composers: composers,
	}
}

// WithVersion to set the version
func (m *BuildTask) WithVersion(version string) *BuildTask {
	m.version = version
	return m
}

// Commands of docker
func (m *BuildTask) Commands(ctx *typbuildtool.Context) []*cli.Command {
	return []*cli.Command{
		{
			Name:  "docker",
			Usage: "Docker utility",
			Subcommands: []*cli.Command{
				{
					Name:  "compose",
					Usage: "Generate docker-compose.yaml",
					Action: func(c *cli.Context) (err error) {
						if len(m.composers) < 1 {
							return errors.New("Nothing to compose")
						}
						var out []byte
						log.Info("Generate docker-compose.yml")
						if out, err = yaml.Marshal(m.dockerCompose()); err != nil {
							return
						}
						if err = ioutil.WriteFile("docker-compose.yml", out, 0644); err != nil {
							return
						}
						return
					},
				},
				{
					Name:  "up",
					Usage: "Spin up docker containers according docker-compose",
					Action: func(ctx *cli.Context) (err error) {
						cmd := exec.Command("docker-compose", "up", "--remove-orphans", "-d")
						cmd.Stderr = os.Stderr
						cmd.Stdout = os.Stdout
						return cmd.Run()
					},
				},
				{
					Name:  "down",
					Usage: "Take down all docker containers according docker-compose",
					Action: func(ctx *cli.Context) (err error) {
						cmd := exec.Command("docker-compose", "down")
						cmd.Stderr = os.Stderr
						cmd.Stdout = os.Stdout
						return cmd.Run()
					},
				},
				{
					Name:  "wipe",
					Usage: "Kill all running docker container",
					Action: func(ctx *cli.Context) (err error) {
						var builder strings.Builder
						cmd := exec.Command("docker", "ps", "-q")
						cmd.Stderr = os.Stderr
						cmd.Stdout = &builder
						if err = cmd.Run(); err != nil {
							return
						}
						dockerIDs := strings.Split(builder.String(), "\n")
						for _, id := range dockerIDs {
							if id != "" {
								cmd := exec.Command("docker", "kill", id)
								cmd.Stderr = os.Stderr
								if err = cmd.Run(); err != nil {
									return
								}
							}
						}
						return
					},
				},
			},
		},
	}
}

func (m *BuildTask) dockerCompose() (root *Recipe) {
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
