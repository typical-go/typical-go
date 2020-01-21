package walker_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcore/walker"
)

func TestAnnotations_Contain(t *testing.T) {
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
