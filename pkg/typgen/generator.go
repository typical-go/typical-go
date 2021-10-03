package typgen

import (
	"errors"

	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	Generator struct {
		Walker      Walker
		Annotations []Annotation
	}
)

//
// Generator
//

var (
	_ typgo.Tasker = (*Generator)(nil)
	_ typgo.Action = (*Generator)(nil)
)

// Task to annotate
func (g *Generator) Task() *typgo.Task {
	return &typgo.Task{
		Name:    "generate",
		Aliases: []string{"g"},
		Usage:   "Generate code based on annotation directive ('@')",
		Action:  g,
	}
}

// Execute annotation
func (g *Generator) Execute(c *typgo.Context) error {
	filePaths := g.walk()
	if len(filePaths) < 1 {
		return errors.New("walker couldn't find any filepath")
	}
	dirs, err := Compile(filePaths...)
	if err != nil {
		return err
	}

	initFile := NewInitFile()
	for _, annot := range g.Annotations {
		ctx := &Context{
			Context:  c,
			InitFile: initFile,
			Annot:    annot,
			Dirs:     Filter(dirs, annot),
		}

		if err := annot.Process(ctx); err != nil {
			return err
		}
	}
	return initFile.WriteTo(c, "internal/generated/init.go")
}

func (a *Generator) walk() []string {
	if a.Walker == nil {
		return Layouts{"internal"}.Walk()
	}
	return a.Walker.Walk()
}
