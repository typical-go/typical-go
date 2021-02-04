package typast_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestAnnotateCmd(t *testing.T) {
	annonateCmd := &typast.AnnotateMe{
		Sources: []string{"internal"},
	}
	d := &typgo.Descriptor{}

	command := annonateCmd.Task(d)
	require.Equal(t, "annotate", command.Name)
	require.Equal(t, []string{"a"}, command.Aliases)
	require.Equal(t, "Annotate the project and generate code", command.Usage)
	require.NoError(t, command.Action(&cli.Context{}))

	_, err := annonateCmd.CreateContext(&typgo.Context{Descriptor: d})
	require.NoError(t, err)
}

func TestAnnotateCmd_Defined(t *testing.T) {
	annonateCmd := typast.AnnotateMe{
		Sources: []string{"internal"},
		Annotators: []typast.Annotator{
			typast.NewAnnotator(func(*typast.Context) error {
				return errors.New("some-error")
			}),
		},
	}
	d := &typgo.Descriptor{}

	command := annonateCmd.Task(d)
	require.EqualError(t, command.Action(&cli.Context{}), "some-error")

	_, err := annonateCmd.CreateContext(&typgo.Context{Descriptor: d})
	require.NoError(t, err)
}

func TestAnnotateCmd_MissingSources(t *testing.T) {
	annonateCmd := typast.AnnotateMe{}

	command := annonateCmd.Task(&typgo.Descriptor{})
	require.EqualError(t, command.Action(&cli.Context{}), "'Sources' is missing")
}

func TestAnnotators_Execute(t *testing.T) {
	testcases := []struct {
		TestName string
		typast.AnnotateMe
		Context     *typgo.Context
		ExpectedErr string
	}{
		{
			Context: &typgo.Context{
				Descriptor: &typgo.Descriptor{},
			},
			AnnotateMe: typast.AnnotateMe{
				Sources: []string{"internal"},
				Annotators: []typast.Annotator{
					typast.NewAnnotator(func(c *typast.Context) error { return errors.New("some-error-1") }),
					typast.NewAnnotator(func(c *typast.Context) error { return errors.New("some-error-2") }),
				},
			},
			ExpectedErr: "some-error-1",
		},
		{
			Context: &typgo.Context{
				Descriptor: &typgo.Descriptor{},
			},
			AnnotateMe: typast.AnnotateMe{
				Sources: []string{"internal"},
				Annotators: []typast.Annotator{
					typast.NewAnnotator(func(c *typast.Context) error { return nil }),
					typast.NewAnnotator(func(c *typast.Context) error { return errors.New("some-error-2") }),
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
