package oskit

import (
	"io"
	"os"
)

// Stdout standard output
var Stdout io.Writer = os.Stdout

// PatchStdout patch stdout
func PatchStdout(w io.Writer) func() {
	Stdout = w
	return func() {
		Stdout = os.Stdout
	}
}

// MkdirAll same with os.Mkdirall and return its remove function
func MkdirAll(path string) func() {
	os.MkdirAll(path, 0777)
	return func() {
		os.RemoveAll(path)
	}
}
