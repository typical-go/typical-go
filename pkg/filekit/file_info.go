package filekit

import (
	"os"
	"time"
)

type (
	// FileInfo os.FileInfo implementation
	FileInfo struct {
		NameField    string
		SizeField    int64
		ModeField    os.FileMode
		ModTimeField time.Time
		IsDirField   bool
		SysField     interface{}
	}
)

var _ os.FileInfo = (*FileInfo)(nil)

// Name is base name of the file
func (f *FileInfo) Name() string { return f.NameField }

// Size is length in bytes for regular files; system-dependent for others
func (f *FileInfo) Size() int64 { return f.SizeField }

// Mode is file mode bits
func (f *FileInfo) Mode() os.FileMode { return f.ModeField }

// ModTime is modification time
func (f *FileInfo) ModTime() time.Time { return f.ModTimeField }

// IsDir is abbreviation for Mode().IsDir()
func (f *FileInfo) IsDir() bool { return f.IsDirField }

// Sys is underlying data source (can return nil)
func (f *FileInfo) Sys() interface{} { return f.SysField }
