package typannot_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typannot"
)

func TestContext_CreateImports(t *testing.T) {
	testcases := []struct {
		TestName string
		*typannot.Context
		ProjectPkg string
		More       []string
		Expected   []string
	}{
		{
			Context: &typannot.Context{
				Dirs: []string{"dir1", "dir2"},
			},
			ProjectPkg: "myproject",
			More:       []string{"github.com/x/x"},
			Expected: []string{
				"myproject/dir1",
				"myproject/dir2",
				"github.com/x/x",
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t,
				tt.Expected,
				tt.CreateImports(tt.ProjectPkg, tt.More...),
			)
		})
	}
}

func TestWalkLayout(t *testing.T) {
	os.MkdirAll("wrapper/some_pkg", os.ModePerm)
	os.MkdirAll("pkg/some_lib", os.ModePerm)
	os.Create("wrapper/some_pkg/some_file.go")
	os.Create("wrapper/some_pkg/not_go.xxx")
	os.Create("pkg/some_lib/lib.go")
	defer func() {
		os.RemoveAll("wrapper")
		os.RemoveAll("pkg")
	}()

	dirs, files := typannot.WalkLayout([]string{"pkg", "wrapper"})
	require.Equal(t, []string{"pkg", "pkg/some_lib", "wrapper", "wrapper/some_pkg"}, dirs)
	require.Equal(t, []string{"pkg/some_lib/lib.go", "wrapper/some_pkg/some_file.go"}, files)
}

func TestFindAnnot(t *testing.T) {
	tag1 := &typannot.Annot{TagName: "@some-tag", Decl: &typannot.Decl{Type: &typannot.FuncType{}}}
	tag2 := &typannot.Annot{TagName: "@some-tag", Decl: &typannot.Decl{Type: &typannot.FuncType{}}}
	tag3 := &typannot.Annot{TagName: "@other", Decl: &typannot.Decl{Type: &typannot.FuncType{}}}
	tag4 := &typannot.Annot{TagName: "@some-tag", Decl: &typannot.Decl{Type: &typannot.StructType{}}}
	tag5 := &typannot.Annot{TagName: "@some-tag", Decl: &typannot.Decl{Type: &typannot.StructType{}}}
	tag6 := &typannot.Annot{TagName: "@other", Decl: &typannot.Decl{Type: &typannot.StructType{}}}
	tag7 := &typannot.Annot{TagName: "@some-tag", Decl: &typannot.Decl{Type: &typannot.InterfaceType{}}}
	tag8 := &typannot.Annot{TagName: "@some-tag", Decl: &typannot.Decl{Type: &typannot.InterfaceType{}}}
	tag9 := &typannot.Annot{TagName: "@other", Decl: &typannot.Decl{Type: &typannot.InterfaceType{}}}

	c := &typannot.Context{
		ASTStore: &typannot.ASTStore{
			Annots: []*typannot.Annot{tag1, tag2, tag3, tag4, tag5, tag6, tag7, tag8, tag9},
		},
	}

	require.Equal(t, []*typannot.Annot{tag1, tag2}, c.FindAnnotByFunc("@some-tag"))
	require.Equal(t, []*typannot.Annot{tag4, tag5}, c.FindAnnotByStruct("@some-tag"))
	require.Equal(t, []*typannot.Annot{tag7, tag8}, c.FindAnnotByInterface("@some-tag"))

}
