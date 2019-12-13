package walker_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typprebuilder/walker"
)

func TestDocTags_ParseDocTag(t *testing.T) {
	tags := walker.ParseNotations("[tag1] some text [tag2] then another text [tag3]")
	require.Equal(t, walker.Notations{"tag1", "tag2", "tag3"}, tags)
}

func TestDocTags_Contain(t *testing.T) {
	testcases := []struct {
		walker.Notations
		name    string
		contain bool
	}{
		{[]string{"tag1", "tag2"}, "tag1", true},
		{[]string{"tag1", "tag2"}, "tag3", false},
		{[]string{"TAG1", "TAG2"}, "tag1", true},
	}
	for _, tt := range testcases {
		require.Equal(t, tt.contain, tt.Contain(tt.name))
	}
}
