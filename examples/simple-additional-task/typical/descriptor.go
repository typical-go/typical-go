package typical

import (
	"github.com/typical-go/typical-go/examples/simple-additional-task/internal/helloworld"
	"github.com/typical-go/typical-go/pkg/typgo"
)

// Descriptor of sample
var Descriptor = typgo.Descriptor{
	Name:    "simple-additional-task",
	Version: "1.0.0",

	EntryPoint: helloworld.Main,

	Compile: &typgo.StdCompile{},

	Run: &typgo.StdRun{},

	Utility: typgo.NewUtility(taskPrintContext), // Add custom task
}
