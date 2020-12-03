package typrls_test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/oskit"
	"github.com/typical-go/typical-go/pkg/typrls"
)

func TestGithub_Publish(t *testing.T) {
	var out strings.Builder
	defer oskit.PatchStdout(&out)()

	os.Unsetenv("GITHUB_TOKEN")
	github := &typrls.Github{}
	require.EqualError(t, github.Publish(nil), "github-release: missing $GITHUB_TOKEN")
}
