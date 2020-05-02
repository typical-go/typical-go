package typical

import (
	"os"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
)

var (
	_ typbuildtool.Runner  = (*ReactDemoModule)(nil)
	_ typbuildtool.Cleaner = (*ReactDemoModule)(nil)
)

// ReactDemoModule is build module for react-demo
type ReactDemoModule struct {
	source string
}

// Run the react-demo
func (m *ReactDemoModule) Run(c *typbuildtool.CliContext) (err error) {
	c.Info("Build react-demo")
	cmd := &buildkit.Command{
		Name: "npm",
		Args: []string{"run", "build"},
		Dir:  m.source,
	}

	return cmd.Run(c.Context)
}

// Clean the react-demo
func (m *ReactDemoModule) Clean(c *typbuildtool.CliContext) (err error) {
	c.Info("Clean react-demo")
	if err := os.RemoveAll(m.source + "/build"); err != nil {
		c.Warnf("React-Demo: Clean: %s", err.Error())
	}
	return
}
