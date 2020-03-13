package typcore_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestNewConfigManager(t *testing.T) {
	t.Run("", func(t *testing.T) {
		loader := typcore.NewConfigManager().Loader()
		require.Equal(t, "*typcore.defaultLoader", reflect.TypeOf(loader).String())
	})
	t.Run("SHOULD implement of typcore.ConfigManager", func(t *testing.T) {
		var _ typcore.ConfigManager = typcore.NewConfigManager()
	})
}

func TestConfiguration(t *testing.T) {
	configuration := typcore.NewConfigManager().
		WithLoader(&dummyLoader{}).
		WithConfigurer(&dummyConfigurer1{}, &dummyConfigurer2{})

	require.IsType(t, &dummyLoader{}, configuration.Loader())

	var b strings.Builder
	require.NoError(t, configuration.Write(&b))
	require.Equal(t, "prefix1_ID=\nprefix1_VOLUME=\nprefix2_TITLE=default-title\nprefix2_CONTENT=default-content\n", b.String())
}

type dummyLoader struct{}

func (*dummyLoader) LoadConfig(string, interface{}) error { return nil }

type dummyConfigurer1 struct{}

func (*dummyConfigurer1) Configure() *typcore.Configuration {
	return typcore.NewConfiguration("prefix1", &struct {
		ID     int64 ``
		Volume int   ``
	}{})
}

type dummyConfigurer2 struct{}

func (*dummyConfigurer2) Configure() *typcore.Configuration {
	return typcore.NewConfiguration("prefix2", &struct {
		Title   string `default:"default-title"`
		Content string `default:"default-content"`
	}{})
}
