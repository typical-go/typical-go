package typgo_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestBuildSys_Run(t *testing.T) {
	var seq []string
	sys := &typgo.BuildSys{
		Commands: []*cli.Command{
			{
				Name:   "cmd1",
				Action: func(*cli.Context) error { return errors.New("some-error") },
			},
			{
				Name:   "cmd2",
				Before: func(*cli.Context) error { seq = append(seq, "1"); return nil },
				Action: func(*cli.Context) error { seq = append(seq, "2"); return nil },
			},
			{
				Name:   "cmd3",
				Before: func(*cli.Context) error { return errors.New("before-error") },
				Action: func(*cli.Context) error { return errors.New("some-error") },
			},
		},
	}
	c := &cli.Context{}
	require.EqualError(t, sys.Run("noname", c), "typgo: no command with name 'noname'")
	require.EqualError(t, sys.Run("cmd1", c), "some-error")
	require.EqualError(t, sys.Run("cmd3", c), "before-error")
	require.NoError(t, sys.Run("cmd2", c))
	require.Equal(t, []string{"1", "2"}, seq)
}

func TestBuildSys_ActionFn(t *testing.T) {
	sys := &typgo.BuildSys{}
	require.NoError(t, sys.ActionFn(nil)(&cli.Context{}))
	action := typgo.NewAction(func(*typgo.Context) error {
		return errors.New("some-error")
	})
	require.EqualError(t, sys.ActionFn(action)(&cli.Context{}), "some-error")
}
