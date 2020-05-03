package typannot_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typast"
)

func TestGetMock(t *testing.T) {
	os.Mkdir("sample", 0777)
	os.Create("sample/source.go")
	defer os.RemoveAll("sample")

	store := &typast.ASTStore{
		Annots: []*typast.Annotation{
			{
				Decl: &typast.Decl{
					Path: "mock.go",
					Name: "someInterface",
					Pkg:  "somePkg",
					Type: typast.Interface,
				},
				TagName: "mock",
			},
			{
				Decl: &typast.Decl{
					Path: "sample/source.go",
					Name: "someInterface2",
					Pkg:  "sample",
					Type: typast.Interface,
				},
				TagName: "mock",
			},
		},
	}

	require.Equal(t, []*typannot.Mock{
		{
			Dir:    ".",
			Pkg:    "somePkg",
			Source: "someInterface",
			Parent: "",
		},
		{
			Dir:    "sample",
			Pkg:    "sample",
			Source: "someInterface2",
			Parent: "",
		},
	}, typannot.GetMock(store))

}
