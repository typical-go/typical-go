package typdocker

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"gopkg.in/yaml.v2"

	"github.com/urfave/cli/v2"
)

const (
	dockerComposeOut = "docker-compose.yml"

	logName = "docker"

	// V3 is version 3
	V3 = "3"
)

type (
	// Command for docker
	Command struct {
		Version   string
		Composers []Composer
	}
)

var _ typgo.Cmd = (*Command)(nil)

// Command of docker
func (m *Command) Command(c *typgo.BuildCli) *cli.Command {
	return &cli.Command{
		Name:  "docker",
		Usage: "Docker utility",
		Subcommands: []*cli.Command{
			m.cmdCompose(c),
			m.cmdUp(c),
			m.cmdDown(c),
			m.cmdWipe(c),
		},
	}
}

func (m *Command) cmdCompose(c *typgo.BuildCli) *cli.Command {
	return &cli.Command{
		Name:   "compose",
		Usage:  "Generate docker-compose.yaml",
		Action: c.ActionFn(logName, m.Compose),
	}
}

// Compose to generate docker-compose.yml
func (m *Command) Compose(c *typgo.Context) (err error) {
	if len(m.Composers) < 1 {
		return errors.New("Nothing to compose")
	}

	root, err := compile(m.Version, m.Composers)
	if err != nil {
		return fmt.Errorf("compile: %w", err)
	}

	out, err := yaml.Marshal(root)
	if err != nil {
		return err
	}

	c.Info("Generate docker-compose.yml")
	return ioutil.WriteFile(dockerComposeOut, out, 0777)
}

// Compile recipes to yaml
func compile(version string, composers []Composer) (*Recipe, error) {
	root := &Recipe{
		Version:  version,
		Services: make(Services),
		Networks: make(Networks),
		Volumes:  make(Volumes),
	}
	for _, composer := range composers {
		obj, err := composer.Compose()
		if err != nil {
			return nil, err
		}
		if obj != nil && obj.Version == version {
			for k, service := range obj.Services {
				root.Services[k] = service
			}
			for k, network := range obj.Networks {
				root.Networks[k] = network
			}
			for k, volume := range obj.Volumes {
				root.Volumes[k] = volume
			}
		}
	}
	return root, nil
}

func (m *Command) cmdWipe(c *typgo.BuildCli) *cli.Command {
	return &cli.Command{
		Name:   "wipe",
		Usage:  "Kill all running docker container",
		Action: c.ActionFn(logName, m.dockerWipe),
	}
}

func (m *Command) dockerWipe(c *typgo.Context) (err error) {
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

func (m *Command) cmdUp(c *typgo.BuildCli) *cli.Command {
	return &cli.Command{
		Name:    "up",
		Aliases: []string{"start"},
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "wipe"},
		},
		Usage:  "Spin up docker containers according docker-compose",
		Action: c.ActionFn(logName, m.dockerUp),
	}
}

func (m *Command) dockerUp(c *typgo.Context) (err error) {
	if c.Bool("wipe") {
		m.dockerWipe(c)
	}

	if _, err = os.Stat(dockerComposeOut); os.IsNotExist(err) {
		if err = m.Compose(c); err != nil {
			return
		}
	}

	return c.Execute(&execkit.Command{
		Name: "docker-compose",
		Args: []string{"up", "--remove-orphans", "-d"},
	})
}

func (m *Command) cmdDown(c *typgo.BuildCli) *cli.Command {
	return &cli.Command{
		Name:    "down",
		Aliases: []string{"stop"},
		Usage:   "Take down all docker containers according docker-compose",
		Action:  c.ActionFn(logName, dockerDown),
	}
}

func dockerDown(c *typgo.Context) error {
	return c.Execute(&execkit.Command{
		Name: "docker-compose",
		Args: []string{"down"},
	})
}
