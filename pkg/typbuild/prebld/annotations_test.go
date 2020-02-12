package prebld_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typbuild/prebld"
)

func TestAnnotations_Contain(t *testing.T) {
	testcases := []struct {
		prebld.Annotations
		name    string
		contain bool
	}{
		{prebld.Annotations{{Name: "tag1"}, {Name: "tag2"}}, "tag1", true},
		{prebld.Annotations{{Name: "tag1"}, {Name: "tag2"}}, "tag3", false},
		{prebld.Annotations{{Name: "TAG1"}, {Name: "TAG2"}}, "tag1", true},
	}
	for _, tt := range testcases {
		require.Equal(t, tt.contain, tt.Contain(tt.name))
	}
}
