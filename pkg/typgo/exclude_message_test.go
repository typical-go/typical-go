package typgo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestExcludeMessage(t *testing.T) {
	require.True(t, typgo.ExcludeMessage("Merge something"))
	require.True(t, typgo.ExcludeMessage("merge something"))
	require.False(t, typgo.ExcludeMessage("asdf"))
}
