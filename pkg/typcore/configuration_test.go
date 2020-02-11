package typcore_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestConfigMap_ValueBy(t *testing.T) {
	m := map[string]typcore.ConfigDetail{
		"key1": configDetail("key1"),
		"key2": configDetail("key2"),
		"key3": configDetail("key3"),
		"key4": configDetail("key4"),
	}
	require.Equal(t, []typcore.ConfigDetail{
		configDetail("key4"),
		configDetail("key1"),
	}, typcore.ConfigDetailsBy(m, "key4", "key1"))
	require.Equal(t, []typcore.ConfigDetail{
		configDetail("key1"),
	}, typcore.ConfigDetailsBy(m, "key1", "not-available"))
}

func configDetail(name string) typcore.ConfigDetail {
	return typcore.ConfigDetail{
		Name: name,
	}
}
