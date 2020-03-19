package helloworld

import (
	"fmt"
	"io"
	"os"

	"github.com/typical-go/typical-go/pkg/typdep"
)

// App of hello world
type App struct {
	io.Writer
}

// New return new instance of application
func New() *App {
	return &App{
		Writer: os.Stdout,
	}
}

// WithWriter return App with new writer
func (a *App) WithWriter(w io.Writer) *App {
	a.Writer = w
	return a
}

// EntryPoint of application
func (a *App) EntryPoint() *typdep.Invocation {
	return typdep.NewInvocation(a.Print)
}

// Print text from greeter
func (a *App) Print(greeter Greeter) {
	fmt.Fprintln(a, greeter.Greet())
}
