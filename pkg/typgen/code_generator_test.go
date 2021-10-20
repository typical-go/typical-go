package typgen_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgen"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestGenerator(t *testing.T) {
	gen := &typgen.CodeGenerator{}
	require.Equal(t, &typgo.Task{
		Name:    "generate",
		Aliases: []string{"g"},
		Usage:   "Generate code based on annotation directive ('@')",
		Action:  gen,
	}, gen.Task())
}

func TestFilter(t *testing.T) {
	annot1 := &typgen.Annotation{Name: "@qwerty"}
	annot2 := &typgen.Annotation{Name: "@qwerty"}
	annot3 := &typgen.Annotation{Name: "@asdf"}
	annot4 := &typgen.Annotation{Name: "@asdf"}
	testcases := []struct {
		TestName string
		Annots   []*typgen.Annotation
		Ator     typgen.Annotator
		Expected []*typgen.Annotation
	}{
		{
			Annots: []*typgen.Annotation{annot1, annot2, annot3, annot4},
			Ator: &annotator{
				AnnotationNameFn: func() string {
					return "@qwerty"
				},
				IsAllowedFn: func(a *typgen.Annotation) bool {
					return true
				},
			},
			Expected: []*typgen.Annotation{annot1, annot2},
		},
		{
			Annots: []*typgen.Annotation{annot1, annot2, annot3, annot4},
			Ator: &annotator{
				AnnotationNameFn: func() string {
					return "@qwerty"
				},
				IsAllowedFn: func(a *typgen.Annotation) bool {
					return false
				},
			},
			Expected: []*typgen.Annotation{},
		},
		{
			Annots: []*typgen.Annotation{annot1, annot2, annot3, annot4},
			Ator: &annotator{
				AnnotationNameFn: func() string {
					return "@zxcv"
				},
				IsAllowedFn: func(a *typgen.Annotation) bool {
					return true
				},
			},
			Expected: []*typgen.Annotation{},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, typgen.Filter(tt.Annots, tt.Ator))
		})
	}
}

func TestExecuteProcessor_WhenBeforeAnnotateError(t *testing.T) {
	ctx := &typgen.Context{}
	ator := &annotator{
		BeforeAnnotateFn: func(c *typgen.Context, a []*typgen.Annotation) error {
			return errors.New("before-error")
		},
	}
	filtered := []*typgen.Annotation{}

	err := typgen.ExecuteAnnotator(ctx, ator, filtered)
	require.EqualError(t, err, "before-error")
}

func TestExecuteProcessor_WhenProcessAnnotError(t *testing.T) {
	ctx := &typgen.Context{
		InitFile: typgen.NewInitFile(),
	}
	ator := &annotator{
		ProcessAnnotFn: func(c *typgen.Context, a *typgen.Annotation) error {
			return errors.New("process-error")
		},
	}
	filtered := []*typgen.Annotation{{}}

	err := typgen.ExecuteAnnotator(ctx, ator, filtered)
	require.EqualError(t, err, "process-error")
}

func TestExecuteProcessor_WhenProcessAnnotedFileError(t *testing.T) {
	ctx := &typgen.Context{
		InitFile: typgen.NewInitFile(),
	}
	ator := &annotator{
		ProcessAnnotatedFileFn: func(c *typgen.Context, f *typgen.File, a []*typgen.Annotation) error {
			return errors.New("process-error")
		},
	}
	filtered := []*typgen.Annotation{
		{
			Decl: &typgen.Decl{
				File: &typgen.File{},
			},
		},
	}

	err := typgen.ExecuteAnnotator(ctx, ator, filtered)
	require.EqualError(t, err, "process-error")
}

func TestExecuteProcessor_WhenAfterAnnotateError(t *testing.T) {
	ctx := &typgen.Context{
		InitFile: typgen.NewInitFile(),
	}
	ator := &annotator{
		AfterAnnotateFn: func(c *typgen.Context, a []*typgen.Annotation) error {
			return errors.New("after-error")
		},
	}
	filtered := []*typgen.Annotation{}

	err := typgen.ExecuteAnnotator(ctx, ator, filtered)
	require.EqualError(t, err, "after-error")
}

func TestExecuteProcessor(t *testing.T) {
	ctx := &typgen.Context{
		InitFile: typgen.NewInitFile(),
	}
	ator := &annotator{}
	filtered := []*typgen.Annotation{}

	err := typgen.ExecuteAnnotator(ctx, ator, filtered)
	require.NoError(t, err)
}

func TestMappedAnnotsByFile(t *testing.T) {
	file1 := &typgen.File{}
	file2 := &typgen.File{}
	annot1 := &typgen.Annotation{
		Decl: &typgen.Decl{
			File: file1,
		},
	}
	annot2 := &typgen.Annotation{
		Decl: &typgen.Decl{
			File: file1,
		},
	}
	annot3 := &typgen.Annotation{
		Decl: &typgen.Decl{
			File: file2,
		},
	}
	annots := []*typgen.Annotation{annot1, annot2, annot3}
	expected := map[*typgen.File][]*typgen.Annotation{
		file1: {annot1, annot2},
		file2: {annot3},
	}
	require.EqualValues(t, expected, typgen.MappedAnnotsByFile(annots))

}

type annotator struct {
	AnnotationNameFn       func() string
	IsAllowedFn            func(*typgen.Annotation) bool
	ProcessAnnotFn         func(*typgen.Context, *typgen.Annotation) error
	ProcessAnnotatedFileFn func(*typgen.Context, *typgen.File, []*typgen.Annotation) error
	BeforeAnnotateFn       func(*typgen.Context, []*typgen.Annotation) error
	AfterAnnotateFn        func(*typgen.Context, []*typgen.Annotation) error
}

var (
	_ typgen.Annotator              = (*annotator)(nil)
	_ typgen.AnnotProcessor         = (*annotator)(nil)
	_ typgen.AnnotatedFileProcessor = (*annotator)(nil)
	_ typgen.PreAnnotator           = (*annotator)(nil)
	_ typgen.PostAnnotator          = (*annotator)(nil)
)

func (a *annotator) AnnotationName() string {
	if a.AnnotationNameFn == nil {
		return ""
	}
	return a.AnnotationNameFn()
}

func (a *annotator) IsAllowed(annot *typgen.Annotation) bool {
	if a.IsAllowedFn == nil {
		return true
	}
	return a.IsAllowedFn(annot)
}

func (a *annotator) ProcessAnnot(c *typgen.Context, filtered *typgen.Annotation) error {
	if a.ProcessAnnotFn == nil {
		return nil
	}
	return a.ProcessAnnotFn(c, filtered)
}

func (a *annotator) ProcessAnnotatedFile(c *typgen.Context, file *typgen.File, annots []*typgen.Annotation) error {
	if a.ProcessAnnotatedFileFn == nil {
		return nil
	}
	return a.ProcessAnnotatedFileFn(c, file, annots)
}

func (a *annotator) BeforeAnnotate(c *typgen.Context, annots []*typgen.Annotation) error {
	if a.BeforeAnnotateFn == nil {
		return nil
	}
	return a.BeforeAnnotateFn(c, annots)
}

func (a *annotator) AfterAnnotate(c *typgen.Context, filtered []*typgen.Annotation) error {
	if a.AfterAnnotateFn == nil {
		return nil
	}
	return a.AfterAnnotateFn(c, filtered)
}
