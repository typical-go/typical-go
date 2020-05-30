package typgo_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestReleasers(t *testing.T) {
	testcases := []struct {
		testName string
		typgo.Release
		context     *typgo.Context
		expectedErr string
	}{
		{
			Release:     typgo.NewRelease(func(*typgo.Context) error { return errors.New("some-error") }),
			expectedErr: "some-error",
		},
		{
			Release: typgo.NewRelease(func(*typgo.Context) error { return nil }),
		},
		{
			Release: typgo.Releases{
				typgo.NewRelease(func(*typgo.Context) error { return nil }),
				typgo.NewRelease(func(*typgo.Context) error { return errors.New("some-error") }),
			},
			expectedErr: "some-error",
		},
		{
			Release: typgo.Releases{
				typgo.NewRelease(func(*typgo.Context) error { return errors.New("some-error") }),
				typgo.NewRelease(func(*typgo.Context) error { return nil }),
			},
			expectedErr: "some-error",
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			err := tt.Release.Release(tt.context)
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
