package bashkit

import (
	"go/build"
	"os"
)

// GOPATH return GOPATH variable
func GOPATH() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}

	return gopath
}
