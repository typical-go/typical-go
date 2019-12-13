package walker_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typprebuilder/walker"
)

func TestDocTags_ParseDocTag(t *testing.T) {
	tags := walker.ParseAnnotations("[tag1] some text [tag2] then another text [tag3]")
	require.Equal(t, walker.Annotations{
		{Name: "tag1"},
		{Name: "tag2"},
		{Name: "tag3"},
	}, tags)
}

func TestDocTags_Contain(t *testing.T) {
	testcases := []struct {
		walker.Annotations
		name    string
		contain bool
	}{
		{walker.Annotations{{Name: "tag1"}, {Name: "tag2"}}, "tag1", true},
		{walker.Annotations{{Name: "tag1"}, {Name: "tag2"}}, "tag3", false},
		{walker.Annotations{{Name: "TAG1"}, {Name: "TAG2"}}, "tag1", true},
	}
	for _, tt := range testcases {
		require.Equal(t, tt.contain, tt.Contain(tt.name))
	}
}
