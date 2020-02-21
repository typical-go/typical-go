package tmpl

// MainSrcData is data for main src template
type MainSrcData struct {
	DescriptorPackage string
}

// MainSrcApp is template for main source for app
const MainSrcApp = `package main

import (
	"github.com/typical-go/typical-go/pkg/typcore"
	"{{.DescriptorPackage}}"
)

func main() {
	typcore.RunApp(&typical.Descriptor)
}
`

// MainSrcBuildTool is template for main source for build tool
const MainSrcBuildTool = `package main

import (
	"github.com/typical-go/typical-go/pkg/typcore"
	"{{.DescriptorPackage}}"
)

func main() {
	typcore.RunBuildTool(&typical.Descriptor)
}
`
