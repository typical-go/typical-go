package typgo_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestRunner(t *testing.T) {
	testcases := []struct {
		testName string
		typgo.Runner
		context     *typgo.Context
		expectedErr string
	}{
		{
			Runner:      typgo.NewRunner(func(*typgo.Context) error { return errors.New("some-error") }),
			expectedErr: "some-error",
		},
		{
			Runner: typgo.NewRunner(func(*typgo.Context) error { return nil }),
		},
		{
			Runner: typgo.Runners{
				typgo.NewRunner(func(*typgo.Context) error { return nil }),
				typgo.NewRunner(func(*typgo.Context) error { return errors.New("some-error") }),
			},
			expectedErr: "some-error",
		},
		{
			Runner: typgo.Runners{
				typgo.NewRunner(func(*typgo.Context) error { return errors.New("some-error") }),
				typgo.NewRunner(func(*typgo.Context) error { return nil }),
			},
			expectedErr: "some-error",
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			err := tt.Run(tt.context)
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
