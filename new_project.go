package main

import (
	"github.com/typical-go/typical-go/command"
	"github.com/typical-go/typical-go/context"
	"github.com/typical-go/typical-go/executor"
	"github.com/typical-go/typical-go/utility/oskit"
)

// NewProject new project
func NewProject(path string) (err error) {
	projectPath := oskit.GOPATH() + "/src/" + path
	typicalPath := projectPath + "/.typical"

	context := context.Context{
		Name: "meh",
		Path: path,
	}

	err = executor.Run(
		command.MakeDirectory(typicalPath),
		command.SaveMetadataContext(context, typicalPath+"/_context.json"),
	)
	return
}
