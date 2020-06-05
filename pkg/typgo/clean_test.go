package typgo_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestCleaner(t *testing.T) {
	testcases := []struct {
		testName string
		typgo.Cleaner
		context     *typgo.Context
		expectedErr string
	}{
		{
			Cleaner:     typgo.NewClean(func(*typgo.Context) error { return errors.New("some-error") }),
			expectedErr: "some-error",
		},
		{
			Cleaner: typgo.NewClean(func(*typgo.Context) error { return nil }),
		},
		{
			Cleaner: typgo.Cleans{
				typgo.NewClean(func(*typgo.Context) error { return errors.New("some-error") }),
				typgo.NewClean(func(*typgo.Context) error { return nil }),
			},
			expectedErr: "some-error",
		},
		{
			Cleaner: typgo.Cleans{
				typgo.NewClean(func(*typgo.Context) error { return nil }),
				typgo.NewClean(func(*typgo.Context) error { return errors.New("some-error") }),
			},
			expectedErr: "some-error",
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			err := tt.Clean(tt.context)
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
