package typgo_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestBuildSys_Run(t *testing.T) {
	sys := &typgo.BuildSys{
		Commands: []*cli.Command{
			{
				Name:   "cmd1",
				Action: func(*cli.Context) error { return errors.New("some-error") },
			},
			{
				Name:   "cmd2",
				Action: func(*cli.Context) error { return nil },
			},
		},
	}
	c := &cli.Context{}
	require.EqualError(t, sys.Run("noname", c), "typgo: no command with name 'noname'")
	require.EqualError(t, sys.Run("cmd1", c), "some-error")
	require.NoError(t, sys.Run("cmd2", c))

}
