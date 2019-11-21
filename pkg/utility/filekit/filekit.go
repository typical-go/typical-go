package filekit

import (
	"io"
	"os"
)

// IsExist reports whether the named file or directory exists.
func IsExist(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// Write to apply write function
func Write(target string, w interface{ Write(io.Writer) error }) (err error) {
	var f *os.File
	if f, err = os.Create(target); err != nil {
		return
	}
	defer f.Close()
	return w.Write(f)
}
