package typdocker

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"gopkg.in/yaml.v2"

	"github.com/urfave/cli/v2"
)

var (
	// DockerComposeYml is yml file
	DockerComposeYml = "docker-compose.yml"
	// Version of docker compose
	Version = "3"
)

type (
	// Command for docker
	Command struct {
		Composers []Composer
	}
	// Composer responsible to compose docker
	Composer interface {
		ComposeV3() (*Recipe, error)
	}
	// ComposeFn function
	ComposeFn    func() (*Recipe, error)
	composerImpl struct {
		fn ComposeFn
	}
)

//
// Command
//

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
		Action: c.ActionFn(m.Compose),
	}
}

// Compose to generate docker-compose.yml
func (m *Command) Compose(c *typgo.Context) (err error) {
	if len(m.Composers) < 1 {
		return errors.New("Nothing to compose")
	}

	root, err := compile(Version, m.Composers)
	if err != nil {
		return fmt.Errorf("compile: %w", err)
	}

	out, err := yaml.Marshal(root)
	if err != nil {
		return err
	}

	fmt.Println("Generate docker-compose.yml")
	return ioutil.WriteFile(DockerComposeYml, out, 0777)
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
		obj, err := composer.ComposeV3()
		if err != nil {
			return nil, err
		}
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
	return root, nil
}

func (m *Command) cmdWipe(c *typgo.BuildCli) *cli.Command {
	return &cli.Command{
		Name:   "wipe",
		Usage:  "Kill all running docker container",
		Action: c.ActionFn(m.dockerWipe),
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
			log.Printf("Fail to kill #%s: %s", id, err.Error())
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
		Action: c.ActionFn(m.dockerUp),
	}
}

func (m *Command) dockerUp(c *typgo.Context) (err error) {
	if c.Bool("wipe") {
		m.dockerWipe(c)
	}
	if _, err = os.Stat(DockerComposeYml); os.IsNotExist(err) {
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
		Action:  c.ActionFn(dockerDown),
	}
}

func dockerDown(c *typgo.Context) error {
	return c.Execute(&execkit.Command{
		Name: "docker-compose",
		Args: []string{"down"},
	})
}

func dockerIDs(ctx context.Context) (ids []string, err error) {
	var out strings.Builder
	cmd := &execkit.Command{
		Name:   "docker",
		Args:   []string{"ps", "-q"},
		Stderr: os.Stderr,
		Stdout: &out,
	}

	execkit.PrintCommand(cmd, os.Stdout)

	if err = cmd.Run(ctx); err != nil {
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
	cmd := &execkit.Command{
		Name:   "docker",
		Args:   []string{"kill", id},
		Stderr: os.Stderr,
	}
	execkit.PrintCommand(cmd, os.Stdout)
	return cmd.Run(ctx)
}

//
// composerImpl
//

var _ Composer = (*composerImpl)(nil)

// NewCompose return new instance of composer
func NewCompose(fn ComposeFn) Composer {
	return &composerImpl{fn: fn}
}

func (i *composerImpl) ComposeV3() (*Recipe, error) {
	return i.fn()
}
