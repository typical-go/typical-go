package typannot_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestAnnotateCmd(t *testing.T) {
	annonateCmd := &typannot.AnnotateCmd{}

	sys := &typgo.BuildSys{
		Descriptor: &typgo.Descriptor{},
	}
	command := annonateCmd.Command(sys)
	require.Equal(t, "annotate", command.Name)
	require.Equal(t, []string{"a"}, command.Aliases)
	require.Equal(t, "Annotate the project and generate code", command.Usage)
	require.NoError(t, command.Action(&cli.Context{}))

	ctx, err := annonateCmd.CreateContext(&typgo.Context{BuildSys: sys})
	require.NoError(t, err)
	require.Equal(t, "internal/generated", ctx.Destination)
}

func TestAnnotateCmd_Defined(t *testing.T) {
	annonateCmd := &typannot.AnnotateCmd{
		Destination: "some-destination",
		Annotators: []typannot.Annotator{
			typannot.NewAnnotator(func(*typannot.Context) error {
				return errors.New("some-error")
			}),
		},
	}
	sys := &typgo.BuildSys{
		Descriptor: &typgo.Descriptor{},
	}
	command := annonateCmd.Command(sys)
	require.EqualError(t, command.Action(&cli.Context{}), "some-error")

	ctx, err := annonateCmd.CreateContext(&typgo.Context{BuildSys: sys})
	require.NoError(t, err)
	require.Equal(t, "some-destination", ctx.Destination)
}

func TestAnnotators_Execute(t *testing.T) {
	testcases := []struct {
		TestName string
		typannot.AnnotateCmd
		Context     *typgo.Context
		ExpectedErr string
	}{
		{
			Context: &typgo.Context{BuildSys: &typgo.BuildSys{
				Descriptor: &typgo.Descriptor{},
			}},
			AnnotateCmd: typannot.AnnotateCmd{
				Annotators: []typannot.Annotator{
					typannot.NewAnnotator(func(c *typannot.Context) error { return errors.New("some-error-1") }),
					typannot.NewAnnotator(func(c *typannot.Context) error { return errors.New("some-error-2") }),
				},
			},
			ExpectedErr: "some-error-1",
		},
		{
			Context: &typgo.Context{BuildSys: &typgo.BuildSys{
				Descriptor: &typgo.Descriptor{},
			}},
			AnnotateCmd: typannot.AnnotateCmd{
				Annotators: []typannot.Annotator{
					typannot.NewAnnotator(func(c *typannot.Context) error { return nil }),
					typannot.NewAnnotator(func(c *typannot.Context) error { return errors.New("some-error-2") }),
				},
			},
			ExpectedErr: "some-error-2",
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			err := tt.Execute(tt.Context)
			if tt.ExpectedErr != "" {
				require.EqualError(t, err, tt.ExpectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
