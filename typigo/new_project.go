package typigo

import (
	"github.com/typical-go/typical-code-generator/command"
	"github.com/typical-go/typical-code-generator/executor"
	"github.com/typical-go/typical-code-generator/metadata"
	"github.com/typical-go/typical-code-generator/utility/oskit"
)

// NewProject new project
func NewProject(path string) (err error) {
	projectPath := oskit.GOPATH() + "/src/" + path
	typicalPath := projectPath + "/.typical"

	context := metadata.Context{
		Name: "meh",
		Path: path,
	}

	err = executor.Run(
		command.MakeDirectory(typicalPath),
		command.SaveMetadataContext(typicalPath+"/_context.json", context),
	)
	return
}
