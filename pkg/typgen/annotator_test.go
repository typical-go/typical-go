package typgen_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgen"
)

func TestFilterFunc(t *testing.T) {
	testcases := []struct {
		TestName   string
		Fn         func(d *typgen.Annotation) bool
		Annotation *typgen.Annotation
		Expected   bool
	}{
		{
			TestName:   "PublicFilter: function name start with lower case",
			Annotation: &typgen.Annotation{Decl: &typgen.Decl{Type: &typgen.Function{Name: "someFunc"}}},
			Fn:         typgen.IsPublic,
			Expected:   false,
		},
		{
			TestName:   "PublicFilter: function name start with upper case",
			Annotation: &typgen.Annotation{Decl: &typgen.Decl{Type: &typgen.Function{Name: "SomeFunc"}}},
			Fn:         typgen.IsPublic,
			Expected:   true,
		},
		{
			TestName:   "FuncFilter: type is function",
			Annotation: &typgen.Annotation{Decl: &typgen.Decl{Type: &typgen.Function{Name: "SomeFunc"}}},
			Fn:         typgen.IsFunc,
			Expected:   true,
		},
		{
			TestName:   "FuncFilter: type is interface",
			Annotation: &typgen.Annotation{Decl: &typgen.Decl{Type: &typgen.Interface{}}},
			Fn:         typgen.IsFunc,
			Expected:   false,
		},
		{
			TestName: "FuncFilter: type is method",
			Annotation: &typgen.Annotation{
				Decl: &typgen.Decl{
					Type: &typgen.Function{Name: "SomeFunc", Recv: []*typgen.Field{{}}},
				},
			},
			Fn:       typgen.IsFunc,
			Expected: false,
		},
		{
			TestName:   "InterfaceFilter: type is interface",
			Annotation: &typgen.Annotation{Decl: &typgen.Decl{Type: &typgen.Interface{}}},
			Fn:         typgen.IsInterface,
			Expected:   true,
		},
		{
			TestName:   "InterfaceFilter: type is function",
			Annotation: &typgen.Annotation{Decl: &typgen.Decl{Type: &typgen.Function{Name: "SomeFunc"}}},
			Fn:         typgen.IsInterface,
			Expected:   false,
		},
		{
			TestName:   "StructFilter: type is interface",
			Annotation: &typgen.Annotation{Decl: &typgen.Decl{Type: &typgen.Struct{}}},
			Fn:         typgen.IsStruct,
			Expected:   true,
		},
		{
			TestName:   "StructFilter: type is function",
			Annotation: &typgen.Annotation{Decl: &typgen.Decl{Type: &typgen.Function{Name: "SomeFunc"}}},
			Fn:         typgen.IsStruct,
			Expected:   false,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, tt.Fn(tt.Annotation))
		})
	}
}
