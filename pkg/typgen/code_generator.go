package typgen

import (
	"errors"
	"strings"

	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	CodeGenerator struct {
		Walker     Walker
		Annotators []Annotator
	}
	Annotator interface {
		AnnotationName() string
		IsAllowed(*Annotation) bool
	}
	AnnotProcessor interface {
		ProcessAnnot(*Context, *Annotation) error
	}
	AnnotatedFileProcessor interface {
		ProcessAnnotatedFile(*Context, *File, []*Annotation) error
	}
	PreAnnotator interface {
		BeforeAnnotate(*Context, []*Annotation) error
	}
	PostAnnotator interface {
		AfterAnnotate(*Context, []*Annotation) error
	}
)

const (
	InitFilePath = "internal/generated/init.go"
)

var (
	_ typgo.Tasker = (*CodeGenerator)(nil)
	_ typgo.Action = (*CodeGenerator)(nil)
)

// Task to annotate
func (g *CodeGenerator) Task() *typgo.Task {
	return &typgo.Task{
		Name:    "generate",
		Aliases: []string{"g"},
		Usage:   "Generate code based on annotation directive ('@')",
		Action:  g,
	}
}

// Execute annotation
func (g *CodeGenerator) Execute(c *typgo.Context) error {
	filePaths := g.walk()
	if len(filePaths) < 1 {
		return errors.New("walker couldn't find any filepath")
	}
	annotations, err := Compile(filePaths...)
	if err != nil {
		return err
	}

	ctx := NewContext(c, annotations)

	for _, ator := range g.Annotators {
		filtered := Filter(ctx.Annotations, ator)
		if len(filtered) > 0 {
			err := ExecuteAnnotator(ctx, ator, filtered)
			if err != nil {
				return err
			}
		}
	}
	return ctx.WriteInitFile(c, InitFilePath)
}

func (a *CodeGenerator) walk() []string {
	if a.Walker == nil {
		return Layouts{"internal"}.Walk()
	}
	return a.Walker.Walk()
}

func Filter(annotations []*Annotation, annotator Annotator) []*Annotation {
	filtered := []*Annotation{}
	name := annotator.AnnotationName()
	for _, annot := range annotations {
		if strings.EqualFold(name, annot.Name) && annotator.IsAllowed(annot) {
			filtered = append(filtered, annot)
		}
	}
	return filtered
}

func ExecuteAnnotator(ctx *Context, ator Annotator, filtered []*Annotation) error {
	if pre, ok := ator.(PreAnnotator); ok {
		if err := pre.BeforeAnnotate(ctx, filtered); err != nil {
			return err
		}
	}

	if proc, ok := ator.(AnnotProcessor); ok {
		ctx.AppendInit("") // NOTE: intentionally put blank
		ctx.AppendInitf("// <<< [Annotator:%s] ", ator.AnnotationName())
		for _, annot := range filtered {
			if err := proc.ProcessAnnot(ctx, annot); err != nil {
				return err
			}
		}
		ctx.AppendInitf("// [Annotator:%s] >>>", ator.AnnotationName())
	}

	if proc, ok := ator.(AnnotatedFileProcessor); ok {
		m := MappedAnnotsByFile(filtered)
		for file, annots := range m {
			if err := proc.ProcessAnnotatedFile(ctx, file, annots); err != nil {
				return err
			}
		}
	}

	if post, ok := ator.(PostAnnotator); ok {
		if err := post.AfterAnnotate(ctx, filtered); err != nil {
			return err
		}
	}

	return nil
}

func MappedAnnotsByFile(annots []*Annotation) map[*File][]*Annotation {
	m := make(map[*File][]*Annotation)
	for _, annot := range annots {
		file := annot.Decl.File
		m[file] = append(m[file], annot)
	}
	return m
}
