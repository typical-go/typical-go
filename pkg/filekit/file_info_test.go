package filekit_test

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/filekit"
)

func TestFileInfo(t *testing.T) {
	info := filekit.FileInfo{
		NameField:    "some-name",
		SizeField:    999,
		ModeField:    777,
		ModTimeField: time.Time{},
		IsDirField:   true,
		SysField:     struct{}{},
	}
	require.Equal(t, "some-name", info.Name())
	require.Equal(t, int64(999), info.Size())
	require.Equal(t, os.FileMode(777), info.Mode())
	require.Equal(t, time.Time{}, info.ModTime())
	require.Equal(t, true, info.IsDir())
	require.Equal(t, struct{}{}, info.Sys())
}
