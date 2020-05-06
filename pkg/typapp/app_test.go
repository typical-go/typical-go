package typapp_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/typapp"
)

func TestProvide(t *testing.T) {
	c1 := &typapp.Constructor{}
	c2 := &typapp.Constructor{}
	c3 := &typapp.Constructor{}
	app := typapp.EntryPoint(nil, "").Imports(c1, c2)
	typapp.Provide(c3)

	require.EqualValues(t,
		[]*typapp.Constructor{c1, c2, c3},
		app.Constructors(),
	)
}

func TestDestoy(t *testing.T) {
	i1 := &typapp.Destructor{}
	i2 := &typapp.Destructor{}
	i3 := &typapp.Destructor{}
	app := typapp.EntryPoint(nil, "").Imports(i1, i2, i3)

	require.EqualValues(t,
		[]*typapp.Destructor{i1, i2, i3},
		app.Destructors(),
	)
}

func TestPrepare(t *testing.T) {
	i1 := &typapp.Preparation{}
	i2 := &typapp.Preparation{}
	i3 := &typapp.Preparation{}
	app := typapp.EntryPoint(nil, "").Imports(i1, i2, i3)

	require.EqualValues(t,
		[]*typapp.Preparation{i1, i2, i3},
		app.Preparations(),
	)
}
