package common

import (
	"bufio"
	"os"
	"strings"
)

// GoMod is naive go.mod file parser
type GoMod struct {
	ModulePackage string
	GoVersion     string
}

// CreateGoMod to create new GoMod instance
func CreateGoMod(path string) (gomod *GoMod, err error) {
	var (
		f             *os.File
		modulePackage string
		goVersion     string
	)

	if f, err = os.Open(path); err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
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
	}, nil
}
