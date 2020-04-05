package typapp_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/typapp"
)

func TestProvide(t *testing.T) {
	c1 := typapp.NewConstructor(nil)
	c2 := typapp.NewConstructor(nil)
	c3 := typapp.NewConstructor(nil)
	app := typapp.EntryPoint(nil, "").
		WithModules(c1, c2)
	typapp.AppendConstructor(c3)

	require.EqualValues(t,
		[]*typapp.Constructor{c1, c2, c3},
		app.Provide(),
	)
}

func TestDestoy(t *testing.T) {
	i1 := typapp.NewDestruction(nil)
	i2 := typapp.NewDestruction(nil)
	i3 := typapp.NewDestruction(nil)
	app := typapp.EntryPoint(nil, "").WithModules(i1, i2, i3)

	require.EqualValues(t,
		[]*typapp.Destruction{i1, i2, i3},
		app.Destroy(),
	)
}

func TestPrepare(t *testing.T) {
	i1 := typapp.NewPreparation(nil)
	i2 := typapp.NewPreparation(nil)
	i3 := typapp.NewPreparation(nil)
	app := typapp.EntryPoint(nil, "").WithModules(i1, i2, i3)

	require.EqualValues(t,
		[]*typapp.Preparation{i1, i2, i3},
		app.Prepare(),
	)
}
