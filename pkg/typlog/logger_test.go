package typlog_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typlog"
)

func TestLogger(t *testing.T) {
	var debugger strings.Builder
	logger := typlog.Logger{
		Out:  &debugger,
		Name: "NONAME",
	}

	logger.Info("some information")
	logger.Infof("formatted information: %s", "FOO")
	logger.Warn("some warning")
	logger.Warnf("formatted warning: %s", "BAR")

	expected := `NONAME:INFO> some information
NONAME:INFO> formatted information: FOO
NONAME:WARN> some warning
NONAME:WARN> formatted warning: BAR
`

	require.Equal(t, expected, debugger.String())
}
