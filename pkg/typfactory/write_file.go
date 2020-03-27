package typfactory

import (
	"io"
	"os"
)

// Writer responsible to write
type Writer interface {
	Write(io.Writer) error
}

// WriteFile to write file
func WriteFile(filename string, perm os.FileMode, w Writer) (err error) {
	var f *os.File
	if f, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm); err != nil {
		return
	}
	defer f.Close()

	return w.Write(f)
}
