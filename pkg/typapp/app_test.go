package typapp_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typdep"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/urfave/cli/v2"
)

func TestNewApp(t *testing.T) {
	t.Run("Constructor parameter", func(t *testing.T) {
		app := typapp.New(&module{})
		require.Equal(t, someEntryPoint, app.EntryPoint())
		require.Equal(t, []*typdep.Constructor{someConstructor}, app.Provide())
		require.Equal(t, []*typdep.Invocation{somePreparation}, app.Prepare())
		require.Equal(t, []*typdep.Invocation{someDestroyer}, app.Destroy())
		require.Equal(t, []*cli.Command{
			{Name: "cmd1"},
			{Name: "cmd2"},
		}, app.Commands(nil))
	})
}

var (
	someEntryPoint  = typdep.NewInvocation(nil)
	someConstructor = typdep.NewConstructor(nil)
	someDestroyer   = typdep.NewInvocation(nil)
	somePreparation = typdep.NewInvocation(nil)
)

type module struct{}

func (*module) EntryPoint() *typdep.Invocation { return someEntryPoint }

func (*module) Provide() []*typdep.Constructor { return []*typdep.Constructor{someConstructor} }

func (*module) Destroy() []*typdep.Invocation { return []*typdep.Invocation{someDestroyer} }

func (*module) Prepare() []*typdep.Invocation { return []*typdep.Invocation{somePreparation} }

func (*module) Commands(c *typapp.Context) []*cli.Command {
	return []*cli.Command{
		{Name: "cmd1"},
		{Name: "cmd2"},
	}
}
