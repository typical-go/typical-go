package typcore_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/app"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestTypicalContext(t *testing.T) {
	os.MkdirAll("app/some_pkg", os.ModePerm)
	os.MkdirAll("pkg/some_lib", os.ModePerm)
	os.Create("app/some_pkg/some_file.go")
	os.Create("app/some_pkg/not_go.xxx")
	os.Create("pkg/some_lib/lib.go")
	ioutil.WriteFile("go.mod", []byte("module github.com/typical-go/typical-go\ngo 1.13"), 0644)
	defer func() {
		os.RemoveAll("app")
		os.RemoveAll("pkg")
		os.Remove("go.mod")
	}()

	ctx, err := typcore.CreateContext(&typcore.Descriptor{
		Name: "some-name",
		App:  app.New(),
		BuildTool: typbuildtool.Create(
			typbuildtool.StandardBuild(),
		),
	})

	require.NoError(t, err)

	// NOTE: ProjectPackage need to set manually because its value get from ldflags
	ctx.ProjectPackage = "some-package"

	require.NoError(t, common.Validate(ctx))
	require.Equal(t, "0.0.1", ctx.Version)
	require.Equal(t, []string{"app", "pkg"}, ctx.ProjectSources)
	require.Equal(t, []string{"app", "app/some_pkg", "pkg", "pkg/some_lib"}, ctx.ProjectDirs)
	require.Equal(t, []string{"app/some_pkg/some_file.go", "pkg/some_lib/lib.go"}, ctx.ProjectFiles)
}
