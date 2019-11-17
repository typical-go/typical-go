package typcfg_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcfg"
)

type SampleSpec struct {
	Field1 string `default:"hello" required:"true"`
	Field2 string `default:"world"`
	Field3 string `ignored:"true"`
	Field4 int    `envconfig:"ALIAS"`
}

func TestConfiguration(t *testing.T) {
	configuration := typcfg.Configuration{
		Prefix: "TEST",
		Spec:   &SampleSpec{},
	}
	require.EqualValues(t, []typcfg.ConfigField{
		{Name: "TEST_FIELD1", Type: "string", Default: "hello", Required: true},
		{Name: "TEST_FIELD2", Type: "string", Default: "world", Required: false},
		{Name: "TEST_ALIAS", Type: "int", Default: "", Required: false},
	}, configuration.ConfigFields())
}
