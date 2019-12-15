package typcore_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestField(t *testing.T) {
	require.EqualValues(t, []typcore.Field{
		{Name: "TEST_FIELD1", Type: "string", Default: "hello", Value: "", IsZero: true, Required: true},
		{Name: "TEST_FIELD2", Type: "string", Default: "world", Value: "mama", IsZero: false, Required: false},
		{Name: "TEST_ALIAS", Type: "int", Default: "", Value: 0, IsZero: true, Required: false},
	}, typcore.Fields("TEST", &SampleSpec{
		Field2: "mama",
	}))
}

type SampleSpec struct {
	Field1 string `default:"hello" required:"true"`
	Field2 string `default:"world"`
	Field3 string `ignored:"true"`
	Field4 int    `envconfig:"ALIAS"`
}
