package typical

import (
	"os"

	"github.com/typical-go/typical-go/pkg/exor"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
)

// ReactDemoModule is build module for react-demo
type ReactDemoModule struct {
	source string
}

// NewReactDemoModule return new constructor of ReactDemoModule
func NewReactDemoModule() *ReactDemoModule {
	return &ReactDemoModule{
		source: "react-demo",
	}
}

// Build the react-demo
func (m *ReactDemoModule) Build(c *typbuildtool.BuildContext) (dists []typbuildtool.BuildDistribution, err error) {
	c.Info("Build react-demo")
	err = exor.NewCommand("npm", "run", "build").
		WithDir(m.source).
		Execute(c.Cli.Context)
	return
}

// Clean the react-demo
func (m *ReactDemoModule) Clean(c *typbuildtool.BuildContext) (err error) {
	c.Info("Clean react-demo")
	if err := os.RemoveAll(m.source + "/build"); err != nil {
		c.Errorf("React-Demo: Clean: %w", err)
	}
	return
}
