package typvar

import "fmt"

type (
	// Build data
	Build struct {
		Binary   string
		Source   string
		Checksum string
	}
)

// GetBuild to get build data
func GetBuild() *Build {
	return &Build{
		Checksum: fmt.Sprintf("%s/checksum", TypicalTmp),
		Binary:   fmt.Sprintf("%s/bin/build-tool", TypicalTmp),
		Source:   fmt.Sprintf("%s/build-tool/main.go", TypicalTmp),
	}
}
