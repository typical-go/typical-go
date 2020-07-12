package typtmpl_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typtmpl"
)

func TestBuildToolMain(t *testing.T) {
	typtmpl.TestTemplate(t, []typtmpl.TestCase{
		{
			Template: &typtmpl.BuildToolMain{DescPkg: "some-package"},
			Expected: `package main

import (
	"github.com/typical-go/typical-go/pkg/typcore"
	"some-package"
)

func main() {
	typcore.LaunchBuild(&typical.Descriptor)
}
`,
		},
	})
}
