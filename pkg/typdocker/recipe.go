package typdocker

// Recipe represent docker-compose.yml
type Recipe struct {
	Version  Version
	Services Services
	Networks Networks
	Volumes  Volumes
}

// Append another compose object
func (c *Recipe) Append(comp *Recipe) {
	for k, service := range comp.Services {
		c.Services[k] = service
	}
	for k, network := range comp.Networks {
		c.Networks[k] = network
	}
	for k, volume := range comp.Volumes {
		c.Volumes[k] = volume
	}
}

// DockerCompose to get the recipe
func (c *Recipe) DockerCompose(version Version) *Recipe {
	return c
}
