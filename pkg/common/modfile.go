package common

import (
	"bufio"
	"io"
	"strings"
)

// Modfile is go.mod file detailsd
type Modfile struct {
	ProjectPackage string
	GoVersion      string
}

// ParseModfile to naively parse go.mod source
func ParseModfile(r io.Reader) *Modfile {
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

	return &Modfile{
		ProjectPackage: modulePackage,
		GoVersion:      goVersion,
	}
}
