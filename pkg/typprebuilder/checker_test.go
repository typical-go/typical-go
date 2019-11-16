package typprebuilder

import (
	"io"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typctx"
)

func TestChecker_CheckBuildTool(t *testing.T) {
	testcases := []struct {
		checker        checker
		checkBuildTool bool
	}{
		{checker{}, false},
		{checker{mockTarget: true}, true},
		{checker{configuration: true}, true},
		{checker{testTarget: true}, true},
		{checker{buildToolBinary: true}, true},
		{checker{contextChecksum: true}, true},
		{checker{buildCommands: true}, true},
	}
	for _, tt := range testcases {
		require.Equal(t, tt.checkBuildTool, tt.checker.checkBuildTool())
	}
}

func TestChecker_CheckReadme(t *testing.T) {
	testcases := []struct {
		checker     checker
		checkReadme bool
	}{
		{
			checker{
				Context: &typctx.Context{ReadmeGenerator: dummyReadmeGenerator{}},
			},
			false,
		},
		{
			checker{
				Context:       &typctx.Context{ReadmeGenerator: dummyReadmeGenerator{}},
				configuration: true,
			},
			true,
		},
		{
			checker{
				Context:       &typctx.Context{ReadmeGenerator: dummyReadmeGenerator{}},
				buildCommands: true,
			},
			true,
		},
		{
			checker{
				Context:    &typctx.Context{ReadmeGenerator: dummyReadmeGenerator{}},
				readmeFile: true,
			},
			true,
		},
		{
			checker{
				Context:    &typctx.Context{},
				readmeFile: false,
			},
			false,
		},
	}
	for i, tt := range testcases {
		require.Equal(t, tt.checkReadme, tt.checker.checkReadme(), i)
	}
}

type dummyReadmeGenerator struct{}

func (dummyReadmeGenerator) Generate(*typctx.Context, io.Writer) error {
	return nil
}
