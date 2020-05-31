package typgo

type (
	// Configurer responsible to create config
	Configurer interface {
		Configurations() []*Configuration
	}

	// Configurers is list of Configurer
	Configurers []Configurer

	// Configuration is alias from typgo.Configuration with Configurer implementation
	Configuration struct {
		Ctor string
		Name string
		Spec interface{}
	}
)

var _ Configurer = (Configurers)(nil)
var _ Configurer = &Configuration{}

// Configurations of configurer
func (c Configurers) Configurations() (cfgs []*Configuration) {
	for _, configurer := range c {
		cfgs = append(cfgs, configurer.Configurations()...)
	}
	return
}
