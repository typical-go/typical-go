package typical

import (
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
)

// ReactDemoModule is build module for react-demo
type ReactDemoModule struct {
	source string
}

var _ typgo.Compiler = (*ReactDemoModule)(nil)

// Compile react demo
func (m *ReactDemoModule) Compile(c *typgo.Context) (err error) {
	c.Info("Build react-demo")
	cmd := &execkit.Command{
		Name: "npm",
		Args: []string{"run", "build"},
		Dir:  m.source,
	}
	return cmd.Run(c.Ctx())
}
