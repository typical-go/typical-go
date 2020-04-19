package typdocker

import (
	"errors"
	"io/ioutil"

	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

const (
	dockerComposeFile = "docker-compose.yml"
)

func (m *DockerUtility) cmdCompose(c *typbuildtool.Context) *cli.Command {
	return &cli.Command{
		Name:   "compose",
		Usage:  "Generate docker-compose.yaml",
		Action: m.composeAction(c),
	}
}

func (m *DockerUtility) composeAction(c *typbuildtool.Context) cli.ActionFunc {
	var (
		out []byte
	)

	return func(cc *cli.Context) (err error) {
		if len(m.composers) < 1 {
			return errors.New("Nothing to compose")
		}

		if out, err = yaml.Marshal(m.recipe()); err != nil {
			return
		}

		c.Info("Generate docker-compose.yml")
		if err = ioutil.WriteFile(dockerComposeFile, out, 0644); err != nil {
			return
		}
		return
	}
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
