package typmodule_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typcfg"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typmodule"
)

func TestIsProvider(t *testing.T) {
	testCases := []struct {
		obj        interface{}
		isProvider bool
	}{
		{dummyObj{}, true},
		{struct{}{}, false},
	}
	for i, tt := range testCases {
		require.Equal(t, tt.isProvider, typmodule.IsProvider(tt.obj), i)
	}
}

func TestIsPreparer(t *testing.T) {
	testCases := []struct {
		obj        interface{}
		isPreparer bool
	}{
		{dummyObj{}, true},
		{struct{}{}, false},
	}
	for i, tt := range testCases {
		require.Equal(t, tt.isPreparer, typmodule.IsPreparer(tt.obj), i)
	}
}

func TestIsDestroyer(t *testing.T) {
	testCases := []struct {
		obj         interface{}
		isDestroyer bool
	}{
		{dummyObj{}, true},
		{struct{}{}, false},
	}
	for i, tt := range testCases {
		require.Equal(t, tt.isDestroyer, typmodule.IsDestroyer(tt.obj), i)
	}
}

func TestConfigurer(t *testing.T) {
	testCases := []struct {
		obj          interface{}
		isConfigurer bool
	}{
		{dummyObj{}, true},
		{struct{}{}, false},
	}
	for i, tt := range testCases {
		require.Equal(t, tt.isConfigurer, typcfg.IsConfigurer(tt.obj), i)
	}
}

func TestValidator(t *testing.T) {
	testCases := []struct {
		obj          interface{}
		isConfigurer bool
	}{
		{dummyObj{}, true},
		{struct{}{}, false},
	}
	for i, tt := range testCases {
		require.Equal(t, tt.isConfigurer, typmodule.IsValidator(tt.obj), i)
	}
}

type dummyObj struct{}

func (dummyObj) Run() interface{}                                                 { return nil }
func (dummyObj) Prepare() []interface{}                                           { return nil }
func (dummyObj) Provide() []interface{}                                           { return nil }
func (dummyObj) Destroy() []interface{}                                           { return nil }
func (dummyObj) Configure() (prefix string, spec interface{}, loadFn interface{}) { return "", nil, nil }
func (dummyObj) Validate() error                                                  { return nil }
