package typast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typast"
)

func TestAnnotations_Contain(t *testing.T) {
	testcases := []struct {
		typast.Annotations
		name    string
		contain bool
	}{
		{typast.Annotations{{Name: "tag1"}, {Name: "tag2"}}, "tag1", true},
		{typast.Annotations{{Name: "tag1"}, {Name: "tag2"}}, "tag3", false},
		{typast.Annotations{{Name: "TAG1"}, {Name: "TAG2"}}, "tag1", true},
	}
	for _, tt := range testcases {
		require.Equal(t, tt.contain, tt.Contain(tt.name))
	}
}
