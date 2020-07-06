package typapp_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func TestDestructor(t *testing.T) {
	dtor := &typapp.Destructor{}
	dtor.Destructors()
	require.Equal(t, []*typapp.Destructor{dtor}, dtor.Destructors())
}

func TestDestroys(t *testing.T) {
	dtor1 := &typapp.Destructor{}
	dtor2 := &typapp.Destructor{}
	dtor3 := &typapp.Destructor{}
	dtor4 := &typapp.Destructor{}
	destroyers := typapp.Destroyers{
		dtor1,
		dtor2,
		typapp.Destroyers{
			dtor3,
			dtor4,
		},
	}
	require.Equal(t, []*typapp.Destructor{
		dtor1,
		dtor2,
		dtor3,
		dtor4,
	}, destroyers.Destructors())
}
