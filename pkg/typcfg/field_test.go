package typcfg_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcfg"
)

func TestField(t *testing.T) {
	require.EqualValues(t, []typcfg.Field{
		{Name: "TEST_FIELD1", Type: "string", Default: "hello", Required: true},
		{Name: "TEST_FIELD2", Type: "string", Default: "world", Required: false},
		{Name: "TEST_ALIAS", Type: "int", Default: "", Required: false},
	}, typcfg.Fields("TEST", &SampleSpec{}))
}
