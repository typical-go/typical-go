package typfactory_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typfactory"
)

func TestBuildToolMain(t *testing.T) {
	testWriter(t,
		testcase{
			Writer: &typfactory.BuildToolMain{DescPkg: "some-package"},
			expected: `package main

import (
	"github.com/typical-go/typical-go/pkg/typcore"
	"some-package"
)

func main() {
	typcore.LaunchBuildTool(&typical.Descriptor)
}
`,
		},
	)
}
