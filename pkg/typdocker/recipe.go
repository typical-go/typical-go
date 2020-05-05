package typdocker

import "gopkg.in/yaml.v2"

var (
	_ Composer = (*Recipe)(nil)
)

// Composer responsible to compose docker
type Composer interface {
	DockerCompose() *Recipe
}

// Recipe represent docker-compose.yml
type Recipe struct {
	Version  string
	Services Services
	Networks Networks
	Volumes  Volumes
}

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
