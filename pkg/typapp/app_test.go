package typapp_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

func TestNewApp(t *testing.T) {
	t.Run("Constructor parameter", func(t *testing.T) {
		app := typapp.New(&module{})
		require.Equal(t, "some-entry-point", app.EntryPoint())
		require.Equal(t, []interface{}{"provide1", "provide2"}, app.Provide())
		require.Equal(t, []interface{}{"prepare1", "prepare2"}, app.Prepare())
		require.Equal(t, []interface{}{"destroy1", "destroy2"}, app.Destroy())
		require.Equal(t, []*cli.Command{
			{Name: "cmd1"},
			{Name: "cmd2"},
		}, app.AppCommands(nil))
	})
	t.Run("With- function", func(t *testing.T) {
		m := &module{}
		app := typapp.New(nil).
			WithEntryPoint(m).
			WithProvide(m).
			WithDestroy(m).
			WithPrepare(m).
			WithCommand(m)
		require.Equal(t, "some-entry-point", app.EntryPoint())
		require.Equal(t, []interface{}{"provide1", "provide2"}, app.Provide())
		require.Equal(t, []interface{}{"prepare1", "prepare2"}, app.Prepare())
		require.Equal(t, []interface{}{"destroy1", "destroy2"}, app.Destroy())
		require.Equal(t, []*cli.Command{
			{Name: "cmd1"},
			{Name: "cmd2"},
		}, app.AppCommands(nil))
	})
}

type module struct{}

func (*module) EntryPoint() interface{} { return "some-entry-point" }

func (*module) Provide() []interface{} { return []interface{}{"provide1", "provide2"} }

func (*module) Destroy() []interface{} { return []interface{}{"destroy1", "destroy2"} }

func (*module) Prepare() []interface{} { return []interface{}{"prepare1", "prepare2"} }

func (*module) AppCommands(c *typcore.AppContext) []*cli.Command {
	return []*cli.Command{
		{Name: "cmd1"},
		{Name: "cmd2"},
	}
}
