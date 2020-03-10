package tmpl

// MainSrcData is data for main src template
type MainSrcData struct {
	DescriptorPackage string
}

// MainSrcBuildTool is template for main source for build tool
const MainSrcBuildTool = `package main

import (
	"github.com/typical-go/typical-go/pkg/typcore"
	"{{.DescriptorPackage}}"
)

func main() {
	typcore.LaunchBuildTool(&typical.Descriptor)
}
`
