package typcore_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestRunner(t *testing.T) {
	runner := typcore.Run(func(*typcore.Descriptor) error {
		return errors.New("some-error")
	})
	require.EqualError(t, runner.Run(nil), "some-error")

}
