package typgo_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestPrebuild(t *testing.T) {
	testcases := []struct {
		testName string
		typgo.Prebuilder
		context     *typgo.PrebuildContext
		expectedErr string
	}{
		{
			Prebuilder:  typgo.NewPrebuild(func(*typgo.PrebuildContext) error { return errors.New("some-error") }),
			expectedErr: "some-error",
		},
		{
			Prebuilder: typgo.NewPrebuild(func(*typgo.PrebuildContext) error { return nil }),
		},
		{
			Prebuilder: typgo.Prebuilds{
				typgo.NewPrebuild(func(*typgo.PrebuildContext) error { return nil }),
				typgo.NewPrebuild(func(*typgo.PrebuildContext) error { return errors.New("some-error") }),
			},
			expectedErr: "some-error",
		},
		{
			Prebuilder: typgo.Prebuilds{
				typgo.NewPrebuild(func(*typgo.PrebuildContext) error { return errors.New("some-error") }),
				typgo.NewPrebuild(func(*typgo.PrebuildContext) error { return nil }),
			},
			expectedErr: "some-error",
		},
	}
	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			err := tt.Prebuild(tt.context)
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
