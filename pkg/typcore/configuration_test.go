package typcore_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestCreateConfigDetails(t *testing.T) {
	spec := struct {
		Field1 string `default:"hello" required:"true"`
		Field2 string `default:"world"`
		Field3 string `ignored:"true"`
		Field4 int    `envconfig:"ALIAS"`
	}{
		Field2: "mama",
	}
	require.EqualValues(t, []typcore.ConfigDetail{
		{Name: "TEST_FIELD1", Type: "string", Default: "hello", Value: "", IsZero: true, Required: true},
		{Name: "TEST_FIELD2", Type: "string", Default: "world", Value: "mama", IsZero: false, Required: false},
		{Name: "TEST_ALIAS", Type: "int", Default: "", Value: 0, IsZero: true, Required: false},
	}, typcore.CreateConfigDetails("TEST", &spec))
}

func TestConfigMap_ValueBy(t *testing.T) {
	configMap := typcore.ConfigMap{
		"key1": configDetail("key1"),
		"key2": configDetail("key2"),
		"key3": configDetail("key3"),
		"key4": configDetail("key4"),
	}
	require.Equal(t, typcore.ConfigDetails{
		configDetail("key4"),
		configDetail("key1"),
	}, configMap.ValueBy("key4", "key1"))
	require.Equal(t, typcore.ConfigDetails{
		configDetail("key1"),
	}, configMap.ValueBy("key1", "not-available"))
}

func configDetail(name string) typcore.ConfigDetail {
	return typcore.ConfigDetail{
		Name: name,
	}
}
