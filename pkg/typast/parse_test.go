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
		{
			`no-bracket`,
			nil,
		},
		{
			`[autowire]`,
			typast.NewAnnotation("autowire"),
		},
		{
			`[mock(pkg=mock2)]`,
			typast.NewAnnotation("mock").PutAttr("pkg", "mock2"),
		},
		{
			`[mock(pkg="mock2")]`,
			typast.NewAnnotation("mock").PutAttr("pkg", "mock2"),
		},
		{
			`[mock(pkg="mock2"]`,
			typast.NewAnnotation("mock"),
		},
		{
			`[mock(pkg)]`,
			typast.NewAnnotation("mock").PutAttr("pkg", ""),
		},
		{
			`[noname(key1="value1" key2="value2")]`,
			typast.NewAnnotation("noname").PutAttr("key1", "value1").PutAttr("key2", "value2"),
		},
		{
			`[noname(key1=value1 key2=value2)]`,
			typast.NewAnnotation("noname").PutAttr("key1", "value1").PutAttr("key2", "value2"),
		},
		{
			`[noname(key1=value1 key2="value2")]`,
			typast.NewAnnotation("noname").PutAttr("key1", "value1").PutAttr("key2", "value2"),
		},
		{
			`[noname(key1=value1 key2 key3=value3)]`,
			typast.NewAnnotation("noname").PutAttr("key1", "value1").PutAttr("key2", "").PutAttr("key3", "value3"),
		},
		{
			`[noname(key1= key2 key3="")]`,
			typast.NewAnnotation("noname").PutAttr("key1", "").PutAttr("key2", "").PutAttr("key3", ""),
		},
		{
			`[noname(key1="" key2 key3=)]`,
			typast.NewAnnotation("noname").PutAttr("key1", "").PutAttr("key2", "").PutAttr("key3", ""),
		},
		{
			`[noname(key1="" key2 key3)]`,
			typast.NewAnnotation("noname").PutAttr("key1", "").PutAttr("key2", "").PutAttr("key3", ""),
		},
		{
			`[noname(key1="" key2 key3 key4=value4)]`,
			typast.NewAnnotation("noname").PutAttr("key1", "").PutAttr("key2", "").PutAttr("key3", "").PutAttr("key4", "value4"),
		},
	}
	for i, tt := range testcases {
		require.Equal(t, tt.Annotation, typast.ParseAnnotation(tt.raw), i)
	}
}

// func TestParseAnnotations(t *testing.T) {

// 	raw := "[tag1] some text [tag2] then another text [tag3(key1=value1)]"
// 	require.Equal(t, []*typast.Annot{
// 		typast.NewAnnot("tag1"),
// 		typast.NewAnnot("tag2"),
// 		typast.NewAnnot("tag3").PutAttr("key1", "value1"),
// 	}, typast.ParseAnnots())
// }
