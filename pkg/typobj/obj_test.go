package typobj_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typobj"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typobj"
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
		require.Equal(t, tt.isProvider, typobj.IsProvider(tt.obj), i)
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
		require.Equal(t, tt.isPreparer, typobj.IsPreparer(tt.obj), i)
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
		require.Equal(t, tt.isDestroyer, typobj.IsDestroyer(tt.obj), i)
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
		require.Equal(t, tt.isConfigurer, typobj.IsConfigurer(tt.obj), i)
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
		require.Equal(t, tt.isConfigurer, typobj.IsValidator(tt.obj), i)
	}
}

type dummyObj struct{}

func (dummyObj) Run() interface{}                                                 { return nil }
func (dummyObj) Prepare() []interface{}                                           { return nil }
func (dummyObj) Provide() []interface{}                                           { return nil }
func (dummyObj) Destroy() []interface{}                                           { return nil }
func (dummyObj) Configure() (prefix string, spec interface{}, loadFn interface{}) { return "", nil, nil }
func (dummyObj) Validate() error                                                  { return nil }
