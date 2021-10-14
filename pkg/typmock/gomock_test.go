package typmock_test

import (
	"errors"
	"flag"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgen"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typmock"
	"github.com/urfave/cli/v2"
)

func TestAnnotate_MockgenError(t *testing.T) {
	typgo.TypicalTmp = ".typical-tmp"
	defer func() { typgo.TypicalTmp = "" }()

	var out strings.Builder
	c := &typgo.Context{
		Context:    cli.NewContext(nil, &flag.FlagSet{}, nil),
		Descriptor: &typgo.Descriptor{},
		Logger:     typgo.Logger{Stdout: &out},
	}
	defer c.PatchBash([]*typgo.MockCommand{
		{
			CommandLine: "go build -o .typical-tmp/bin/mockgen github.com/golang/mock/mockgen",
		},
		{
			CommandLine: ".typical-tmp/bin/mockgen -destination internal/generated/parent/mypkg_mock/some_interface.go -package mypkg_mock /parent/mypkg SomeInterface",
			ReturnError: errors.New("some-error"),
		},
	})(t)

	gomock := &typmock.GoMock{}
	ctx := &typgen.Context{Context: c}
	annot := &typgen.Annotation{
		Name: "@mock",
		Decl: &typgen.Decl{
			File: &typgen.File{Name: "mypkg", Path: "parent/mypkg/some_interface.go"},
			Type: &typgen.Interface{TypeDecl: typgen.TypeDecl{Name: "SomeInterface"}},
		},
	}
	require.NoError(t, gomock.ProcessAnnot(ctx, annot))
	require.Equal(t, `> go build -o .typical-tmp/bin/mockgen github.com/golang/mock/mockgen
> .typical-tmp/bin/mockgen -destination internal/generated/parent/mypkg_mock/some_interface.go -package mypkg_mock /parent/mypkg SomeInterface
> Fail to mock '/parent/mypkg': some-error
`, out.String())
}

func TestGoMock(t *testing.T) {
	gomock := &typmock.GoMock{}
	require.Equal(t, "@mock", gomock.AnnotationName())
	require.True(t, gomock.IsAllowed(&typgen.Annotation{
		Decl: &typgen.Decl{
			Type: &typgen.Interface{
				TypeDecl: typgen.TypeDecl{
					Name: "SomeInterface",
				},
			},
		},
	}))
	require.Equal(t, &typgo.Task{
		Name:   "mock",
		Usage:  "Generate mock class",
		Action: gomock,
	}, gomock.Task())

	require.EqualError(t, gomock.Execute(nil), "walker couldn't find any filepath")
}
