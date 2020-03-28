package typlog_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typlog"
)

func TestLogger(t *testing.T) {
	var output strings.Builder
	defer typlog.SetOutput(&output)()

	var logger typlog.Logger

	logger.Info("some information")
	logger.Infof("formatted information: %s", "FOO")
	logger.Warn("some warning")
	logger.Warnf("formatted warning: %s", "BAR")

	expected := `[TYPICAL][INFO] some information
[TYPICAL][INFO] formatted information: FOO
[TYPICAL][WARN] some warning
[TYPICAL][WARN] formatted warning: BAR
`

	require.Equal(t, expected, output.String())
}

func TestChangeSignature(t *testing.T) {
	var output strings.Builder
	defer typlog.SetOutput(&output)()

	var logger typlog.Logger
	defer logger.SetLogSignature("TEST", 0)()

	logger.Info("some-info")

	require.Equal(t, "[TEST][INFO] some-info\n", output.String())
}
