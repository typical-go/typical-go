package helloworld

import (
	"github.com/typical-go/typical-go/pkg/typdep"
)

// App of hello world
type App struct {
	ConfigName string
}

// New return new instance of application
func New() *App {
	return &App{
		ConfigName: "APP",
	}
}

// WithConfigPrefix return Module with new config prefix
func (a *App) WithConfigPrefix(name string) *App {
	a.ConfigName = name
	return a
}

// EntryPoint of application
func (a *App) EntryPoint() *typdep.Invocation {
	return typdep.NewInvocation(start)
}

func start(greeter *Greeter) {
	greeter.Greet()
}
