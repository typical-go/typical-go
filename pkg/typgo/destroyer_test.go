package typgo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestDestructor(t *testing.T) {
	dtor := &typgo.Destructor{}
	dtor.Destructors()
	require.Equal(t, []*typgo.Destructor{dtor}, dtor.Destructors())
}

func TestDestroys(t *testing.T) {
	dtor1 := &typgo.Destructor{}
	dtor2 := &typgo.Destructor{}
	dtor3 := &typgo.Destructor{}
	dtor4 := &typgo.Destructor{}
	destroyers := typgo.Destroyers{
		dtor1,
		dtor2,
		typgo.Destroyers{
			dtor3,
			dtor4,
		},
	}
	require.Equal(t, []*typgo.Destructor{
		dtor1,
		dtor2,
		dtor3,
		dtor4,
	}, destroyers.Destructors())
}
