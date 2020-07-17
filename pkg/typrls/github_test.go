package typrls_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typrls"
)

func TestExcludeMessage(t *testing.T) {
	require.True(t, typrls.ExcludeMessage("Merge something"))
	require.True(t, typrls.ExcludeMessage("merge something"))
	require.False(t, typrls.ExcludeMessage("asdf"))
}
