package typapp

import (
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/urfave/cli/v2"
)

var (
	_ Provider          = (*Module)(nil)
	_ Destroyer         = (*Module)(nil)
	_ Preparer          = (*Module)(nil)
	_ Commander         = (*Module)(nil)
	_ typcfg.Configurer = (*Module)(nil)
)

// Module for application
type Module struct {
	providers   []Provider
	destroyers  []Destroyer
	preparers   []Preparer
	commanders  []Commander
	configurers []typcfg.Configurer
}

// NewModule return new instance of Module
func NewModule() *Module {
	return &Module{}
}

// Provide the module
func (m *Module) Provide(providers ...Provider) *Module {
	m.providers = providers
	return m
}

// Destroy the module
func (m *Module) Destroy(destroyers ...Destroyer) *Module {
	m.destroyers = destroyers
	return m
}

// Prepare the module
func (m *Module) Prepare(prepares ...Preparer) *Module {
	m.preparers = prepares
	return m
}

// Command the module
func (m *Module) Command(commanders ...Commander) *Module {
	m.commanders = commanders
	return m
}

// Configure tbe module
func (m *Module) Configure(configurers ...typcfg.Configurer) *Module {
	m.configurers = configurers
	return m
}

// Constructors of the module
func (m *Module) Constructors() (constructions []*Constructor) {
	for _, provider := range m.providers {
		constructions = append(constructions, provider.Constructors()...)
	}
	return
}

// Destructions of the module
func (m *Module) Destructions() (destructions []*Destruction) {
	for _, destroyer := range m.destroyers {
		destructions = append(destructions, destroyer.Destructions()...)
	}
	return
}

// Preparations of the module
func (m *Module) Preparations() (preparations []*Preparation) {
	for _, prepare := range m.preparers {
		preparations = append(preparations, prepare.Preparations()...)
	}
	return
}

// Commands of module
func (m *Module) Commands(c *Context) (cmds []*cli.Command) {
	for _, commander := range m.commanders {
		cmds = append(cmds, commander.Commands(c)...)
	}
	return
}

// Configurations of module
func (m *Module) Configurations() (cfgs []*typcfg.Configuration) {
	for _, configurer := range m.configurers {
		cfgs = append(cfgs, configurer.Configurations()...)
	}
	return
}
