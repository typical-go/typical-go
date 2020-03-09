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

// Run app
func (*App) Run(d *typcore.Descriptor) error {
	fmt.Println("Hello World")
	return nil
}
