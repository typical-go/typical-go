package typapp_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/urfave/cli/v2"
)

func TestModule(t *testing.T) {
	t.Run("SHOULD implement provider", func(t *testing.T) {
		var _ typapp.Provider = &typapp.Module{}
	})
	t.Run("SHOULD implement destroyer", func(t *testing.T) {
		var _ typapp.Destroyer = &typapp.Module{}
	})
	t.Run("SHOULD implement preparer", func(t *testing.T) {
		var _ typapp.Preparer = &typapp.Module{}
	})
	t.Run("SHOULD implement commander", func(t *testing.T) {
		var _ typapp.Commander = &typapp.Module{}
	})

}

func TestNewModule(t *testing.T) {
	cmd := &cli.Command{}
	preparation := typapp.NewPreparation(nil)
	constructor := typapp.NewConstructor(nil)
	destruction := typapp.NewDestruction(nil)

	mod := typapp.NewModule().
		WithCommanders(typapp.NewCommander(func(*typapp.Context) []*cli.Command {
			return []*cli.Command{cmd}
		})).
		WithPrepares(preparation).
		WithDestoyers(destruction).
		WithProviders(constructor)

	require.Equal(t, []*cli.Command{cmd}, mod.Commands(nil))
	require.Equal(t, []*typapp.Preparation{preparation}, mod.Prepare())
	require.Equal(t, []*typapp.Constructor{constructor}, mod.Provide())
	require.Equal(t, []*typapp.Destruction{destruction}, mod.Destroy())
}
