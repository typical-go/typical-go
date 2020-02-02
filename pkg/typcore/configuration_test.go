package typcore_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestNewConfiguration(t *testing.T) {
	t.Run("New configuration instance using default config loader", func(t *testing.T) {
		loader := typcore.NewConfiguration().Loader()
		require.Equal(t, "*typcore.defaultConfigLoader", reflect.TypeOf(loader).String())
	})
	t.Run("Configuration must implement of ConfigurationInterface", func(t *testing.T) {
		var _ typcore.ConfigurationInterface = typcore.NewConfiguration()
	})
}

func TestConfiguration(t *testing.T) {
	configuration := typcore.NewConfiguration().
		WithLoader(&dummyLoader{}).
		WithConfigure(&dummyConfigurer1{}, &dummyConfigurer2{})
	require.IsType(t, &dummyLoader{}, configuration.Loader())
	require.Equal(t, []interface{}{"loadFn1", "loadFn2"}, configuration.Provide())
	keys, configMap := configuration.ConfigMap()
	require.Equal(t, typcore.ConfigDetails{
		{Name: "prefix1_ID", Type: "int64", Default: "", Value: int64(0), IsZero: true, Required: false},
		{Name: "prefix1_VOLUME", Type: "int", Default: "", Value: 0, IsZero: true, Required: false},
		{Name: "prefix2_TITLE", Type: "string", Default: "", Value: "", IsZero: true, Required: false},
		{Name: "prefix2_CONTENT", Type: "string", Default: "", Value: "", IsZero: true, Required: false},
	}, configMap.ValueBy(keys...))
}

type dummyLoader struct{}

func (*dummyLoader) Load(string, interface{}) error { return nil }

type dummyConfigurer1 struct{}

func (*dummyConfigurer1) Configure(loader typcore.ConfigLoader) (string, interface{}, interface{}) {
	return "prefix1",
		&struct {
			ID     int64
			Volume int
		}{},
		"loadFn1"
}

type dummyConfigurer2 struct{}

func (*dummyConfigurer2) Configure(loader typcore.ConfigLoader) (string, interface{}, interface{}) {
	return "prefix2", &struct {
		Title   string
		Content string
	}{}, "loadFn2"
}
