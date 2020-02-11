package typcore

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/typcore/walker"
)

// Build is interface of build
type Build interface {
	Run(*BuildContext) error
}

// BuildContext is context of prebuild
type BuildContext struct {
	*Descriptor
	Dirs         []string
	Files        []string
	Declarations []*walker.Declaration
}

// DeclFunc to handle declaration
type DeclFunc func(*walker.Declaration) error

// AnnotationFunc to handle annotation
type AnnotationFunc func(decl *walker.Declaration, ann *walker.Annotation) error

// EachDecl to handle each declaration
func (b *BuildContext) EachDecl(fn DeclFunc) (err error) {
	for _, decl := range b.Declarations {
		if err = fn(decl); err != nil {
			return
		}
	}
	return
}

// EachAnnotation to handle each annotation
func (b *BuildContext) EachAnnotation(name string, declType walker.DeclType, fn AnnotationFunc) (err error) {
	return b.EachDecl(func(decl *walker.Declaration) (err error) {
		annotation := decl.Annotations.Get(name)
		if annotation != nil {
			if decl.Type == declType {
				if err = fn(decl, annotation); err != nil {
					return
				}
			} else {
				log.Warnf("[%s] has no effect to %s:%s", name, declType, decl.SourceName)
			}
		}
		return
	})
}

// Walk function
func (b *BuildContext) addFile(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		b.Dirs = append(b.Dirs, path)
	} else if isWalkTarget(path) {
		b.Files = append(b.Files, path)
	}
	return nil
}

func isWalkTarget(filename string) bool {
	return strings.HasSuffix(filename, ".go") &&
		!strings.HasSuffix(filename, "_test.go")
}
