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
	command := task.Task(&typgo.BuildSys{})
	require.Equal(t, "some-name", command.Name)
	require.Equal(t, []string{"some-alias"}, command.Aliases)
	require.Equal(t, "some-usage", command.Usage)
	require.Equal(t, []cli.Flag{&cli.StringFlag{Name: "some-flag"}}, command.Flags)
	require.Equal(t, true, command.SkipFlagParsing)
	require.EqualError(t, command.Action(&cli.Context{}), "some-error")
}
