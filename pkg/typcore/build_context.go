package typcore

import (
	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/typcore/walker"
)

// BuildContext is context of prebuild
type BuildContext struct {
	*Descriptor
	*ProjectInfo
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
