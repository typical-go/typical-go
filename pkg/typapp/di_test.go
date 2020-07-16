package typapp_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func TestAppendConstructor(t *testing.T) {
	c1 := &typapp.Constructor{}
	c2 := &typapp.Constructor{}

	typapp.AppendCtor(c1, c2)
	require.Equal(t, []*typapp.Constructor{c1, c2}, typapp.GetCtors())
}

func TestAppendDestructor(t *testing.T) {
	d1 := &typapp.Destructor{}
	d2 := &typapp.Destructor{}

	typapp.AppendDtor(d1, d2)
	require.Equal(t, []*typapp.Destructor{d1, d2}, typapp.GetDtors())
}
