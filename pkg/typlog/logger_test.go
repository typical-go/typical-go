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

	expected := `INFO: (NONAME) some information
INFO: (NONAME) formatted information: FOO
WARN: (NONAME) some warning
WARN: (NONAME) formatted warning: BAR
`

	require.Equal(t, expected, debugger.String())
}
