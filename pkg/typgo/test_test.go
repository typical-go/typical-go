package typgo_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestTests(t *testing.T) {
	testcases := []struct {
		testName string
		typgo.Tester
		context     *typgo.Context
		expectedErr string
	}{
		{
			Tester:      typgo.NewTest(func(*typgo.Context) error { return errors.New("some-error") }),
			expectedErr: "some-error",
		},
		{
			Tester: typgo.NewTest(func(*typgo.Context) error { return nil }),
		},
		{
			Tester: typgo.Tests{
				typgo.NewTest(func(*typgo.Context) error { return errors.New("some-error") }),
				typgo.NewTest(func(*typgo.Context) error { return nil }),
			},
			expectedErr: "some-error",
		},
		{
			Tester: typgo.Tests{
				typgo.NewTest(func(*typgo.Context) error { return nil }),
				typgo.NewTest(func(*typgo.Context) error { return errors.New("some-error") }),
			},
			expectedErr: "some-error",
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			err := tt.Test(tt.context)
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
