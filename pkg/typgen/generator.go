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

var _ typgo.Tasker = (*Generator)(nil)
var _ typgo.Action = (*Generator)(nil)

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
	for _, annot := range g.Annotations {
		if err := ExecuteAnnotation(c, annot, dirs); err != nil {
			return err
		}
	}
	return nil
}

func (a *Generator) walk() []string {
	if a.Walker == nil {
		return Layouts{"internal"}.Walk()
	}
	return a.Walker.Walk()
}
