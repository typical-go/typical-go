package typcore_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestConfigStore(t *testing.T) {
	var (
		store = &typcore.ConfigStore{}

		c1 = struct{}{}
		c2 = struct{}{}
		f1 = &typcore.ConfigField{}
		f2 = &typcore.ConfigField{}
		f3 = &typcore.ConfigField{}
		f4 = &typcore.ConfigField{}
	)

	store.Add(&typcore.ConfigBean{
		Constructor: c1,
		Keys:        []string{"key1", "key2"},
		FieldMap:    map[string]*typcore.ConfigField{"key1": f1, "key2": f2},
	})
	store.Add(&typcore.ConfigBean{
		Constructor: c2,
		Keys:        []string{"key3", "key4"},
		FieldMap:    map[string]*typcore.ConfigField{"key3": f3, "key4": f4},
	})

	require.EqualValues(t, []interface{}{c1, c2}, store.Provide())
	require.EqualValues(t, []string{"key1", "key2", "key3", "key4"}, store.Keys())
	require.EqualValues(t, map[string]*typcore.ConfigField{"key1": f1, "key2": f2, "key3": f3, "key4": f4}, store.FieldMap())
	require.EqualValues(t, []*typcore.ConfigField{f1, f3}, store.Fields("key1", "key3"))

}
