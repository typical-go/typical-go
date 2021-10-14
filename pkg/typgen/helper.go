package typgen

import (
	"path/filepath"
	"strings"
)

const (
	DefaultParent = "internal/generated"
)

func CreateTargetDir(d *Annotation, suffix string) string {
	dir := filepath.Dir(d.Path())
	if dir == "." {
		return DefaultParent
	}
	dir = strings.ReplaceAll(dir, "internal/", "")
	return DefaultParent + "/" + dir + "_" + suffix
}
