package common

import (
	"os"
)

// IsFileExist reports whether the named file or directory exists.
func IsFileExist(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
