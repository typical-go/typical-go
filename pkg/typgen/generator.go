package typgen

import (
	"errors"

	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	Generator struct {
		Walker     Walker
		Annotators []Annotator
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
	annotations, err := Compile(filePaths...)
	if err != nil {
		return err
	}

	initFile := NewInitFile()
	ctx := &Context{
		Context:     c,
		InitFile:    initFile,
		Annotations: annotations,
	}

	for _, a := range g.Annotators {

		annots := Filter(annotations, a)

		if len(annots) > 0 {
			ctx.PutInit("") // NOTE: intentionally put blank
			ctx.PutInitSprintf("// <<< [Annotator:%s] ", a.AnnotationName())
			for _, annot := range Filter(annotations, a) {
				if err := a.ProcessAnnot(ctx, annot); err != nil {
					return err
				}
			}
			ctx.PutInitSprintf("// [Annotator:%s] >>>", a.AnnotationName())
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
