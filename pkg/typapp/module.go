package typapp

import (
	"github.com/urfave/cli/v2"
)

// Module for application
type Module struct {
	Provider  Provider
	Destroyer Destroyer
	Preparer  Preparer
	Commander Commander
}

// Provide dependency
func (m *Module) Provide() []*Constructor {
	if m.Provider != nil {
		return m.Provider.Provide()
	}
	return []*Constructor{}
}

// Destroy dependency
func (m *Module) Destroy() []*Destruction {
	if m.Destroyer != nil {
		return m.Destroyer.Destroy()
	}
	return []*Destruction{}
}

// Prepare dependency
func (m *Module) Prepare() []*Preparation {
	if m.Preparer != nil {
		return m.Preparer.Prepare()
	}
	return []*Preparation{}
}

// Commands of application
func (m *Module) Commands(c *Context) []*cli.Command {
	if m.Commander != nil {
		return m.Commander.Commands(c)
	}
	return []*cli.Command{}
}
