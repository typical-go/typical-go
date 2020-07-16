package execkit_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
)

func TestHelper(t *testing.T) {
	err := execkit.Run(
		context.Background(),
		execkit.NewRunner(func(context.Context) error {
			return errors.New("some-error")
		}),
	)
	require.EqualError(t, err, "some-error")
}
