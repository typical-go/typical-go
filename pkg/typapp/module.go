package typapp

import (
	"github.com/urfave/cli/v2"
)

// Module for application
type Module struct {
	providers  []Provider
	destroyers []Destroyer
	preparers  []Preparer
	commanders []Commander
}

// NewModule return new instance of Module
func NewModule() *Module {
	return &Module{}
}

// WithProviders return Module with new providers
func (m *Module) WithProviders(providers ...Provider) *Module {
	m.providers = providers
	return m
}

// WithDestoyers return Module with new destroyers
func (m *Module) WithDestoyers(destroyers ...Destroyer) *Module {
	m.destroyers = destroyers
	return m
}

// WithPrepares return Module with new preparers
func (m *Module) WithPrepares(prepares ...Preparer) *Module {
	m.preparers = prepares
	return m
}

// WithCommanders return Module with new commanders
func (m *Module) WithCommanders(commanders ...Commander) *Module {
	m.commanders = commanders
	return m
}

// Provide dependency
func (m *Module) Provide() (constructions []*Constructor) {
	for _, provider := range m.providers {
		constructions = append(constructions, provider.Provide()...)
	}
	return
}

// Destroy dependency
func (m *Module) Destroy() (destructions []*Destruction) {
	for _, destroyer := range m.destroyers {
		destructions = append(destructions, destroyer.Destroy()...)
	}
	return
}

// Prepare dependency
func (m *Module) Prepare() (preparations []*Preparation) {
	for _, prepare := range m.preparers {
		preparations = append(preparations, prepare.Prepare()...)
	}
	return
}

// Commands of application
func (m *Module) Commands(c *Context) (cmds []*cli.Command) {
	for _, commander := range m.commanders {
		cmds = append(cmds, commander.Commands(c)...)
	}
	return
}
