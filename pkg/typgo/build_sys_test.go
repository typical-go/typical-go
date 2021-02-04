package typgo_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestBuildSys_Action(t *testing.T) {
	sys := &typgo.Descriptor{}
	action := typgo.NewAction(func(*typgo.Context) error {
		return errors.New("some-error")
	})

	require.NoError(t, sys.Action(nil)(&cli.Context{}))
	require.EqualError(t, sys.Action(action)(&cli.Context{}), "some-error")
}
