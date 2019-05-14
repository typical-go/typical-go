package typigo

import (
	"github.com/typical-go/typical-code-generator/utility/bashkit"
	"github.com/typical-go/typical-code-generator/utility/linux"
)

func NewProject(name, archetype, path string) (err error) {
	projectPath := bashkit.GOPATH() + "/src/" + path
	typicalPath := projectPath + "/.typical"

	screen := bashkit.NewScreen()
	err = screen.Run(
		linux.MakeDirectory(typicalPath),
	)
	return
}
