package typast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typast"
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
