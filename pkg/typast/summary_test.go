package typast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typast"
)

func TestFindAnnot(t *testing.T) {
	tag1 := &typast.Annot{TagName: "@some-tag", Decl: &typast.Decl{Type: &typast.FuncDecl{}}}
	tag2 := &typast.Annot{TagName: "@some-tag", Decl: &typast.Decl{Type: &typast.FuncDecl{}}}
	tag3 := &typast.Annot{TagName: "@other", Decl: &typast.Decl{Type: &typast.FuncDecl{}}}
	tag4 := &typast.Annot{TagName: "@some-tag", Decl: &typast.Decl{Type: &typast.StructDecl{}}}
	tag5 := &typast.Annot{TagName: "@some-tag", Decl: &typast.Decl{Type: &typast.StructDecl{}}}
	tag6 := &typast.Annot{TagName: "@other", Decl: &typast.Decl{Type: &typast.StructDecl{}}}
	tag7 := &typast.Annot{TagName: "@some-tag", Decl: &typast.Decl{Type: &typast.InterfaceDecl{}}}
	tag8 := &typast.Annot{TagName: "@some-tag", Decl: &typast.Decl{Type: &typast.InterfaceDecl{}}}
	tag9 := &typast.Annot{TagName: "@other", Decl: &typast.Decl{Type: &typast.InterfaceDecl{}}}

	c := &typast.Summary{
		Annots: []*typast.Annot{tag1, tag2, tag3, tag4, tag5, tag6, tag7, tag8, tag9},
	}

	require.Equal(t,
		[]*typast.Annot{tag1, tag2},
		c.FindAnnot(
			func(a *typast.Annot) bool {
				_, ok := a.Type.(*typast.FuncDecl)
				return ok && a.TagName == "@some-tag"
			},
		),
	)

}
