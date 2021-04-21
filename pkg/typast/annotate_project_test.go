package typast_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var (
	someStructDecl = &typast.Decl{
		File: typast.File{
			Path:    "sample_test.go",
			Package: "typast_test",
		},
		Type: &typast.StructDecl{
			TypeDecl: typast.TypeDecl{
				GenDecl: typast.GenDecl{
					Docs: []string{
						"// sampleStruct",
						"// @tag1",
						"// @tag2 (key1:\"\", key2: \"\", key3:\"value3\")",
					},
				},
				Name: "sampleStruct",
			},
			Fields: []*typast.Field{
				{
					Names:     []string{"sampleInt"},
					Type:      "int",
					StructTag: reflect.StructTag(`default:"value1"`),
				},
				{
					Names:     []string{"sampleString"},
					Type:      "string",
					StructTag: reflect.StructTag(`default:"value2"`),
				},
			},
		},
	}

	someFunctionDecl2 = &typast.Decl{
		File: typast.File{
			Path:    "sample_test.go",
			Package: "typast_test",
		},
		Type: &typast.FuncDecl{
			Name:   "sampleFunction2",
			Params: &typast.FieldList{},
			Docs: []string{
				"// GetWriter to get writer to greet the world",
				"// @ctor",
			},
		},
	}

	someInterface2Decl = &typast.Decl{
		File: typast.File{
			Path:    "sample_test.go",
			Package: "typast_test",
		},
		Type: &typast.InterfaceDecl{
			TypeDecl: typast.TypeDecl{
				Name: "sampleInterface2",
				Docs: []string{"// @tag3"},
			},
		},
	}

	someStruct2Decl = &typast.Decl{
		File: typast.File{
			Path:    "sample_test.go",
			Package: "typast_test",
		},
		Type: &typast.StructDecl{
			TypeDecl: typast.TypeDecl{
				Name: "sampleStruct2",
				Docs: []string{
					"// sampleStruct2 asdf",
					"// @tag4",
				},
			},
		},
	}
)

func TestAnnotateProject(t *testing.T) {
	annotateProject := &typast.AnnotateProject{}
	require.Equal(t, &typgo.Task{
		Name:    "annotate",
		Aliases: []string{"a"},
		Usage:   "Annotate the project and generate code",
		Action:  annotateProject,
	}, annotateProject.Task())
}

func TestAnnotateProject_Execute(t *testing.T) {
	var directives typast.Directives
	action := &typast.AnnotateProject{
		Walker: typast.FilePaths{"sample_test.go"},
		Annotators: []typast.Annotator{
			&typast.Annotation{
				Processor: typast.NewProcessor(func(c *typgo.Context, d typast.Directives) error {
					directives = d
					return nil
				}),
			},
		},
	}
	require.NoError(t, action.Execute(&typgo.Context{}))
	require.EqualValues(t, typast.Directives{
		{Decl: someStructDecl, TagName: "@tag1"},
		{Decl: someStructDecl, TagName: "@tag2", TagParam: `key1:"", key2: "", key3:"value3"`},
		{Decl: someFunctionDecl2, TagName: "@ctor"},
		{Decl: someInterface2Decl, TagName: "@tag3"},
		{Decl: someStruct2Decl, TagName: "@tag4"},
	}, directives)
}

func TestAnnotateProject_Error(t *testing.T) {
	testcases := []struct {
		TestName        string
		AnnotateProject *typast.AnnotateProject
		ExpectedErr     string
	}{
		{
			AnnotateProject: &typast.AnnotateProject{
				Walker: typast.FilePaths{"bad_file.go"},
			},
			ExpectedErr: "open bad_file.go: no such file or directory",
		},
		{
			AnnotateProject: &typast.AnnotateProject{
				Walker: typast.FilePaths{"sample_test.go"},
				Annotators: []typast.Annotator{
					&typast.Annotation{
						Processor: typast.NewProcessor(func(c *typgo.Context, d typast.Directives) error {
							return errors.New("some-error")
						}),
					},
				},
			},
			ExpectedErr: "some-error",
		},
		{
			AnnotateProject: &typast.AnnotateProject{},
			ExpectedErr:     "walker couldn't find any filepath",
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			err := tt.AnnotateProject.Execute(&typgo.Context{})
			if tt.ExpectedErr != "" {
				require.EqualError(t, err, tt.ExpectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestParseRawAnnot(t *testing.T) {
	testcases := []struct {
		TestName         string
		Raw              string
		ExpectedTagName  string
		ExpectedTagAttrs string
	}{
		{
			TestName:        "tag only",
			Raw:             `@tag1`,
			ExpectedTagName: "@tag1",
		},
		{
			TestName:        "tag only with space",
			Raw:             `@tag2 extra1`,
			ExpectedTagName: "@tag2",
		},
		{
			TestName:         "with attribute",
			Raw:              `@tag3("name":"wire1")`,
			ExpectedTagName:  "@tag3",
			ExpectedTagAttrs: `"name":"wire1"`,
		},
		{
			TestName:         "there is space between tagname and attribute",
			Raw:              `@tag4 ("name":"wire1")`,
			ExpectedTagName:  "@tag4",
			ExpectedTagAttrs: `"name":"wire1"`,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			tagName, tagAttrs := typast.ParseRawAnnot(tt.Raw)
			require.Equal(t, tt.ExpectedTagName, tagName)
			require.Equal(t, tt.ExpectedTagAttrs, tagAttrs)
		})
	}
}
