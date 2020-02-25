package typbuild

import (
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typbuild/prebld"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Context of build
type Context struct {
	*typcore.TypicalContext
	Declarations []*prebld.Declaration
}

// DeclFunc to handle declaration
type DeclFunc func(*prebld.Declaration) error

// AnnotationFunc to handle annotation
type AnnotationFunc func(decl *prebld.Declaration, ann *prebld.Annotation) error

// EachDecl to handle each declaration
func (c *Context) EachDecl(fn DeclFunc) (err error) {
	for _, decl := range c.Declarations {
		if err = fn(decl); err != nil {
			return
		}
	}
	return
}

// EachAnnotation to handle each annotation
func (c *Context) EachAnnotation(name string, declType prebld.DeclType, fn AnnotationFunc) (err error) {
	return c.EachDecl(func(decl *prebld.Declaration) (err error) {
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
