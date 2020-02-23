package buildkit

import (
	"bufio"
	"io"
	"strings"
)

// GoMod is go.mod file details
type GoMod struct {
	ModulePackage string
	GoVersion     string
}

// ParseGoMod to naively parse go mod source
func ParseGoMod(r io.Reader) *GoMod {
	var (
		modulePackage string
		goVersion     string
	)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		row := scanner.Text()
		if strings.HasPrefix(row, "module") {
			modulePackage = strings.TrimSpace(row[6:])
			continue
		}
		if strings.HasPrefix(row, "go") {
			goVersion = strings.TrimSpace(row[2:])
			continue
		}
	}

	return &GoMod{
		ModulePackage: modulePackage,
		GoVersion:     goVersion,
	}
}
