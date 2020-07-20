package typgo_test

import (
	"errors"
	"flag"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestRunCompile(t *testing.T) {
	testcases := []typgo.CmdTestCase{
		{
			Cmd: &typgo.RunCmd{
				Action: typgo.NewAction(func(*typgo.Context) error {
					return errors.New("some-error")
				}),
			},
			Expected: typgo.Command{
				Name:            "run",
				Aliases:         []string{"r"},
				Usage:           "Run the project in local environment",
				SkipFlagParsing: true,
			},
			ExpectedError: "some-error",
		},
	}
	for _, tt := range testcases {
		tt.Run(t)
	}
}

func TestStdRun(t *testing.T) {
	stdRun := &typgo.StdRun{}
	c := &typgo.Context{
		Context: cli.NewContext(nil, &flag.FlagSet{}, nil),
		BuildSys: &typgo.BuildSys{
			Descriptor: &typgo.Descriptor{Name: "some-name"},
		},
	}

	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: []string{"bin/some-name"}},
	})
	defer unpatch(t)

	require.NoError(t, stdRun.Execute(c))
}

func TestStdRun_Predefined(t *testing.T) {
	stdRun := &typgo.StdRun{
		Binary: "some-binary",
	}
	c := &typgo.Context{
		Context: cli.NewContext(nil, &flag.FlagSet{}, nil),
		BuildSys: &typgo.BuildSys{
			Descriptor: &typgo.Descriptor{Name: "some-name"},
		},
	}

	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: []string{"some-binary"}},
	})
	defer unpatch(t)

	require.NoError(t, stdRun.Execute(c))
}
