package typgo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestConstructor(t *testing.T) {
	ctor := &typgo.Constructor{}
	require.Equal(t, []*typgo.Constructor{ctor}, ctor.Constructors())
}

func TestProviders(t *testing.T) {
	ctor1 := &typgo.Constructor{}
	ctor2 := &typgo.Constructor{}
	ctor3 := &typgo.Constructor{}
	ctor4 := &typgo.Constructor{}
	providers := typgo.Providers{
		ctor1,
		ctor2,
		typgo.Providers{
			ctor3,
			ctor4,
		},
	}
	require.Equal(t, []*typgo.Constructor{
		ctor1,
		ctor2,
		ctor3,
		ctor4,
	}, providers.Constructors())
}
