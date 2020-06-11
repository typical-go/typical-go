package typgo_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestContext_Execute(t *testing.T) {
	var debugger strings.Builder
	tmpWriter := typgo.CtxExecWriter
	typgo.CtxExecWriter = &debugger
	defer func() {
		typgo.CtxExecWriter = tmpWriter
	}()

	c := &typgo.Context{
		Context: &cli.Context{},
	}
	err := c.Execute(&dummyExec{})
	require.EqualError(t, err, "some-err")
	require.Equal(t, "\n$ dummy-exec\n", debugger.String())
}

type dummyExec struct{}

func (*dummyExec) Run(ctx context.Context) error {
	return errors.New("some-err")
}

func (dummyExec) String() string {
	return "dummy-exec"
}
