package typast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typast"
)

func TestParseAnnotation(t *testing.T) {
	testcases := []struct {
		raw string
		*typast.Annot
	}{
		{`no-bracket`, nil},
		{`[autowire]`, typast.NewAnnot("autowire")},
		{`[mock(pkg=mock2)]`, typast.NewAnnot("mock").Put("pkg", "mock2")},
		{`[mock(pkg="mock2")]`, typast.NewAnnot("mock").Put("pkg", "mock2")},
		{`[mock(pkg="mock2"]`, typast.NewAnnot("mock")},
		{`[mock(pkg)]`, typast.NewAnnot("mock").Put("pkg", "")},
		{
			`[noname(key1="value1" key2="value2")]`,
			typast.NewAnnot("noname").Put("key1", "value1").Put("key2", "value2"),
		},
		{
			`[noname(key1=value1 key2=value2)]`,
			typast.NewAnnot("noname").Put("key1", "value1").Put("key2", "value2"),
		},
		{
			`[noname(key1=value1 key2="value2")]`,
			typast.NewAnnot("noname").Put("key1", "value1").Put("key2", "value2"),
		},
		{
			`[noname(key1=value1 key2 key3=value3)]`,
			typast.NewAnnot("noname").Put("key1", "value1").Put("key2", "").Put("key3", "value3"),
		},
		{
			`[noname(key1= key2 key3="")]`,
			typast.NewAnnot("noname").Put("key1", "").Put("key2", "").Put("key3", ""),
		},
		{
			`[noname(key1="" key2 key3=)]`,
			typast.NewAnnot("noname").Put("key1", "").Put("key2", "").Put("key3", ""),
		},
		{
			`[noname(key1="" key2 key3)]`,
			typast.NewAnnot("noname").Put("key1", "").Put("key2", "").Put("key3", ""),
		},
		{
			`[noname(key1="" key2 key3 key4=value4)]`,
			typast.NewAnnot("noname").Put("key1", "").Put("key2", "").Put("key3", "").Put("key4", "value4"),
		},
	}
	for i, tt := range testcases {
		require.Equal(t, tt.Annot, typast.ParseAnnotation(tt.raw), i)
	}
}

func TestParseAnnotations(t *testing.T) {
	raw := "[tag1] some text [tag2] then another text [tag3(key1=value1)]"
	require.Equal(t, []*typast.Annot{
		typast.NewAnnot("tag1"),
		typast.NewAnnot("tag2"),
		typast.NewAnnot("tag3").Put("key1", "value1"),
	}, typast.ParseAnnots(raw))
}
