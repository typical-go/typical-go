package typannot_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typannot"
)

func TestFindAnnot(t *testing.T) {
	tag1 := &typannot.Annot{TagName: "@some-tag", Decl: &typannot.Decl{DeclType: &typannot.FuncDecl{}}}
	tag2 := &typannot.Annot{TagName: "@some-tag", Decl: &typannot.Decl{DeclType: &typannot.FuncDecl{}}}
	tag3 := &typannot.Annot{TagName: "@other", Decl: &typannot.Decl{DeclType: &typannot.FuncDecl{}}}
	tag4 := &typannot.Annot{TagName: "@some-tag", Decl: &typannot.Decl{DeclType: &typannot.StructDecl{}}}
	tag5 := &typannot.Annot{TagName: "@some-tag", Decl: &typannot.Decl{DeclType: &typannot.StructDecl{}}}
	tag6 := &typannot.Annot{TagName: "@other", Decl: &typannot.Decl{DeclType: &typannot.StructDecl{}}}
	tag7 := &typannot.Annot{TagName: "@some-tag", Decl: &typannot.Decl{DeclType: &typannot.InterfaceDecl{}}}
	tag8 := &typannot.Annot{TagName: "@some-tag", Decl: &typannot.Decl{DeclType: &typannot.InterfaceDecl{}}}
	tag9 := &typannot.Annot{TagName: "@other", Decl: &typannot.Decl{DeclType: &typannot.InterfaceDecl{}}}

	c := &typannot.Summary{
		Annots: []*typannot.Annot{tag1, tag2, tag3, tag4, tag5, tag6, tag7, tag8, tag9},
	}

	require.Equal(t, []*typannot.Annot{tag1, tag2}, c.FindAnnotByFunc("@some-tag"))
	require.Equal(t, []*typannot.Annot{tag4, tag5}, c.FindAnnotByStruct("@some-tag"))
	require.Equal(t, []*typannot.Annot{tag7, tag8}, c.FindAnnotByInterface("@some-tag"))

}
