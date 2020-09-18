package typrls_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typrls"
)

func TestReleasers(t *testing.T) {
	testcases := []struct {
		testName string
		typrls.Releaser
		context     *typrls.Context
		expectedErr string
	}{
		{
			Releaser:    typrls.NewReleaser(func(*typrls.Context) error { return errors.New("some-error") }),
			expectedErr: "some-error",
		},
		{
			Releaser: typrls.NewReleaser(func(*typrls.Context) error { return nil }),
		},
		{
			Releaser: typrls.Releasers{
				typrls.NewReleaser(func(*typrls.Context) error { return nil }),
				typrls.NewReleaser(func(*typrls.Context) error { return errors.New("some-error") }),
			},
			expectedErr: "some-error",
		},
		{
			Releaser: typrls.Releasers{
				typrls.NewReleaser(func(*typrls.Context) error { return errors.New("some-error") }),
				typrls.NewReleaser(func(*typrls.Context) error { return nil }),
			},
			expectedErr: "some-error",
		},
		{
			Releaser: typrls.Releasers{
				typrls.NewReleaser(func(*typrls.Context) error { return nil }),
				typrls.NewReleaser(func(*typrls.Context) error { return nil }),
			},
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
