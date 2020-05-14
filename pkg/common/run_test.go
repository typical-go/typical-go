package common_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/common"
)

func TestGracefulRun(t *testing.T) {
	testcases := []struct {
		testName string
		startFn  common.Fn
		stopFn   common.Fn
		expected common.Errors
	}{
		{
			startFn: func() error { return nil },
			stopFn:  func() error { return nil },
		},
		{
			startFn: func() error { return nil },
			stopFn:  func() error { return errors.New("some-stop-error") },
			expected: common.Errors{
				errors.New("some-stop-error"),
			},
		},
		{
			startFn: func() error { return errors.New("some-start-error") },
			stopFn:  func() error { return nil },
			expected: common.Errors{
				errors.New("some-start-error"),
			},
		},
		{
			startFn: func() error { return errors.New("some-start-error") },
			stopFn:  func() error { return errors.New("some-stop-error") },
			expected: common.Errors{
				errors.New("some-start-error"),
				errors.New("some-stop-error"),
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			require.EqualValues(t, tt.expected, common.GracefulRun(tt.startFn, tt.stopFn))
		})
	}
}
