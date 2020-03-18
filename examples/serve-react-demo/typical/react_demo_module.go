package typical

import (
	"github.com/typical-go/typical-go/pkg/exor"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
)

// ReactDemoModule is build module for react-demo
type ReactDemoModule struct {
}

// NewReactDemoModule return new constructor of ReactDemoModule
func NewReactDemoModule() *ReactDemoModule {
	return &ReactDemoModule{}
}

// Build the react-demo
func (m *ReactDemoModule) Build(c *typbuildtool.Context) (dists []typbuildtool.BuildDistribution, err error) {
	c.Info("Build react-demo")
	err = exor.NewCommand("npm", "run", "build").
		WithDir("react-demo").
		Execute(c.Cli.Context)
	return
}
