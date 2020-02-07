package tmpl

// MainSrcData is data for main src template
type MainSrcData struct {
	ImportTypical string
}

// MainSrcApp is template for main source for app
const MainSrcApp = `package main

import (
	"github.com/typical-go/typical-go/pkg/typapp"
	"{{.ImportTypical}}"
)

func main() {
	typapp.Run(&typical.Descriptor)
}
`

// MainSrcBuildTool is template for main source for build tool
const MainSrcBuildTool = `package main

import (
	"github.com/typical-go/typical-go/pkg/typbuild"
	"{{.ImportTypical}}"
)

func main() {
	typbuild.Run(&typical.Descriptor)
}
`
