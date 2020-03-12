package typcore_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typdep"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestConfigStore(t *testing.T) {

	store := typcore.NewConfigStore()
	c1 := typdep.NewConstructor(nil)
	c2 := typdep.NewConstructor(nil)
	f1 := &typcore.ConfigField{}
	f2 := &typcore.ConfigField{}
	f3 := &typcore.ConfigField{}
	f4 := &typcore.ConfigField{}

	store.Put("bean1", typcore.NewConfigBean("bean1", []*typcore.ConfigField{f1, f2}, c1))
	store.Put("bean2", typcore.NewConfigBean("bean2", []*typcore.ConfigField{f3, f4}, c2))

	require.EqualValues(t, []*typdep.Constructor{c1, c2}, store.Provide())
	require.EqualValues(t, []*typcore.ConfigField{f1, f2, f3, f4}, store.Fields())

}
