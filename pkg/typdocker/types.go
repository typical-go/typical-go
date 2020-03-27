package typdocker

// Composer responsible to compose docker
type Composer interface {
	DockerCompose(version Version) *Recipe
}

// Services descriptor in docker-compose.yml
type Services map[string]interface{}

// Networks descriptor in docker-compose.yml
type Networks map[string]interface{}

// Volumes descriptor in docker-compose.yml
type Volumes map[string]interface{}

// Network in docker-compose.yaml
type Network struct {
	Driver string `yaml:"driver,omitempty"`
}

// Service in docker-compose.yaml
type Service struct {
	Image       string            `yaml:"image,omitempty"`
	Command     string            `yaml:"command,omitempty"`
	Environment map[string]string `yaml:"environment,omitempty"`
	Volumes     []string          `yaml:"volumes,omitempty"`
	Ports       []string          `yaml:"ports,omitempty"`
	Networks    []string          `yaml:"networks,omitempty"`
	Restart     string            `yaml:"restart,omitempty"`
}
