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
