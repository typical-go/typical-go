package typgo_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestTask(t *testing.T) {

	task := &typgo.Task{
		Name:            "some-name",
		Aliases:         []string{"some-alias"},
		Usage:           "some-usage",
		Flags:           []cli.Flag{&cli.StringFlag{Name: "some-flag"}},
		SkipFlagParsing: true,

		Action: typgo.NewAction(func(*typgo.Context) error {
			return errors.New("some-error")
		}),
	}

	require.Equal(t, task, task.Task())

	cmd := typgo.CliCommand(&typgo.Descriptor{}, task)

	require.Equal(t, "some-name", cmd.Name)
	require.Equal(t, []string{"some-alias"}, cmd.Aliases)
	require.Equal(t, "some-usage", cmd.Usage)
	require.Equal(t, []cli.Flag{&cli.StringFlag{Name: "some-flag"}}, cmd.Flags)
	require.Equal(t, true, cmd.SkipFlagParsing)

}
