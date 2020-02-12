package prebld_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typbuild/prebld"
)

func TestParseAnnotation(t *testing.T) {
	testcases := []struct {
		raw string
		*prebld.Annotation
	}{
		{`no-bracket`, nil},
		{`[autowire]`, prebld.NewAnnotation("autowire")},
		{`[mock(pkg=mock2)]`, prebld.NewAnnotation("mock").Put("pkg", "mock2")},
		{`[mock(pkg="mock2")]`, prebld.NewAnnotation("mock").Put("pkg", "mock2")},
		{`[mock(pkg="mock2"]`, prebld.NewAnnotation("mock")},
		{`[mock(pkg)]`, prebld.NewAnnotation("mock").Put("pkg", "")},
		{
			`[noname(key1="value1" key2="value2")]`,
			prebld.NewAnnotation("noname").Put("key1", "value1").Put("key2", "value2"),
		},
		{
			`[noname(key1=value1 key2=value2)]`,
			prebld.NewAnnotation("noname").Put("key1", "value1").Put("key2", "value2"),
		},
		{
			`[noname(key1=value1 key2="value2")]`,
			prebld.NewAnnotation("noname").Put("key1", "value1").Put("key2", "value2"),
		},
		{
			`[noname(key1=value1 key2 key3=value3)]`,
			prebld.NewAnnotation("noname").Put("key1", "value1").Put("key2", "").Put("key3", "value3"),
		},
		{
			`[noname(key1= key2 key3="")]`,
			prebld.NewAnnotation("noname").Put("key1", "").Put("key2", "").Put("key3", ""),
		},
		{
			`[noname(key1="" key2 key3=)]`,
			prebld.NewAnnotation("noname").Put("key1", "").Put("key2", "").Put("key3", ""),
		},
		{
			`[noname(key1="" key2 key3)]`,
			prebld.NewAnnotation("noname").Put("key1", "").Put("key2", "").Put("key3", ""),
		},
		{
			`[noname(key1="" key2 key3 key4=value4)]`,
			prebld.NewAnnotation("noname").Put("key1", "").Put("key2", "").Put("key3", "").Put("key4", "value4"),
		},
	}
	for i, tt := range testcases {
		require.Equal(t, tt.Annotation, prebld.ParseAnnotation(tt.raw), i)
	}
}

func TestParseAnnotations(t *testing.T) {
	raw := "[tag1] some text [tag2] then another text [tag3(key1=value1)]"
	require.Equal(t, prebld.Annotations{
		prebld.NewAnnotation("tag1"),
		prebld.NewAnnotation("tag2"),
		prebld.NewAnnotation("tag3").Put("key1", "value1"),
	}, prebld.ParseAnnotations(raw))
}
