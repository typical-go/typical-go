package typmock_test

import (
	"errors"
	"flag"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typmock"
	"github.com/urfave/cli/v2"
)

func TestCommand(t *testing.T) {
	mockCmd := &typmock.MockCmd{}
	command := mockCmd.Command(&typgo.BuildSys{})

	require.Equal(t, "mock", command.Name)
}

func TestAnnotate(t *testing.T) {
	c := &typgo.Context{
		BuildSys: &typgo.BuildSys{Descriptor: &typgo.Descriptor{ProjectLayouts: []string{"."}}},
		Context:  cli.NewContext(nil, &flag.FlagSet{}, nil),
	}
	summary := &typannot.Summary{
		Annots: []*typannot.Annot{
			{
				TagName: "@mock",
				Decl: &typannot.Decl{
					File: typannot.File{
						Package: "mypkg",
						Path:    "parent/path/some_interface.go",
					},
					DeclType: &typannot.InterfaceDecl{
						TypeDecl: typannot.TypeDecl{Name: "SomeInterface"},
					},
				}},
		},
	}

	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: "go build -o /bin/mockgen github.com/golang/mock/mockgen"},
		{CommandLine: "rm -rf parent/path_mock"},
		{CommandLine: "/bin/mockgen -destination parentmypkg_mock/some_interface.go -package mypkg_mock /parent/path SomeInterface"},
	})
	defer unpatch(t)

	require.NoError(t, typmock.Annotate(c, summary))
}

func TestAnnotate_InstallMockgenError(t *testing.T) {
	c := &typgo.Context{
		BuildSys: &typgo.BuildSys{Descriptor: &typgo.Descriptor{ProjectLayouts: []string{"."}}},
		Context:  cli.NewContext(nil, &flag.FlagSet{}, nil),
	}

	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{
			CommandLine: "go build -o /bin/mockgen github.com/golang/mock/mockgen",
			ReturnError: errors.New("some-error"),
		},
	})
	defer unpatch(t)

	require.EqualError(t, typmock.Annotate(c, &typannot.Summary{}), "some-error")
}

func TestAnnotate_MockgenError(t *testing.T) {
	var out strings.Builder
	typmock.Stdout = &out
	defer func() { typmock.Stdout = os.Stdout }()

	c := &typgo.Context{
		BuildSys: &typgo.BuildSys{Descriptor: &typgo.Descriptor{ProjectLayouts: []string{"."}}},
		Context:  cli.NewContext(nil, &flag.FlagSet{}, nil),
	}
	summary := &typannot.Summary{
		Annots: []*typannot.Annot{
			{
				TagName: "@mock",
				Decl: &typannot.Decl{
					File: typannot.File{
						Package: "mypkg",
						Path:    "parent/path/some_interface.go",
					},
					DeclType: &typannot.InterfaceDecl{
						TypeDecl: typannot.TypeDecl{Name: "SomeInterface"},
					},
				}},
		},
	}

	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: "go build -o /bin/mockgen github.com/golang/mock/mockgen"},
		{CommandLine: "rm -rf parent/path_mock"},
		{
			CommandLine: "/bin/mockgen -destination parentmypkg_mock/some_interface.go -package mypkg_mock /parent/path SomeInterface",
			ReturnError: errors.New("some-error"),
		},
	})
	defer unpatch(t)

	require.NoError(t, typmock.Annotate(c, summary))
	require.Equal(t, "Fail to mock '/parent/path.SomeInterface': some-error\n", out.String())
}
