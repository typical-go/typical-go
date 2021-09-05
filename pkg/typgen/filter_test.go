package typgen_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgen"
)

func TestFilter(t *testing.T) {
	testcases := []struct {
		TestName  string
		Filter    typgen.Filter
		Directive *typgen.Directive
		Expected  bool
	}{
		{
			TestName:  "NewFilter",
			Directive: &typgen.Directive{Decl: &typgen.Decl{Type: &typgen.Function{Name: "someFunc"}}},
			Filter:    typgen.NewFilter(func(d *typgen.Directive) bool { return true }),
			Expected:  true,
		},
		{
			TestName:  "NewFilter",
			Directive: &typgen.Directive{TagName: "@tagname1"},
			Filter:    typgen.TagNameFilter{"@tagname1"},
			Expected:  true,
		},
		{
			TestName:  "NewFilter",
			Directive: &typgen.Directive{TagName: "@tagname2"},
			Filter:    typgen.TagNameFilter{"@tagname1"},
			Expected:  false,
		},
		{
			TestName:  "NewFilter",
			Directive: &typgen.Directive{Decl: &typgen.Decl{Type: &typgen.Function{Name: "someFunc"}}},
			Filter:    typgen.NewFilter(func(d *typgen.Directive) bool { return true }),
			Expected:  true,
		},
		{
			TestName:  "PublicFilter: function name start with lower case",
			Directive: &typgen.Directive{Decl: &typgen.Decl{Type: &typgen.Function{Name: "someFunc"}}},
			Filter:    &typgen.PublicFilter{},
			Expected:  false,
		},
		{
			TestName:  "PublicFilter: function name start with upper case",
			Directive: &typgen.Directive{Decl: &typgen.Decl{Type: &typgen.Function{Name: "SomeFunc"}}},
			Filter:    &typgen.PublicFilter{},
			Expected:  true,
		},
		{
			TestName:  "FuncFilter: type is function",
			Directive: &typgen.Directive{Decl: &typgen.Decl{Type: &typgen.Function{Name: "SomeFunc"}}},
			Filter:    &typgen.FuncFilter{},
			Expected:  true,
		},
		{
			TestName:  "FuncFilter: type is interface",
			Directive: &typgen.Directive{Decl: &typgen.Decl{Type: &typgen.Interface{}}},
			Filter:    &typgen.FuncFilter{},
			Expected:  false,
		},
		{
			TestName: "FuncFilter: type is method",
			Directive: &typgen.Directive{
				Decl: &typgen.Decl{
					Type: &typgen.Function{Name: "SomeFunc", Recv: []*typgen.Field{{}}},
				},
			},
			Filter:   &typgen.FuncFilter{},
			Expected: false,
		},
		{
			TestName:  "InterfaceFilter: type is interface",
			Directive: &typgen.Directive{Decl: &typgen.Decl{Type: &typgen.Interface{}}},
			Filter:    &typgen.InterfaceFilter{},
			Expected:  true,
		},
		{
			TestName:  "InterfaceFilter: type is function",
			Directive: &typgen.Directive{Decl: &typgen.Decl{Type: &typgen.Function{Name: "SomeFunc"}}},
			Filter:    &typgen.InterfaceFilter{},
			Expected:  false,
		},
		{
			TestName:  "StructFilter: type is interface",
			Directive: &typgen.Directive{Decl: &typgen.Decl{Type: &typgen.Struct{}}},
			Filter:    &typgen.StructFilter{},
			Expected:  true,
		},
		{
			TestName:  "StructFilter: type is function",
			Directive: &typgen.Directive{Decl: &typgen.Decl{Type: &typgen.Function{Name: "SomeFunc"}}},
			Filter:    &typgen.StructFilter{},
			Expected:  false,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, tt.Filter.IsAllowed(tt.Directive))
		})
	}
}
