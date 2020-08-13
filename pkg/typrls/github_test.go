package typrls_test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typrls"
)

func TestGithub_Release(t *testing.T) {
	var out strings.Builder
	typrls.Stdout = &out
	defer func() { typrls.Stdout = os.Stdout }()

	os.Unsetenv("GITHUB_TOKEN")
	github := &typrls.Github{}
	require.NoError(t, github.Release(nil))
	require.Equal(t, "Skip Github Release due to missing 'GITHUB_TOKEN'\n", out.String())
}
