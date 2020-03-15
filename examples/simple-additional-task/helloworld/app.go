package helloworld

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typcore"
)

// App of hello-world
type App struct {
}

// New return new instance of application
func New() *App {
	return &App{}
}

// RunApp to run the hello-world
func (*App) RunApp(d *typcore.Descriptor) error {
	fmt.Println("Hello World")
	return nil
}
