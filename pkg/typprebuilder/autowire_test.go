package typprebuilder_test

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typprebuilder"

	"github.com/typical-go/typical-go/pkg/typprebuilder/walker"
)

func TestAutowire_IsAction(t *testing.T) {
	testcases := []struct {
		typprebuilder.Autowires
		event    *walker.FuncDeclEvent
		isAction bool
	}{
		{
			event: &walker.FuncDeclEvent{
				FuncDecl: &ast.FuncDecl{},
			},
			isAction: false,
		},
		{
			event: &walker.FuncDeclEvent{
				FuncDecl: &ast.FuncDecl{},
				Name:     "NewSomething",
			},
			isAction: true,
		},
		{
			event: &walker.FuncDeclEvent{
				FuncDecl: funcDeclWithComment("some doc"),
				Name:     "SomeFunction",
			},
			isAction: false,
		},
		{
			event: &walker.FuncDeclEvent{
				FuncDecl: funcDeclWithComment("some doc [autowire]"),
				Name:     "SomeFunction",
			},
			isAction: true,
		},
		{
			event: &walker.FuncDeclEvent{
				FuncDecl: funcDeclWithComment("some doc [Autowire]"),
				Name:     "SomeFunction",
			},
			isAction: true,
		},
		{
			event: &walker.FuncDeclEvent{
				FuncDecl: funcDeclWithComment("some doc [AUTOWIRE]"),
				Name:     "SomeFunction",
			},
			isAction: true,
		},
		{
			event: &walker.FuncDeclEvent{
				FuncDecl: funcDeclWithComment("some doc [nowire]"),
				Name:     "NewSomeFunction",
			},
			isAction: false,
		},
		{
			event: &walker.FuncDeclEvent{
				FuncDecl: funcDeclWithComment("some doc [nowire][autowire]"),
				Name:     "NewSomeFunction",
			},
			isAction: false,
		},
		{
			event: &walker.FuncDeclEvent{
				FuncDecl: funcDeclWithComment("some doc [nowire][autowire]"),
				Name:     "SomeFunction",
			},
			isAction: true,
		},
	}
	for _, tt := range testcases {
		require.Equal(t, tt.isAction, tt.IsAction(tt.event))
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
