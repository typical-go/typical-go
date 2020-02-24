package typcfg_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/typical-go/typical-go/pkg/typdep"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestNewConfiguration(t *testing.T) {
	t.Run("New configuration instance using default config loader", func(t *testing.T) {
		loader := typcfg.New().Loader()
		require.Equal(t, "*typcfg.defaultLoader", reflect.TypeOf(loader).String())
	})
	t.Run("Configuration must implement of typcore.Configuration", func(t *testing.T) {
		var _ typcore.Configuration = typcfg.New()
	})
}

func TestConfiguration(t *testing.T) {
	configuration := typcfg.New().
		WithLoader(&dummyLoader{}).
		WithConfigure(&dummyConfigurer1{}, &dummyConfigurer2{})
	store := configuration.Store()
	require.IsType(t, &dummyLoader{}, configuration.Loader())
	require.Equal(t, []*typdep.Constructor{constructor1, constructor2}, store.Provide())

	require.EqualValues(t, []*typcore.ConfigField{
		{Name: "prefix1_ID", Type: "int64", Default: "", Value: int64(0), IsZero: true, Required: false},
		{Name: "prefix1_VOLUME", Type: "int", Default: "", Value: 0, IsZero: true, Required: false},
		{Name: "prefix2_TITLE", Type: "string", Default: "default-title", Value: "", IsZero: true, Required: false},
		{Name: "prefix2_CONTENT", Type: "string", Default: "default-content", Value: "", IsZero: true, Required: false},
	}, store.Fields(store.Keys()...))

	var b strings.Builder
	require.NoError(t, configuration.Write(&b))
	require.Equal(t, "prefix1_ID=\nprefix1_VOLUME=\nprefix2_TITLE=default-title\nprefix2_CONTENT=default-content\n", b.String())
}

var (
	constructor1 *typdep.Constructor
	constructor2 *typdep.Constructor
)

type dummyLoader struct{}

func (*dummyLoader) Load(string, interface{}) error { return nil }

type dummyConfigurer1 struct{}

func (*dummyConfigurer1) Configure(loader typcfg.Loader) *typcfg.Detail {
	return &typcfg.Detail{
		Prefix: "prefix1",
		Spec: &struct {
			ID     int64 ``
			Volume int   ``
		}{},
		Constructor: constructor1,
	}
}

type dummyConfigurer2 struct{}

func (*dummyConfigurer2) Configure(loader typcfg.Loader) *typcfg.Detail {
	return &typcfg.Detail{
		Prefix: "prefix2",
		Spec: &struct {
			Title   string `default:"default-title"`
			Content string `default:"default-content"`
		}{},
		Constructor: constructor2,
	}
}
