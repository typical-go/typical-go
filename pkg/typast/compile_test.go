package typast_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typast"
)

var (
	someInterfaceDecl = &typast.Decl{
		File: typast.File{
			Path:    "sample_test.go",
			Package: "typast_test",
		},
		DeclType: &typast.InterfaceDecl{
			TypeDecl: typast.TypeDecl{Name: "sampleInterface"},
		},
	}

	someStructDecl = &typast.Decl{
		File: typast.File{
			Path:    "sample_test.go",
			Package: "typast_test",
		},
		DeclType: &typast.StructDecl{
			TypeDecl: typast.TypeDecl{
				GenDecl: typast.GenDecl{
					Docs: []string{
						"// sampleStruct",
						"// @tag1",
						"// @tag2 (key1:\"\", key2: \"\", key3:\"value3\")",
					},
				},
				Name: "sampleStruct",
			},
			Fields: []*typast.Field{
				{
					Names:     []string{"sampleInt"},
					Type:      "int",
					StructTag: reflect.StructTag(`default:"value1"`),
				},
				{
					Names:     []string{"sampleString"},
					Type:      "string",
					StructTag: reflect.StructTag(`default:"value2"`),
				},
			},
		},
	}

	someFunctionDecl = &typast.Decl{
		File: typast.File{
			Path:    "sample_test.go",
			Package: "typast_test",
		},
		DeclType: &typast.FuncDecl{
			Name: "sampleFunction",
			Params: &typast.FieldList{
				List: []*typast.Field{
					{Names: []string{"param1"}, Type: "int"},
					{Names: []string{"param2"}, Type: "int"},
				},
			},
		},
	}

	someFunctionDecl2 = &typast.Decl{
		File: typast.File{
			Path:    "sample_test.go",
			Package: "typast_test",
		},
		DeclType: &typast.FuncDecl{
			Name:   "sampleFunction2",
			Params: &typast.FieldList{},
			Docs: []string{
				"// GetWriter to get writer to greet the world",
				"// @ctor",
			},
		},
	}

	someInterface2Decl = &typast.Decl{
		File: typast.File{
			Path:    "sample_test.go",
			Package: "typast_test",
		},
		DeclType: &typast.InterfaceDecl{
			TypeDecl: typast.TypeDecl{
				Name: "sampleInterface2",
				Docs: []string{"// @tag3"},
			},
		},
	}

	someStruct2Decl = &typast.Decl{
		File: typast.File{
			Path:    "sample_test.go",
			Package: "typast_test",
		},
		DeclType: &typast.StructDecl{
			TypeDecl: typast.TypeDecl{
				Name: "sampleStruct2",
				Docs: []string{
					"// sampleStruct2 asdf",
					"// @tag4",
				},
			},
		},
	}
	someStruct3Decl = &typast.Decl{
		File: typast.File{
			Path:    "sample_test.go",
			Package: "typast_test",
		},
		DeclType: &typast.StructDecl{
			TypeDecl: typast.TypeDecl{Name: "sampleStruct3"},
			Fields: []*typast.Field{
				{Names: []string{"Name"}, Type: "string"},
				{Names: []string{"Address"}, Type: "string"},
			},
		},
	}

	someMethod = &typast.Decl{
		File: typast.File{
			Path:    "sample_test.go",
			Package: "typast_test",
		},
		DeclType: &typast.FuncDecl{
			Name: "someMethod",
			Recv: &typast.FieldList{
				List: []*typast.Field{
					{Names: []string{"s"}, Type: "*sampleStruct3"},
				},
			},
			Params: &typast.FieldList{},
		},
	}
)

func TestCompile(t *testing.T) {
	summary, err := typast.Compile("sample_test.go")
	require.NoError(t, err)

	require.EqualValues(t, someInterfaceDecl, summary.Decls[0])
	require.EqualValues(t, someStructDecl, summary.Decls[1])
	require.EqualValues(t, someFunctionDecl, summary.Decls[2])
	require.EqualValues(t, someFunctionDecl2, summary.Decls[3])
	require.EqualValues(t, someInterface2Decl, summary.Decls[4])
	require.EqualValues(t, someStruct2Decl, summary.Decls[5])
	require.EqualValues(t, someStruct3Decl, summary.Decls[6])
	require.EqualValues(t, someMethod, summary.Decls[7])

	// require.EqualValues(t, []*typast.Annot{
	// 	{
	// 		Decl:    someStructDecl,
	// 		TagName: "@tag1",
	// 	},
	// 	{
	// 		Decl:     someStructDecl,
	// 		TagName:  "@tag2",
	// 		TagParam: `key1:"", key2: "", key3:"value3"`,
	// 	},
	// 	{
	// 		Decl:    someFunctionDecl2,
	// 		TagName: "@ctor",
	// 	},
	// 	{
	// 		Decl:    someInterface2Decl,
	// 		TagName: "@tag3",
	// 	},
	// 	{
	// 		Decl:    someStruct2Decl,
	// 		TagName: "@tag4",
	// 	},
	// }, summary.Annots)
}

func TestCompile_FileNotFound(t *testing.T) {
	_, err := typast.Compile("not_found.go")
	require.EqualError(t, err, "open not_found.go: no such file or directory")
}

func TestWalk(t *testing.T) {
	os.MkdirAll("wrapper/some_pkg", os.ModePerm)
	os.MkdirAll("pkg/some_lib", os.ModePerm)
	os.Create("wrapper/some_pkg/some_file.go")
	os.Create("wrapper/some_pkg/not_go.xxx")
	os.Create("pkg/some_lib/lib.go")
	defer func() {
		os.RemoveAll("wrapper")
		os.RemoveAll("pkg")
	}()

	dirs, files := typast.Walk([]string{"pkg", "wrapper"})
	require.Equal(t, []string{"pkg", "pkg/some_lib", "wrapper", "wrapper/some_pkg"}, dirs)
	require.Equal(t, []string{"pkg/some_lib/lib.go", "wrapper/some_pkg/some_file.go"}, files)
}
