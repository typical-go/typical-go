package typtmpl_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typtmpl"
)

func TestBuildToolMain(t *testing.T) {
	testTemplate(t,
		testcase{
			Template: &typtmpl.BuildToolMain{DescPkg: "some-package"},
			expected: `package main

import (
	"github.com/typical-go/typical-go/pkg/typcore"
	"some-package"
)

func main() {
	typcore.LaunchBuild(&typical.Descriptor)
}
`,
		},
	)
}
