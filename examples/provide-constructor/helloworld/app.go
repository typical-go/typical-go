package helloworld

import (
	"github.com/typical-go/typical-go/pkg/typdep"
)

// App of hello world
type App struct{}

// New return new instance of application
func New() *App {
	return &App{}
}

// EntryPoint of application
func (a *App) EntryPoint() *typdep.Invocation {
	return typdep.NewInvocation(start)
}

func start(greeter *Greeter) {
	greeter.Greet()
}
