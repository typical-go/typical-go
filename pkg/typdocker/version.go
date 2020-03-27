package typdocker

import (
	"strings"
)

// Major version from docker-composer version
func Major(version string) string {
	i := strings.IndexRune(version, '.')
	if i < 0 {
		return version
	}

	return version[:i]
}
