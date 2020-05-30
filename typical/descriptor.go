package typical

import (
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/wrapper"
)

// Descriptor of typical-go
var Descriptor = typgo.Descriptor{
	Name:    "typical-go",
	Version: "0.9.55",

	EntryPoint: wrapper.Main,
	Layouts:    []string{"wrapper", "pkg"},

	Test:    &typgo.StdTest{},
	Compile: &typgo.StdCompile{},
	Run:     &typgo.StdRun{},
	Release: &typgo.Github{Owner: "typical-go", RepoName: "typical-go"},

	Utility: typgo.NewUtility(taskExamples), // Test all the examples
}
