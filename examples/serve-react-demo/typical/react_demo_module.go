package typical

import (
	"os"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var (
	_ typgo.Build = (*ReactDemoModule)(nil)
)

// ReactDemoModule is build module for react-demo
type ReactDemoModule struct {
	source string
}

// Execute build
func (m *ReactDemoModule) Execute(c *typgo.Context, phase typgo.Phase) (ok bool, err error) {
	switch phase {
	case typgo.RunPhase:
		return true, m.executeRun(c)
	case typgo.CleanPhase:
		return true, m.executeClean(c)
	}
	return
}

func (m *ReactDemoModule) executeRun(c *typgo.Context) (err error) {
	c.Info("Build react-demo")
	cmd := &execkit.Command{
		Name: "npm",
		Args: []string{"run", "build"},
		Dir:  m.source,
	}

	return cmd.Run(c.Ctx())
}

func (m *ReactDemoModule) executeClean(c *typgo.Context) (err error) {
	c.Info("Clean react-demo")
	if err := os.RemoveAll(m.source + "/build"); err != nil {
		c.Warnf("React-Demo: Clean: %s", err.Error())
	}
	return
}
