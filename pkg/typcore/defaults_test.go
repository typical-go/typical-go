package typcore_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/typicalgo"
)

func TestDefaultProjectSources(t *testing.T) {
	testcases := []struct {
		*typcore.Descriptor
		expected []string
	}{
		{
			Descriptor: &typcore.Descriptor{App: typicalgo.New()},
			expected:   []string{"typicalgo"},
		},
		{
			Descriptor: &typcore.Descriptor{App: typapp.New(typicalgo.New())},
			expected:   []string{"typicalgo"},
		},
	}

	for _, tt := range testcases {
		require.Equal(t, tt.expected, typcore.DefaultProjectSources(tt.Descriptor))
	}
}
