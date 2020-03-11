package server

import (
	"net/http"

	"github.com/typical-go/typical-go/pkg/typcore"
)

// App of server
type App struct{}

// New instance of App
func New() *App {
	return &App{}
}

// Run app
func (*App) Run(d *typcore.Descriptor) error {
	fs := http.FileServer(http.Dir("react-demo/build"))
	http.Handle("/", fs)

	return http.ListenAndServe(":3000", nil)
}
