package walker

import "go/ast"

// FuncDeclListener is listen function declarion event
type FuncDeclListener interface {
	IsAction(*FuncDeclEvent) bool
	ActionPerformed(*FuncDeclEvent) error
}

// FuncDeclEvent is an event when function declared
type FuncDeclEvent struct {
	*ast.FuncDecl
	Name string
	File *ast.File
}
