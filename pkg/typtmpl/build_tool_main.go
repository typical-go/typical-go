package typtmpl

import (
	"io"
)

var _ Template = (*BuildToolMain)(nil)

const buildtoolMain = `package main

import (
	"github.com/typical-go/typical-go/pkg/typgo"
	"{{.DescPkg}}"
)

func main() {
	typgo.LaunchBuildTool(&typical.Descriptor)
}
`

// BuildToolMain is writer to generate main.go for app
type BuildToolMain struct {
	DescPkg string
}

// Execute build tool main template
func (t *BuildToolMain) Execute(w io.Writer) (err error) {
	return Execute("buildToolMain", buildtoolMain, t, w)
}
