package typcore_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/typicalgo"
)

func TestTypicalContext(t *testing.T) {
	os.MkdirAll("typicalgo/some_pkg", os.ModePerm)
	os.MkdirAll("pkg/some_lib", os.ModePerm)
	os.Create("typicalgo/some_pkg/some_file.go")
	os.Create("typicalgo/some_pkg/not_go.xxx")
	os.Create("pkg/some_lib/lib.go")
	ioutil.WriteFile("go.mod", []byte("module github.com/typical-go/typical-go\ngo 1.13"), 0644)
	defer func() {
		os.RemoveAll("typicalgo")
		os.RemoveAll("pkg")
		os.Remove("go.mod")
	}()

	ctx, err := typcore.CreateContext(&typcore.Descriptor{
		Name:      "some-name",
		App:       typicalgo.New(),
		BuildTool: typbuildtool.New(),
	})

	// NOTE: ProjectPackage need to set manually because its value get from ldflags
	ctx.ProjectPackage = "some-package"

	require.NoError(t, err)
	require.NoError(t, common.Validate(ctx))
	require.Equal(t, "0.0.1", ctx.Version)
	require.Equal(t, []string{"typicalgo", "pkg"}, ctx.ProjectSources)
	require.Equal(t, []string{"typicalgo", "typicalgo/some_pkg", "pkg", "pkg/some_lib"}, ctx.Dirs)
	require.Equal(t, []string{"typicalgo/some_pkg/some_file.go", "pkg/some_lib/lib.go"}, ctx.Files)
}

func TestTypicalContext_Validate(t *testing.T) {
	testcases := []struct {
		*typcore.Context
		expectedError string
	}{
		{
			Context:       &typcore.Context{},
			expectedError: "TypicalContext: Descriptor can't be empty",
		},
		{
			Context: &typcore.Context{
				Descriptor: validDescriptor,
			},
			expectedError: "TypicalContext: ProjectPackage can't be empty",
		},
		{
			Context: &typcore.Context{
				Descriptor:     validDescriptor,
				ProjectPackage:  "some-package",
				ProjectSources: []string{"not-exist"},
			},
			expectedError: "TypicalContext: Source 'not-exist' is not exist",
		},
		{
			Context: &typcore.Context{
				Descriptor:    validDescriptor,
				ProjectPackage: "some-package",
			},
		},
	}

	for _, tt := range testcases {
		err := common.Validate(tt.Context)
		if tt.expectedError == "" {
			require.NoError(t, err)
		} else {
			require.EqualError(t, err, tt.expectedError)
		}
	}

}
