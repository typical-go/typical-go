package typcore_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/wrapper"
)

func TestTypicalContext(t *testing.T) {
	os.MkdirAll("wrapper/some_pkg", os.ModePerm)
	os.MkdirAll("pkg/some_lib", os.ModePerm)
	os.Create("wrapper/some_pkg/some_file.go")
	os.Create("wrapper/some_pkg/not_go.xxx")
	os.Create("pkg/some_lib/lib.go")
	ioutil.WriteFile("go.mod", []byte("module github.com/typical-go/typical-go\ngo 1.13"), 0644)
	defer func() {
		os.RemoveAll("wrapper")
		os.RemoveAll("pkg")
		os.Remove("go.mod")
	}()

	ctx, err := typcore.CreateContext(&typcore.Descriptor{
		Name: "some-name",
		App:  wrapper.New(),
		BuildTool: typbuildtool.BuildSequences(
			typbuildtool.StandardBuild(),
		),
	})

	require.NoError(t, err)

	// NOTE: ProjectPackage need to set manually because its value get from ldflags
	ctx.ProjectPkg = "some-package"

	require.NoError(t, common.Validate(ctx))
	require.Equal(t, "0.0.1", ctx.Version)
	require.Equal(t, []string{"wrapper", "pkg"}, ctx.AppSources)
	require.Equal(t, []string{"wrapper", "wrapper/some_pkg", "pkg", "pkg/some_lib"}, ctx.AppDirs)
	require.Equal(t, []string{"wrapper/some_pkg/some_file.go", "pkg/some_lib/lib.go"}, ctx.AppFiles)
}
