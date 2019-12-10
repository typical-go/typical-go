package typobj_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typobj"
)

func TestField(t *testing.T) {
	require.EqualValues(t, []typobj.Field{
		{Name: "TEST_FIELD1", Type: "string", Default: "hello", Value: "", IsZero: true, Required: true},
		{Name: "TEST_FIELD2", Type: "string", Default: "world", Value: "mama", IsZero: false, Required: false},
		{Name: "TEST_ALIAS", Type: "int", Default: "", Value: 0, IsZero: true, Required: false},
	}, typobj.Fields("TEST", &SampleSpec{
		Field2: "mama",
	}))
}
