package tmpl

type MainData struct {
	ImportTypical string
}

var MainAppSrc = `package main

import (
	"github.com/typical-go/typical-go/pkg/typapp"
	"{{.ImportTypical}}"
)

func main() {
	typapp.Run(&typical.Descriptor)
}
`

var MainBuildToolSrc = `package main

import (
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"{{.ImportTypical}}"
)

func main() {
	typbuildtool.Run(&typical.Descriptor)
}
`
