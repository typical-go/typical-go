package typapp

import (
	"github.com/typical-go/typical-go/pkg/typdep"
)

// MainInvocation invoke when the application start
type MainInvocation struct {
	*typdep.Invocation
}

// NewMainInvocation return new instance of MainInvocation
func NewMainInvocation(fn interface{}) *MainInvocation {
	return &MainInvocation{
		Invocation: typdep.NewInvocation(fn),
	}
}

// EntryPoint of the module
func (m *MainInvocation) EntryPoint() *MainInvocation {
	return m
}
