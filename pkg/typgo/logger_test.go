package typgo_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestLogger(t *testing.T) {
	var out strings.Builder
	c := typgo.Logger{
		Stdout:      &out,
		ProjectName: "some-project",
		TaskNames:   []string{"some-command"},
	}
	c.Info("some-info")
	c.Infof("some-info %d\n", 9999)
	c.Warn("some-warning")
	c.Warnf("some-warning %d\n", 9999)
	require.Equal(t, "some-project:some-command> some-info\nsome-project:some-command> some-info 9999\nsome-project:some-command> some-warning\nsome-project:some-command> some-warning 9999\n", out.String())
}

func TestLogger_NoPanic(t *testing.T) {
	c := typgo.Logger{}
	c.Info("some-info")
	c.Infof("some-info %d\n", 9999)
	c.Warn("some-warning")
	c.Warnf("some-warning %d\n", 9999)
}
