package typical

import (
	"github.com/typical-go/typical-go/examples/serve-react-demo/internal/server"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
)

// Descriptor of sample
var Descriptor = typgo.Descriptor{
	Name:    "server-echo-react",
	Version: "1.0.0",

	EntryPoint: server.Main,
	Layouts:    []string{"internal"},

	Compile: typgo.Compiles{
		typgo.NewCompile(func(c *typgo.Context) (err error) {
			return c.Execute(&execkit.Command{
				Name: "npm",
				Args: []string{"run", "build"},
				Dir:  "react-demo",
			})
		}),
		&typgo.StdCompile{},
	},
	Run:   &typgo.StdRun{},
	Clean: &typgo.StdClean{},
}
