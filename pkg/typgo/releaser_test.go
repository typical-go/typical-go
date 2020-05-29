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
		typgo.Releaser
		context     *typgo.Context
		expectedErr string
	}{
		{
			Releaser:    typgo.NewReleaser(func(*typgo.Context) error { return errors.New("some-error") }),
			expectedErr: "some-error",
		},
		{
			Releaser: typgo.NewReleaser(func(*typgo.Context) error { return nil }),
		},
		{
			Releaser: typgo.Releasers{
				typgo.NewReleaser(func(*typgo.Context) error { return nil }),
				typgo.NewReleaser(func(*typgo.Context) error { return errors.New("some-error") }),
			},
			expectedErr: "some-error",
		},
		{
			Releaser: typgo.Releasers{
				typgo.NewReleaser(func(*typgo.Context) error { return errors.New("some-error") }),
				typgo.NewReleaser(func(*typgo.Context) error { return nil }),
			},
			expectedErr: "some-error",
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			err := tt.Release(tt.context)
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
