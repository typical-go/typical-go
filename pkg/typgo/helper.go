package typgo

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typvar"
)

// IsTag return true if annotation same type and tag names
func IsTag(a *typast.Annot, typ typast.DeclType, tagNames ...string) bool {
	if a.Decl.Type == typ {
		for _, tagName := range tagNames {
			if strings.EqualFold(tagName, a.TagName) {
				return true
			}
		}
	}
	return false
}

// IsFuncTag return true if annotation function has tag names
func IsFuncTag(a *typast.Annot, tagNames ...string) bool {
	return IsTag(a, typast.Function, tagNames...)
}

// ExcludeMessage return true is message mean to be exclude
func ExcludeMessage(msg string) bool {
	msg = strings.ToLower(msg)
	for _, prefix := range typvar.ExclMsgPrefix {
		if strings.HasPrefix(msg, strings.ToLower(prefix)) {
			return true
		}
	}
	return false
}

// WalkLayout return dirs and files
func WalkLayout(layouts []string) (dirs, files []string) {
	for _, layout := range layouts {
		filepath.Walk(layout, func(path string, info os.FileInfo, err error) error {
			if info == nil {
				return nil
			}

			if info.IsDir() {
				dirs = append(dirs, path)
				return nil
			}

			if isGoSource(path) {
				files = append(files, path)
			}
			return nil
		})
	}
	return
}

func isGoSource(path string) bool {
	return strings.HasSuffix(path, ".go") &&
		!strings.HasSuffix(path, "_test.go")
}
