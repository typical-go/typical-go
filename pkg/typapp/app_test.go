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
	app := typapp.EntryPoint(nil, "").Imports(c1, c2)
	typapp.AppendConstructor(c3)

	require.EqualValues(t,
		[]*typapp.Constructor{c1, c2, c3},
		app.Constructors(),
	)
}

func TestDestoy(t *testing.T) {
	i1 := typapp.NewDestructor(nil)
	i2 := typapp.NewDestructor(nil)
	i3 := typapp.NewDestructor(nil)
	app := typapp.EntryPoint(nil, "").Imports(i1, i2, i3)

	require.EqualValues(t,
		[]*typapp.Destructor{i1, i2, i3},
		app.Destructors(),
	)
}

func TestPrepare(t *testing.T) {
	i1 := typapp.NewPreparation(nil)
	i2 := typapp.NewPreparation(nil)
	i3 := typapp.NewPreparation(nil)
	app := typapp.EntryPoint(nil, "").Imports(i1, i2, i3)

	require.EqualValues(t,
		[]*typapp.Preparation{i1, i2, i3},
		app.Preparations(),
	)
}
