package typgo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestDescriptor_ValidateName(t *testing.T) {

	testcases := []struct {
		testname string
		name     string
		expected bool
	}{
		{
			name:     "asdf",
			expected: true,
		},
		{
			name:     "Asdf",
			expected: true,
		},
		{
			name:     "As_df",
			expected: true,
		},
		{
			name:     "as-df",
			expected: true,
		},
		{
			name:     "Asdf!",
			expected: false,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testname, func(t *testing.T) {
			require.Equal(t, tt.expected, typgo.ValidateName(tt.name))
		})
	}

}

func TestDecriptor_Validate_ReturnError(t *testing.T) {
	testcases := []struct {
		testname string
		*typgo.Descriptor
		expectedErr string
	}{
		{
			Descriptor: validDescriptor,
		},
		{
			Descriptor: &typgo.Descriptor{
				Name: "Typical Go",
				BuildSequences: []interface{}{
					struct{}{},
				},
			},
			expectedErr: "Descriptor: bad name",
		},
	}
	for i, tt := range testcases {
		t.Run(tt.testname, func(t *testing.T) {
			err := tt.Validate()
			if tt.expectedErr == "" {
				require.NoError(t, err, i)
			} else {
				require.EqualError(t, err, tt.expectedErr, i)
			}
		})
	}
}

var (
	validDescriptor = &typgo.Descriptor{
		Name: "some-name",
		BuildSequences: []interface{}{
			struct{}{},
		},
	}
)
