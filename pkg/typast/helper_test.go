package typast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestIsPublic(t *testing.T) {
	testnames := []struct {
		TestName string
		Type     typast.Type
		Expected bool
	}{
		{
			Type:     &typast.FuncDecl{Name: "someFunc"},
			Expected: false,
		},
		{
			Type:     &typast.FuncDecl{Name: "SomeFunc"},
			Expected: true,
		},
	}
	for _, tt := range testnames {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, typast.IsPublic(tt.Type))
		})
	}
}

func Test_EqualFunc(t *testing.T) {
	testcases := []struct {
		TestName string
		TagName  string
		Annot    *typast.Annot
		Expected bool
	}{
		{
			TestName: "private function",
			TagName:  "@tagname",
			Annot: &typast.Annot{
				TagName: "@tagname",
				Decl:    &typast.Decl{Type: &typast.FuncDecl{Name: "someFunction"}},
			},
			Expected: false,
		},
		{
			TestName: "wrong tagname",
			TagName:  "@wrong",
			Annot: &typast.Annot{
				TagName: "@tagname",
				Decl:    &typast.Decl{Type: &typast.FuncDecl{Name: "SomeFunction"}},
			},
			Expected: false,
		},
		{
			TestName: "public function",
			TagName:  "@tagname",
			Annot: &typast.Annot{
				TagName: "@tagname",
				Decl:    &typast.Decl{Type: &typast.FuncDecl{Name: "SomeFunction"}},
			},
			Expected: true,
		},
		{
			TestName: "not function",
			TagName:  "@tagname",
			Annot: &typast.Annot{
				TagName: "@tagname",
				Decl: &typast.Decl{
					Type: &typast.InterfaceDecl{TypeDecl: typast.TypeDecl{Name: "SomeInterface"}},
				},
			},
			Expected: false,
		},
		{
			TestName: "method function",
			TagName:  "@tagname",
			Annot: &typast.Annot{
				TagName: "@tagname",
				Decl:    &typast.Decl{Type: &typast.FuncDecl{Name: "SomeMethod", Recv: &typast.FieldList{}}},
			},
			Expected: false,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, typast.EqualFunc(tt.Annot, tt.TagName))
		})
	}
}

func Test_EqualInterface(t *testing.T) {
	testcases := []struct {
		TestName string
		TagName  string
		Annot    *typast.Annot
		Expected bool
	}{
		{
			TestName: "private interface",
			TagName:  "@tagname",
			Annot: &typast.Annot{
				TagName: "@tagname",
				Decl: &typast.Decl{Type: &typast.InterfaceDecl{
					TypeDecl: typast.TypeDecl{Name: "someInterface"},
				}},
			},
			Expected: false,
		},
		{
			TestName: "wrong tagname",
			TagName:  "@wrong",
			Annot: &typast.Annot{
				TagName: "@tagname",
				Decl: &typast.Decl{Type: &typast.InterfaceDecl{
					TypeDecl: typast.TypeDecl{Name: "SomeInterface"},
				}},
			},
			Expected: false,
		},
		{
			TestName: "public interface",
			TagName:  "@tagname",
			Annot: &typast.Annot{
				TagName: "@tagname",
				Decl: &typast.Decl{Type: &typast.InterfaceDecl{
					TypeDecl: typast.TypeDecl{Name: "SomeInterface"},
				}},
			},
			Expected: true,
		},
		{
			TestName: "not struct",
			TagName:  "@tagname",
			Annot: &typast.Annot{
				TagName: "@tagname",
				Decl: &typast.Decl{Type: &typast.StructDecl{
					TypeDecl: typast.TypeDecl{Name: "SomeStruct"},
				}},
			},
			Expected: false,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, typast.EqualInterface(tt.Annot, tt.TagName))
		})
	}
}

func Test_EqualStruct(t *testing.T) {
	testcases := []struct {
		TestName string
		TagName  string
		Annot    *typast.Annot
		Expected bool
	}{
		{
			TestName: "private struct",
			TagName:  "@tagname",
			Annot: &typast.Annot{
				TagName: "@tagname",
				Decl: &typast.Decl{Type: &typast.StructDecl{
					TypeDecl: typast.TypeDecl{Name: "someStruct"},
				}},
			},
			Expected: false,
		},
		{
			TestName: "wrong tagname",
			TagName:  "@wrong",
			Annot: &typast.Annot{
				TagName: "@tagname",
				Decl: &typast.Decl{Type: &typast.StructDecl{
					TypeDecl: typast.TypeDecl{Name: "SomeStruct"},
				}},
			},
			Expected: false,
		},
		{
			TestName: "public struct",
			TagName:  "@tagname",
			Annot: &typast.Annot{
				TagName: "@tagname",
				Decl: &typast.Decl{Type: &typast.StructDecl{
					TypeDecl: typast.TypeDecl{Name: "SomeStruct"},
				}},
			},
			Expected: true,
		},
		{
			TestName: "not interface",
			TagName:  "@tagname",
			Annot: &typast.Annot{
				TagName: "@tagname",
				Decl: &typast.Decl{Type: &typast.InterfaceDecl{
					TypeDecl: typast.TypeDecl{Name: "SomeInterface"},
				}},
			},
			Expected: false,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, typast.EqualStruct(tt.Annot, tt.TagName))
		})
	}
}

func TestCtorAnnotation_CreateCtors(t *testing.T) {
	ctx := &typast.Context{
		Summary: &typast.Summary{
			Annots: []*typast.Annot{
				{
					TagName: "@ctor",
					Decl: &typast.Decl{
						Type: &typast.FuncDecl{Name: "NewObject"},
						File: typast.File{Package: "pkg", Path: "project/pkg/file.go"},
					},
				},
				{
					TagName:  "@ctor",
					TagParam: `name:"obj2"`,
					Decl: &typast.Decl{
						File: typast.File{Package: "pkg2", Path: "project/pkg2/file.go"},
						Type: &typast.FuncDecl{Name: "NewObject2"},
					},
				},
			},
		},
	}

	typgo.ProjectPkg = "github.com/typical-go/typical-go"

	annots, imports := typast.FindAnnot(ctx, "@ctor", typast.EqualFunc)
	require.Equal(t, map[string]string{
		"github.com/typical-go/typical-go/project/pkg":  "a",
		"github.com/typical-go/typical-go/project/pkg2": "b",
	}, imports)
	require.Equal(t, []*typast.Annot2{
		{
			Import:      "github.com/typical-go/typical-go/project/pkg",
			ImportAlias: "a",
			Annot: &typast.Annot{
				TagName: "@ctor",
				Decl: &typast.Decl{
					Type: &typast.FuncDecl{Name: "NewObject"},
					File: typast.File{Package: "pkg", Path: "project/pkg/file.go"},
				},
			},
		},
		{
			Import:      "github.com/typical-go/typical-go/project/pkg2",
			ImportAlias: "b",
			Annot: &typast.Annot{
				TagName:  "@ctor",
				TagParam: `name:"obj2"`,
				Decl: &typast.Decl{
					File: typast.File{Package: "pkg2", Path: "project/pkg2/file.go"},
					Type: &typast.FuncDecl{Name: "NewObject2"},
				},
			},
		},
	}, annots)
}
