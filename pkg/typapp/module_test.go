package typapp_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typcfg"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/urfave/cli/v2"
)

func TestNewModule(t *testing.T) {
	cmd := &cli.Command{}
	preparation := typapp.NewPreparation(nil)
	constructor := typapp.NewConstructor(nil)
	destruction := typapp.NewDestruction(nil)
	configuration := typcfg.NewConfiguration("", nil)

	mod := typapp.NewModule().
		Commanders(typapp.NewCommander(func(*typapp.Context) []*cli.Command {
			return []*cli.Command{cmd}
		})).
		Prepares(preparation).
		Destroys(destruction).
		Provides(constructor).
		Configures(configuration)

	require.Equal(t, []*cli.Command{cmd}, mod.Commands(nil))
	require.Equal(t, []*typapp.Preparation{preparation}, mod.Preparations())
	require.Equal(t, []*typapp.Constructor{constructor}, mod.Constructors())
	require.Equal(t, []*typapp.Destruction{destruction}, mod.Destructions())
	require.Equal(t, []*typcfg.Configuration{configuration}, mod.Configurations())
}
