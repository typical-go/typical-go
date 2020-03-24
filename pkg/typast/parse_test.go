package typast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typast"
)

func TestParseAnnotation(t *testing.T) {
	testcases := []struct {
		raw string
		*typast.Annotation
	}{
		{`no-bracket`, nil},
		{`[autowire]`, typast.NewAnnotation("autowire")},
		{`[mock(pkg=mock2)]`, typast.NewAnnotation("mock").Put("pkg", "mock2")},
		{`[mock(pkg="mock2")]`, typast.NewAnnotation("mock").Put("pkg", "mock2")},
		{`[mock(pkg="mock2"]`, typast.NewAnnotation("mock")},
		{`[mock(pkg)]`, typast.NewAnnotation("mock").Put("pkg", "")},
		{
			`[noname(key1="value1" key2="value2")]`,
			typast.NewAnnotation("noname").Put("key1", "value1").Put("key2", "value2"),
		},
		{
			`[noname(key1=value1 key2=value2)]`,
			typast.NewAnnotation("noname").Put("key1", "value1").Put("key2", "value2"),
		},
		{
			`[noname(key1=value1 key2="value2")]`,
			typast.NewAnnotation("noname").Put("key1", "value1").Put("key2", "value2"),
		},
		{
			`[noname(key1=value1 key2 key3=value3)]`,
			typast.NewAnnotation("noname").Put("key1", "value1").Put("key2", "").Put("key3", "value3"),
		},
		{
			`[noname(key1= key2 key3="")]`,
			typast.NewAnnotation("noname").Put("key1", "").Put("key2", "").Put("key3", ""),
		},
		{
			`[noname(key1="" key2 key3=)]`,
			typast.NewAnnotation("noname").Put("key1", "").Put("key2", "").Put("key3", ""),
		},
		{
			`[noname(key1="" key2 key3)]`,
			typast.NewAnnotation("noname").Put("key1", "").Put("key2", "").Put("key3", ""),
		},
		{
			`[noname(key1="" key2 key3 key4=value4)]`,
			typast.NewAnnotation("noname").Put("key1", "").Put("key2", "").Put("key3", "").Put("key4", "value4"),
		},
	}
	for i, tt := range testcases {
		require.Equal(t, tt.Annotation, typast.ParseAnnotation(tt.raw), i)
	}
}

func TestParseAnnotations(t *testing.T) {
	raw := "[tag1] some text [tag2] then another text [tag3(key1=value1)]"
	require.Equal(t, []*typast.Annotation{
		typast.NewAnnotation("tag1"),
		typast.NewAnnotation("tag2"),
		typast.NewAnnotation("tag3").Put("key1", "value1"),
	}, typast.ParseAnnotations(raw))
}
