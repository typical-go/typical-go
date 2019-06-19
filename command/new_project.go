package main

import (
	"github.com/typical-go/typical-go/appcontext"
	"github.com/typical-go/typical-go/command/stmt"
	"github.com/typical-go/typical-go/utility/oskit"
)

// NewProject new project
func NewProject(path string) (err error) {
	projectPath := oskit.GOPATH() + "/src/" + path
	typicalPath := projectPath + "/.typical"

	context := appcontext.Context{
		Name: "meh",
		Path: path,
	}

	err = Start(
		stmt.MakeDirectory(typicalPath),
		stmt.SaveContext(context, typicalPath+"/_context.json"),
	)
	return
}
