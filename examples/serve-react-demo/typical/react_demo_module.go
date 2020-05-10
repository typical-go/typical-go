package typical

import (
	"os"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var (
	_ typgo.Runner  = (*ReactDemoModule)(nil)
	_ typgo.Cleaner = (*ReactDemoModule)(nil)
)

// ReactDemoModule is build module for react-demo
type ReactDemoModule struct {
	source string
}

// Run the react-demo
func (m *ReactDemoModule) Run(c *typgo.CliContext) (err error) {
	c.Info("Build react-demo")
	cmd := &execkit.Command{
		Name: "npm",
		Args: []string{"run", "build"},
		Dir:  m.source,
	}

	return cmd.Run(c.Cli.Context)
}

// Clean the react-demo
func (m *ReactDemoModule) Clean(c *typgo.CliContext) (err error) {
	c.Info("Clean react-demo")
	if err := os.RemoveAll(m.source + "/build"); err != nil {
		c.Warnf("React-Demo: Clean: %s", err.Error())
	}
	return
}
