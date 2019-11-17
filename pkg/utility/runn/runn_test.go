package runn_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/utility/runn"
)

func TestExecute(t *testing.T) {
	err := runn.Execute(
		runner{"some-error-2"},
		runner{"some-error-1"},
	)
	require.EqualError(t, err, "some-error-2")

}
