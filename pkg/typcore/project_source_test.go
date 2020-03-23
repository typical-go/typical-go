package typcore_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/typicalgo"
)

func TestRetrieveProjectSources(t *testing.T) {
	testcases := []struct {
		*typcore.Descriptor
		expected      []string
		expectedError string
	}{
		{
			Descriptor:    &typcore.Descriptor{App: typicalgo.New()},
			expectedError: "ProjectSource 'typicalgo' is not exist",
		},
		{
			Descriptor:    &typcore.Descriptor{App: typapp.Create(typicalgo.New())},
			expectedError: "ProjectSource 'typicalgo' is not exist",
		},
	}

	for _, tt := range testcases {
		sources, err := typcore.ProjectSources(tt.Descriptor)
		if tt.expectedError == "" {
			require.NoError(t, err)
			require.Equal(t, tt.expected, sources)
		} else {
			require.EqualError(t, err, tt.expectedError)
		}

	}
}
