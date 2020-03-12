package typcore_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestNewConfiguration(t *testing.T) {
	t.Run("New configuration instance using default config loader", func(t *testing.T) {
		loader := typcore.NewConfiguration().Loader()
		require.Equal(t, "*typcore.defaultLoader", reflect.TypeOf(loader).String())
	})
	t.Run("SHOULD implement of typcore.Configuration", func(t *testing.T) {
		var _ typcore.Configuration = typcore.NewConfiguration()
	})
}

func TestConfiguration(t *testing.T) {
	configuration := typcore.NewConfiguration().
		WithLoader(&dummyLoader{}).
		AppendConfigurer(&dummyConfigurer1{}, &dummyConfigurer2{})
	store := configuration.Store()
	require.IsType(t, &dummyLoader{}, configuration.Loader())

	require.EqualValues(t, []*typcore.ConfigField{
		{Name: "prefix1_ID", Type: "int64", Default: "", Value: int64(0), IsZero: true, Required: false},
		{Name: "prefix1_VOLUME", Type: "int", Default: "", Value: 0, IsZero: true, Required: false},
		{Name: "prefix2_TITLE", Type: "string", Default: "default-title", Value: "", IsZero: true, Required: false},
		{Name: "prefix2_CONTENT", Type: "string", Default: "default-content", Value: "", IsZero: true, Required: false},
	}, store.Fields())

	var b strings.Builder
	require.NoError(t, configuration.Write(&b))
	require.Equal(t, "prefix1_ID=\nprefix1_VOLUME=\nprefix2_TITLE=default-title\nprefix2_CONTENT=default-content\n", b.String())
}

type dummyLoader struct{}

func (*dummyLoader) LoadConfig(string, interface{}) error { return nil }

type dummyConfigurer1 struct{}

func (*dummyConfigurer1) Configure() *typcore.ConfigBean {
	return &typcore.ConfigBean{
		Name: "prefix1",
		Spec: &struct {
			ID     int64 ``
			Volume int   ``
		}{},
	}
}

type dummyConfigurer2 struct{}

func (*dummyConfigurer2) Configure() *typcore.ConfigBean {
	return &typcore.ConfigBean{
		Name: "prefix2",
		Spec: &struct {
			Title   string `default:"default-title"`
			Content string `default:"default-content"`
		}{},
	}
}
