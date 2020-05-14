package typdocker

import "gopkg.in/yaml.v2"

var (
	_ Composer = (*Recipe)(nil)
)

type (
	// Composer responsible to compose docker
	Composer interface {
		DockerCompose() *Recipe
	}

	// Recipe represent docker-compose.yml
	Recipe struct {
		Version  string
		Services Services
		Networks Networks
		Volumes  Volumes
	}

	// Services descriptor in docker-compose.yml
	Services map[string]interface{}

	// Networks descriptor in docker-compose.yml
	Networks map[string]interface{}

	// Volumes descriptor in docker-compose.yml
	Volumes map[string]interface{}

	// Network in docker-compose.yaml
	Network struct {
		Driver string `yaml:"driver,omitempty"`
	}

	// Service in docker-compose.yaml
	Service struct {
		Image       string            `yaml:"image,omitempty"`
		Command     string            `yaml:"command,omitempty"`
		Environment map[string]string `yaml:"environment,omitempty"`
		Volumes     []string          `yaml:"volumes,omitempty"`
		Ports       []string          `yaml:"ports,omitempty"`
		Networks    []string          `yaml:"networks,omitempty"`
		Restart     string            `yaml:"restart,omitempty"`
	}
)

// DockerCompose to get the recipe
func (c *Recipe) DockerCompose() *Recipe {
	return c
}

// ComposeRecipe to compose recipe
func ComposeRecipe(version string, composers []Composer) ([]byte, error) {
	root := &Recipe{
		Version:  version,
		Services: make(Services),
		Networks: make(Networks),
		Volumes:  make(Volumes),
	}
	for _, composer := range composers {
		obj := composer.DockerCompose()
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

	return yaml.Marshal(root)
}
