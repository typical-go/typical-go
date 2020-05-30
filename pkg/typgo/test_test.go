package typgo_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestTests(t *testing.T) {
	testcases := []struct {
		testName    string
		Test        typgo.Test
		context     *typgo.Context
		expectedErr string
	}{
		{
			Test:        typgo.NewTest(func(*typgo.Context) error { return errors.New("some-error") }),
			expectedErr: "some-error",
		},
		{
			Test: typgo.NewTest(func(*typgo.Context) error { return nil }),
		},
		{
			Test: typgo.Tests{
				typgo.NewTest(func(*typgo.Context) error { return errors.New("some-error") }),
				typgo.NewTest(func(*typgo.Context) error { return nil }),
			},
			expectedErr: "some-error",
		},
		{
			Test: typgo.Tests{
				typgo.NewTest(func(*typgo.Context) error { return nil }),
				typgo.NewTest(func(*typgo.Context) error { return errors.New("some-error") }),
			},
			expectedErr: "some-error",
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			err := tt.Test.Test(tt.context)
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
