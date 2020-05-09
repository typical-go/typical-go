package typcfg

type (
	// Configurer responsible to create config
	Configurer interface {
		Configurations() []*Configuration
	}

	// Configurers is list of Configurer
	Configurers []Configurer
)

// Configurations of configurer
func (c Configurers) Configurations() (cfgs []*Configuration) {
	for _, configurer := range c {
		cfgs = append(cfgs, configurer.Configurations()...)
	}
	return
}
