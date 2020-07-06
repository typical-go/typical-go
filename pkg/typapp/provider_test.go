package typapp_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func TestConstructor(t *testing.T) {
	ctor := &typapp.Constructor{}
	require.Equal(t, []*typapp.Constructor{ctor}, ctor.Constructors())
}

func TestProviders(t *testing.T) {
	ctor1 := &typapp.Constructor{}
	ctor2 := &typapp.Constructor{}
	ctor3 := &typapp.Constructor{}
	ctor4 := &typapp.Constructor{}
	providers := typapp.Providers{
		ctor1,
		ctor2,
		typapp.Providers{
			ctor3,
			ctor4,
		},
	}
	require.Equal(t, []*typapp.Constructor{
		ctor1,
		ctor2,
		ctor3,
		ctor4,
	}, providers.Constructors())
}
