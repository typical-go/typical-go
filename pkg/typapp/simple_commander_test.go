package typapp_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/urfave/cli/v2"
)

func TestSimpleCommander_Commands(t *testing.T) {
	c1 := &cli.Command{}
	c2 := &cli.Command{}
	c3 := &cli.Command{}
	cmd := typapp.NewCommander(
		func(*typapp.Context) []*cli.Command {
			return []*cli.Command{c1, c2}
		},
		func(*typapp.Context) []*cli.Command {
			return []*cli.Command{c3}
		},
	)
	require.Equal(t, []*cli.Command{c1, c2, c3}, cmd.Commands(nil))

}
