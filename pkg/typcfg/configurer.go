package typcfg

var _ Config = (Configs)(nil)
var _ Config = &Configuration{}

type (
	// Config responsible to create config
	Config interface {
		Configurations() []*Configuration
	}

	// Configs is list of Configurer
	Configs []Config

	// Configuration is alias from typgo.Configuration with Configurer implementation
	Configuration struct {
		CtorName string
		Name     string
		Spec     interface{}
	}
)

// Configurations of configurer
func (c Configs) Configurations() (cfgs []*Configuration) {
	for _, configurer := range c {
		cfgs = append(cfgs, configurer.Configurations()...)
	}
	return
}
