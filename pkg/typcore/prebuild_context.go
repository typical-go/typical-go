package typcore

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/typcore/walker"
	"github.com/typical-go/typical-go/pkg/typenv"
)

// PrebuildContext is context of prebuild
type PrebuildContext struct {
	context.Context
	*ProjectDescriptor
	*ProjectInfo
	Declarations []*walker.Declaration
}

// DeclFunc to handle declaration
type DeclFunc func(*walker.Declaration) error

// AnnotationFunc to handle annotation
type AnnotationFunc func(decl *walker.Declaration, ann *walker.Annotation) error

// CreatePrebuildContext to create PrebuildContext
func CreatePrebuildContext(ctx context.Context, d *ProjectDescriptor) (pc *PrebuildContext, err error) {
	var (
		projInfo     *ProjectInfo
		declarations []*walker.Declaration
	)
	if projInfo, err = ReadProject(typenv.Layout.App); err != nil {
		log.Fatal(err.Error())
	}
	if declarations, err = walker.Walk(projInfo.Files); err != nil {
		return
	}
	pc = &PrebuildContext{
		ProjectDescriptor: d,
		Declarations:      declarations,
		Context:           ctx,
		ProjectInfo:       projInfo,
	}
	return
}

// EachDecl to handle each declaration
func (c *PrebuildContext) EachDecl(fn DeclFunc) (err error) {
	for _, decl := range c.Declarations {
		if err = fn(decl); err != nil {
			return
		}
	}
	return
}

// EachAnnotation to handle each annotation
func (c *PrebuildContext) EachAnnotation(name string, declType walker.DeclType, fn AnnotationFunc) (err error) {
	return c.EachDecl(func(decl *walker.Declaration) (err error) {
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
