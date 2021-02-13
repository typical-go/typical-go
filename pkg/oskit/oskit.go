package oskit

import (
	"os"
)

// MkdirAll same with os.Mkdirall and return its remove function
func MkdirAll(path string) func() {
	os.MkdirAll(path, 0777)
	return func() {
		os.RemoveAll(path)
	}
}
