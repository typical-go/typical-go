package typast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typast"
)

func TestIsPublic(t *testing.T) {
	testnames := []struct {
		TestName  string
		Directive *typast.Directive
		CheckFn   func(*typast.Directive) bool
		Expected  bool
	}{
		{
			TestName:  "IsPublic: function name start with lower case",
			Directive: &typast.Directive{Decl: &typast.Decl{Type: &typast.FuncDecl{Name: "someFunc"}}},
			CheckFn:   typast.IsPublic,
			Expected:  false,
		},
		{
			TestName:  "IsPublic: function name start with upper case",
			Directive: &typast.Directive{Decl: &typast.Decl{Type: &typast.FuncDecl{Name: "SomeFunc"}}},
			CheckFn:   typast.IsPublic,
			Expected:  true,
		},
		{
			TestName:  "IsFunc: type is function",
			Directive: &typast.Directive{Decl: &typast.Decl{Type: &typast.FuncDecl{Name: "SomeFunc"}}},
			CheckFn:   typast.IsFunc,
			Expected:  true,
		},
		{
			TestName:  "IsFunc: type is interface",
			Directive: &typast.Directive{Decl: &typast.Decl{Type: &typast.InterfaceDecl{}}},
			CheckFn:   typast.IsFunc,
			Expected:  false,
		},
		{
			TestName: "IsFunc: type is method",
			Directive: &typast.Directive{
				Decl: &typast.Decl{
					Type: &typast.FuncDecl{Name: "SomeFunc", Recv: &typast.FieldList{}},
				},
			},
			CheckFn:  typast.IsFunc,
			Expected: false,
		},
		{
			TestName:  "IsMethod: type is function",
			Directive: &typast.Directive{Decl: &typast.Decl{Type: &typast.FuncDecl{Name: "SomeFunc"}}},
			CheckFn:   typast.IsMethod,
			Expected:  false,
		},
		{
			TestName:  "IsMethod: type is interface",
			Directive: &typast.Directive{Decl: &typast.Decl{Type: &typast.InterfaceDecl{}}},
			CheckFn:   typast.IsMethod,
			Expected:  false,
		},
		{
			TestName: "IsMethod: type is method",
			Directive: &typast.Directive{
				Decl: &typast.Decl{
					Type: &typast.FuncDecl{Name: "SomeFunc", Recv: &typast.FieldList{}},
				},
			},
			CheckFn:  typast.IsMethod,
			Expected: true,
		},
		{
			TestName:  "IsInterface: type is interface",
			Directive: &typast.Directive{Decl: &typast.Decl{Type: &typast.InterfaceDecl{}}},
			CheckFn:   typast.IsInterface,
			Expected:  true,
		},
		{
			TestName:  "IsInterface: type is function",
			Directive: &typast.Directive{Decl: &typast.Decl{Type: &typast.FuncDecl{Name: "SomeFunc"}}},
			CheckFn:   typast.IsInterface,
			Expected:  false,
		},
		{
			TestName:  "IsStruct: type is interface",
			Directive: &typast.Directive{Decl: &typast.Decl{Type: &typast.StructDecl{}}},
			CheckFn:   typast.IsStruct,
			Expected:  true,
		},
		{
			TestName:  "IsStruct: type is function",
			Directive: &typast.Directive{Decl: &typast.Decl{Type: &typast.FuncDecl{Name: "SomeFunc"}}},
			CheckFn:   typast.IsStruct,
			Expected:  false,
		},
	}
	for _, tt := range testnames {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, tt.CheckFn(tt.Directive))
		})
	}
}
