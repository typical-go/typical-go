package typast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typast"
)

func TestFilter(t *testing.T) {
	testcases := []struct {
		TestName  string
		Filter    typast.Filter
		Directive *typast.Directive
		Expected  bool
	}{
		{
			TestName:  "NewFilter",
			Directive: &typast.Directive{Decl: &typast.Decl{Type: &typast.FuncDecl{Name: "someFunc"}}},
			Filter:    typast.NewFilter(func(d *typast.Directive) bool { return true }),
			Expected:  true,
		},
		{
			TestName:  "NewFilter",
			Directive: &typast.Directive{TagName: "@tagname1"},
			Filter:    typast.TagNameFilter{"@tagname1"},
			Expected:  true,
		},
		{
			TestName:  "NewFilter",
			Directive: &typast.Directive{TagName: "@tagname2"},
			Filter:    typast.TagNameFilter{"@tagname1"},
			Expected:  false,
		},
		{
			TestName:  "NewFilter",
			Directive: &typast.Directive{Decl: &typast.Decl{Type: &typast.FuncDecl{Name: "someFunc"}}},
			Filter:    typast.NewFilter(func(d *typast.Directive) bool { return true }),
			Expected:  true,
		},
		{
			TestName:  "PublicFilter: function name start with lower case",
			Directive: &typast.Directive{Decl: &typast.Decl{Type: &typast.FuncDecl{Name: "someFunc"}}},
			Filter:    &typast.PublicFilter{},
			Expected:  false,
		},
		{
			TestName:  "PublicFilter: function name start with upper case",
			Directive: &typast.Directive{Decl: &typast.Decl{Type: &typast.FuncDecl{Name: "SomeFunc"}}},
			Filter:    &typast.PublicFilter{},
			Expected:  true,
		},
		{
			TestName:  "FuncFilter: type is function",
			Directive: &typast.Directive{Decl: &typast.Decl{Type: &typast.FuncDecl{Name: "SomeFunc"}}},
			Filter:    &typast.FuncFilter{},
			Expected:  true,
		},
		{
			TestName:  "FuncFilter: type is interface",
			Directive: &typast.Directive{Decl: &typast.Decl{Type: &typast.InterfaceDecl{}}},
			Filter:    &typast.FuncFilter{},
			Expected:  false,
		},
		{
			TestName: "FuncFilter: type is method",
			Directive: &typast.Directive{
				Decl: &typast.Decl{
					Type: &typast.FuncDecl{Name: "SomeFunc", Recv: &typast.FieldList{}},
				},
			},
			Filter:   &typast.FuncFilter{},
			Expected: false,
		},
		{
			TestName:  "InterfaceFilter: type is interface",
			Directive: &typast.Directive{Decl: &typast.Decl{Type: &typast.InterfaceDecl{}}},
			Filter:    &typast.InterfaceFilter{},
			Expected:  true,
		},
		{
			TestName:  "InterfaceFilter: type is function",
			Directive: &typast.Directive{Decl: &typast.Decl{Type: &typast.FuncDecl{Name: "SomeFunc"}}},
			Filter:    &typast.InterfaceFilter{},
			Expected:  false,
		},
		{
			TestName:  "StructFilter: type is interface",
			Directive: &typast.Directive{Decl: &typast.Decl{Type: &typast.StructDecl{}}},
			Filter:    &typast.StructFilter{},
			Expected:  true,
		},
		{
			TestName:  "StructFilter: type is function",
			Directive: &typast.Directive{Decl: &typast.Decl{Type: &typast.FuncDecl{Name: "SomeFunc"}}},
			Filter:    &typast.StructFilter{},
			Expected:  false,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, tt.Filter.IsAllowed(tt.Directive))
		})
	}
}
