package walker_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcore/walker"
)

func TestParseAnnotation(t *testing.T) {
	testcases := []struct {
		raw string
		*walker.Annotation
	}{
		{`no-bracket`, nil},
		{`[autowire]`, walker.NewAnnotation("autowire")},
		{`[mock(pkg=mock2)]`, walker.NewAnnotation("mock").Put("pkg", "mock2")},
		{`[mock(pkg="mock2")]`, walker.NewAnnotation("mock").Put("pkg", "mock2")},
		{`[mock(pkg="mock2"]`, walker.NewAnnotation("mock")},
		{`[mock(pkg)]`, walker.NewAnnotation("mock").Put("pkg", "")},
		{`[noname(key1="value1" key2="value2")]`,
			walker.NewAnnotation("noname").Put("key1", "value1").Put("key2", "value2")},
		{`[noname(key1=value1 key2=value2)]`,
			walker.NewAnnotation("noname").Put("key1", "value1").Put("key2", "value2")},
		{`[noname(key1=value1 key2="value2")]`,
			walker.NewAnnotation("noname").Put("key1", "value1").Put("key2", "value2")},
		{`[noname(key1=value1 key2 key3=value3)]`,
			walker.NewAnnotation("noname").Put("key1", "value1").Put("key2", "").Put("key3", "value3")},
		{`[noname(key1= key2 key3="")]`,
			walker.NewAnnotation("noname").Put("key1", "").Put("key2", "").Put("key3", "")},
		{`[noname(key1="" key2 key3=)]`,
			walker.NewAnnotation("noname").Put("key1", "").Put("key2", "").Put("key3", "")},
		{`[noname(key1="" key2 key3)]`,
			walker.NewAnnotation("noname").Put("key1", "").Put("key2", "").Put("key3", "")},
		{`[noname(key1="" key2 key3 key4=value4)]`,
			walker.NewAnnotation("noname").Put("key1", "").Put("key2", "").Put("key3", "").Put("key4", "value4")},
	}
	for i, tt := range testcases {
		require.Equal(t, tt.Annotation, walker.ParseAnnotation(tt.raw), i)
	}
}

func TestParseAnnotations(t *testing.T) {
	raw := "[tag1] some text [tag2] then another text [tag3(key1=value1)]"
	require.Equal(t, walker.Annotations{
		walker.NewAnnotation("tag1"),
		walker.NewAnnotation("tag2"),
		walker.NewAnnotation("tag3").Put("key1", "value1"),
	}, walker.ParseAnnotations(raw))
}
