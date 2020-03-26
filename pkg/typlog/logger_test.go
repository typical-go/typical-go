package typlog_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typlog"
)

func TestLogger(t *testing.T) {
	var debugger strings.Builder

	logger := typlog.New().
		WithWriter(&debugger)

	logger.Info("some information")
	logger.Infof("formatted information: %s", "FOO")
	logger.Warn("some warning")
	logger.Warnf("formatted warning: %s", "BAR")

	require.Equal(t, `[TYPICAL][INFO] some information
[TYPICAL][INFO] formatted information: FOO
[TYPICAL][WARN] some warning
[TYPICAL][WARN] formatted warning: BAR
`, debugger.String())
}
