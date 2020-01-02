package typcore_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestField(t *testing.T) {
	names, configMap := typcore.ConfigMap("TEST", &sampleSpec{
		Field2: "mama",
	})
	require.Equal(t, []string{"TEST_FIELD1", "TEST_FIELD2", "TEST_ALIAS"}, names)
	require.EqualValues(t, map[string]typcore.ConfigDetail{
		"TEST_FIELD1": {Name: "TEST_FIELD1", Type: "string", Default: "hello", Value: "", IsZero: true, Required: true},
		"TEST_FIELD2": {Name: "TEST_FIELD2", Type: "string", Default: "world", Value: "mama", IsZero: false, Required: false},
		"TEST_ALIAS":  {Name: "TEST_ALIAS", Type: "int", Default: "", Value: 0, IsZero: true, Required: false},
	}, configMap)
}

type sampleSpec struct {
	Field1 string `default:"hello" required:"true"`
	Field2 string `default:"world"`
	Field3 string `ignored:"true"`
	Field4 int    `envconfig:"ALIAS"`
}
