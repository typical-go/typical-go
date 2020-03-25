package server

import (
	"net/http"

	"github.com/typical-go/typical-go/pkg/typcore"
)

// Main function to run server
func Main(d *typcore.Descriptor) error {
	fs := http.FileServer(http.Dir("react-demo/build"))
	http.Handle("/", fs)

	return http.ListenAndServe(":3000", nil)
}
