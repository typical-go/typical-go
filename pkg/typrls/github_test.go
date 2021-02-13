package typrls_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typrls"
)

func TestGithub_Publish(t *testing.T) {
	os.Unsetenv("GITHUB_TOKEN")
	github := &typrls.Github{}
	require.EqualError(t, github.Publish(nil), "github-release: missing $GITHUB_TOKEN")
}
