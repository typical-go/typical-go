package exor_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/exor"
)

func TestRun(t *testing.T) {
	testcases := []struct {
		executors     []exor.Executor
		expectedError string
	}{
		{},
		{
			executors: []exor.Executor{
				errorRunner{"some-error"},
				exor.NewCommand("wrong-command", "bad-argument"),
			},
			expectedError: "some-error",
		},
		{
			executors: []exor.Executor{
				errorRunner{"some-error-1"},
				errorRunner{"some-error-2"},
			},
			expectedError: "some-error-1",
		},
	}

	for i, tt := range testcases {
		err := exor.Execute(context.Background(), tt.executors...)
		if tt.expectedError == "" {
			require.NoError(t, err)
		} else {
			require.EqualError(t, err, tt.expectedError, i)
		}
	}
}

type errorRunner struct {
	msg string
}

func (r errorRunner) Execute(ctx context.Context) error {
	if r.msg == "" {
		return nil
	}
	return errors.New(r.msg)
}
