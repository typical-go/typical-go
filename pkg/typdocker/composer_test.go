package typdocker_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typdocker"
)

func TestNewCompose(t *testing.T) {
	expectedErr := errors.New("some-error")
	expectedRecipe := &typdocker.Recipe{}

	composer := typdocker.NewCompose(func() (*typdocker.Recipe, error) {
		return expectedRecipe, expectedErr
	})

	recipe, err := composer.Compose()
	require.Equal(t, expectedRecipe, recipe)
	require.Equal(t, expectedErr, err)
}
