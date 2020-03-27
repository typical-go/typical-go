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
	t.Run("GIVEN empty attribute", func(t *testing.T) {
		mod := &typapp.Module{}
		require.Equal(t, []*cli.Command{}, mod.Commands(nil))
		require.Equal(t, []*typapp.Preparation{}, mod.Prepare())
		require.Equal(t, []*typapp.Constructor{}, mod.Provide())
		require.Equal(t, []*typapp.Destruction{}, mod.Destroy())
	})
	t.Run("GIVE some attribute", func(t *testing.T) {
		cmd := &cli.Command{}
		preparation := typapp.NewPreparation(nil)
		constructor := typapp.NewConstructor(nil)
		destruction := typapp.NewDestruction(nil)

		mod := &typapp.Module{
			Commander: typapp.NewCommander(func(*typapp.Context) []*cli.Command {
				return []*cli.Command{cmd}
			}),
			Preparer:  preparation,
			Provider:  constructor,
			Destroyer: destruction,
		}
		require.Equal(t, []*cli.Command{cmd}, mod.Commands(nil))
		require.Equal(t, []*typapp.Preparation{preparation}, mod.Prepare())
		require.Equal(t, []*typapp.Constructor{constructor}, mod.Provide())
		require.Equal(t, []*typapp.Destruction{destruction}, mod.Destroy())
	})
}
