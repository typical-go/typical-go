package typmock_test

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typmock"
	"github.com/urfave/cli/v2"
)

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
					Name:    "SomeInterface",
					Package: "mypkg",
					Path:    "parent/path/some_interface.go",
					Type:    &typannot.InterfaceType{},
				}},
		},
	}

	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: []string{"go", "build", "-o", "/bin/mockgen", "github.com/golang/mock/mockgen"}},
		{CommandLine: []string{"rm", "-rf", "parent/path_mock"}},
		{CommandLine: []string{"/bin/mockgen", "-destination", "parentmypkg_mock/some_interface.go", "-package", "mypkg_mock", "/parent/path", "SomeInterface"}},
	})
	defer unpatch(t)

	require.NoError(t, typmock.Annotate(c, summary))
}
