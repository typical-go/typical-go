package typannot_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typannot"
)

func TestCreateAnnotation(t *testing.T) {
	testcases := []struct {
		TestName         string
		Raw              string
		ExpectedTagName  string
		ExpectedTagAttrs string
	}{
		{
			TestName:        "tag only",
			Raw:             `@tag1`,
			ExpectedTagName: "@tag1",
		},
		{
			TestName:        "tag only with space",
			Raw:             `@tag2 extra1`,
			ExpectedTagName: "@tag2",
		},
		{
			TestName:         "with attribute",
			Raw:              `@tag3("name":"wire1")`,
			ExpectedTagName:  "@tag3",
			ExpectedTagAttrs: `"name":"wire1"`,
		},
		{
			TestName:         "there is space between tagname and attribute",
			Raw:              `@tag4 ("name":"wire1")`,
			ExpectedTagName:  "@tag4",
			ExpectedTagAttrs: `"name":"wire1"`,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			tagName, tagAttrs := typannot.ParseAnnot(tt.Raw)
			require.Equal(t, tt.ExpectedTagName, tagName)
			require.Equal(t, tt.ExpectedTagAttrs, tagAttrs)
		})
	}
}

func TestAnnot_CheckFunc(t *testing.T) {
	testcases := []struct {
		TestName string
		*typannot.Annot
		TagName  string
		Expected bool
	}{
		{
			Annot:    &typannot.Annot{TagName: "@tagname", Decl: &typannot.Decl{Type: &typannot.FuncType{}}},
			TagName:  "@tagname",
			Expected: true,
		},
		{
			Annot:    &typannot.Annot{TagName: "@tagname", Decl: &typannot.Decl{Type: &typannot.FuncType{}}},
			TagName:  "@tagname1",
			Expected: false,
		},
		{
			Annot:    &typannot.Annot{TagName: "@tagname", Decl: &typannot.Decl{Type: &typannot.StructType{}}},
			TagName:  "@tagname",
			Expected: false,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, tt.CheckFunc(tt.TagName))
		})
	}
}

func TestAnnot_CheckStruct(t *testing.T) {
	testcases := []struct {
		TestName string
		*typannot.Annot
		TagName  string
		Expected bool
	}{
		{
			Annot:    &typannot.Annot{TagName: "@tagname", Decl: &typannot.Decl{Type: &typannot.StructType{}}},
			TagName:  "@tagname",
			Expected: true,
		},
		{
			Annot:    &typannot.Annot{TagName: "@tagname", Decl: &typannot.Decl{Type: &typannot.StructType{}}},
			TagName:  "@tagname1",
			Expected: false,
		},
		{
			Annot:    &typannot.Annot{TagName: "@tagname", Decl: &typannot.Decl{Type: &typannot.FuncType{}}},
			TagName:  "@tagname",
			Expected: false,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, tt.CheckStruct(tt.TagName))
		})
	}
}

func TestAnnot_CheckInterface(t *testing.T) {
	testcases := []struct {
		TestName string
		*typannot.Annot
		TagName  string
		Expected bool
	}{
		{
			Annot:    &typannot.Annot{TagName: "@tagname", Decl: &typannot.Decl{Type: &typannot.InterfaceType{}}},
			TagName:  "@tagname",
			Expected: true,
		},
		{
			Annot:    &typannot.Annot{TagName: "@tagname", Decl: &typannot.Decl{Type: &typannot.InterfaceType{}}},
			TagName:  "@tagname1",
			Expected: false,
		},
		{
			Annot:    &typannot.Annot{TagName: "@tagname", Decl: &typannot.Decl{Type: &typannot.FuncType{}}},
			TagName:  "@tagname",
			Expected: false,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, tt.CheckInterface(tt.TagName))
		})
	}
}
