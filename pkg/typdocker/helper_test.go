package typdocker_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typdocker"
)

func TestMajor(t *testing.T) {
	testcases := []struct {
		version string
		major   string
	}{
		{"3", "3"},
		{"2", "2"},
		{"2.0.1", "2"},
		{"3.0", "3"},
	}

	for _, tt := range testcases {
		require.Equal(t, tt.major, typdocker.Major(tt.version))
	}

}
