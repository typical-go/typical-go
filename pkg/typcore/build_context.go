package typcore

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"go.uber.org/dig"

	"github.com/typical-go/typical-go/pkg/common"
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

// Invoke function
func (b *BuildContext) Invoke(c *cli.Context, fn interface{}) (err error) {
	di := dig.New()
	if err = di.Provide(func() *cli.Context { return c }); err != nil {
		return
	}
	if b.Configuration != nil {
		if err = provide(di, b.Configuration.Provide()...); err != nil {
			return
		}
	}
	startFn := func() error {
		return di.Invoke(fn)
	}
	errs := common.NewApplication(startFn).Run()
	for _, err := range errs {
		log.Error(err.Error())
	}
	return
}

// ActionFunc to return action function that required config and object only
func (b *BuildContext) ActionFunc(fn interface{}) func(ctx *cli.Context) error {
	return func(c *cli.Context) (err error) {
		return b.Invoke(c, fn)
	}
}

func invoke(di *dig.Container, fns ...interface{}) (err error) {
	for _, fn := range fns {
		if err = di.Invoke(fn); err != nil {
			return
		}
	}
	return
}

func provide(di *dig.Container, fns ...interface{}) (err error) {
	for _, fn := range fns {
		if err = di.Provide(fn); err != nil {
			return
		}
	}
	return
}
