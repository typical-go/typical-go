package typical

import (
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
)

// Descriptor of sample
var Descriptor = typgo.Descriptor{
	Name:    "server-echo-react",
	Version: "1.0.0",
	Layouts: []string{"internal"},

	Compile: &typgo.StdCompile{
		Before: typgo.NewCompiler(npmBuild),
	},
	Run:   &typgo.StdRun{},
	Clean: &typgo.StdClean{},
}

func npmBuild(c *typgo.Context) error {
	return c.Execute(&execkit.Command{
		Name: "npm",
		Args: []string{"run", "build"},
		Dir:  "react-demo",
	})
}
