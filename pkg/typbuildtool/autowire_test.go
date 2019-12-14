package typbuildtool_test

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typbuildtool/walker"
)

func TestAutowire(t *testing.T) {
	testcases := []struct {
		typbuildtool.Autowires
		event    *walker.FuncDeclEvent
		autowire []string
	}{
		{
			event: &walker.FuncDeclEvent{
				Name:     "SomeFunction",
				FuncDecl: &ast.FuncDecl{},
				File:     &ast.File{Name: &ast.Ident{Name: "pkg"}},
			},
		},
		{
			event: &walker.FuncDeclEvent{
				FuncDecl: &ast.FuncDecl{},
				Name:     "NewSomething",
				File:     &ast.File{Name: ast.NewIdent("pkg")},
			},
			autowire: []string{"pkg.NewSomething"},
		},
		{
			event: &walker.FuncDeclEvent{
				FuncDecl: funcDeclWithComment("some doc"),
				Name:     "SomeFunction",
				File:     &ast.File{Name: ast.NewIdent("pkg")},
			},
		},
		{
			event: &walker.FuncDeclEvent{
				FuncDecl: funcDeclWithComment("some doc [autowire]"),
				Name:     "SomeFunction",
				File:     &ast.File{Name: ast.NewIdent("pkg")},
			},
			autowire: []string{"pkg.SomeFunction"},
		},
		{
			event: &walker.FuncDeclEvent{
				FuncDecl: funcDeclWithComment("some doc [Autowire]"),
				Name:     "SomeFunction",
				File:     &ast.File{Name: ast.NewIdent("pkg")},
			},
			autowire: []string{"pkg.SomeFunction"},
		},
		{
			event: &walker.FuncDeclEvent{
				FuncDecl: funcDeclWithComment("some doc [AUTOWIRE]"),
				Name:     "SomeFunction",
				File:     &ast.File{Name: ast.NewIdent("pkg")},
			},
			autowire: []string{"pkg.SomeFunction"},
		},
		{
			event: &walker.FuncDeclEvent{
				FuncDecl: funcDeclWithComment("some doc [nowire]"),
				Name:     "NewSomeFunction",
				File:     &ast.File{Name: ast.NewIdent("pkg")},
			},
		},
		{
			event: &walker.FuncDeclEvent{
				FuncDecl: funcDeclWithComment("some doc [nowire][autowire]"),
				Name:     "NewSomeFunction",
				File:     &ast.File{Name: ast.NewIdent("pkg")},
			},
		},
		{
			event: &walker.FuncDeclEvent{
				FuncDecl: funcDeclWithComment("some doc [nowire][autowire]"),
				Name:     "SomeFunction",
				File:     &ast.File{Name: ast.NewIdent("pkg")},
			},
			autowire: []string{"pkg.SomeFunction"},
		},
	}
	for _, tt := range testcases {
		require.NoError(t, tt.OnFuncDecl(tt.event))
		require.EqualValues(t, tt.autowire, tt.Autowires)
	}
}

func funcDeclWithComment(text string) *ast.FuncDecl {
	return &ast.FuncDecl{
		Doc: astComment(text),
	}
}

func astComment(text string) *ast.CommentGroup {
	return &ast.CommentGroup{
		List: []*ast.Comment{
			{Text: text},
		},
	}
}
