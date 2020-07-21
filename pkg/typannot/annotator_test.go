package typannot_test

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestAnnotators_Execute(t *testing.T) {
	testcases := []struct {
		TestName string
		typannot.Annotators
		Context     *typgo.Context
		ExpectedErr string
	}{
		{
			Context: &typgo.Context{BuildSys: &typgo.BuildSys{
				Descriptor: &typgo.Descriptor{},
			}},
			Annotators: typannot.Annotators{
				typannot.NewAnnotator(func(c *typannot.Context) error { return errors.New("some-error-1") }),
				typannot.NewAnnotator(func(c *typannot.Context) error { return errors.New("some-error-2") }),
			},
			ExpectedErr: "some-error-1",
		},
		{
			Context: &typgo.Context{BuildSys: &typgo.BuildSys{
				Descriptor: &typgo.Descriptor{},
			}},
			Annotators: typannot.Annotators{
				typannot.NewAnnotator(func(c *typannot.Context) error { return nil }),
				typannot.NewAnnotator(func(c *typannot.Context) error { return errors.New("some-error-2") }),
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

func TestWalkLayout(t *testing.T) {
	os.MkdirAll("wrapper/some_pkg", os.ModePerm)
	os.MkdirAll("pkg/some_lib", os.ModePerm)
	os.Create("wrapper/some_pkg/some_file.go")
	os.Create("wrapper/some_pkg/not_go.xxx")
	os.Create("pkg/some_lib/lib.go")
	defer func() {
		os.RemoveAll("wrapper")
		os.RemoveAll("pkg")
	}()

	dirs, files := typannot.WalkLayout([]string{
		"pkg",
		"wrapper",
	})

	require.Equal(t, []string{
		"pkg",
		"pkg/some_lib",
		"wrapper",
		"wrapper/some_pkg",
	}, dirs)

	require.Equal(t, []string{
		"pkg/some_lib/lib.go",
		"wrapper/some_pkg/some_file.go",
	}, files)
}
