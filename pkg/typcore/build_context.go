package typcore

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"go.uber.org/dig"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcore/walker"
	"github.com/typical-go/typical-go/pkg/typenv"
)

// BuildContext is context of prebuild
type BuildContext struct {
	*ProjectDescriptor
	*ProjectInfo
	Declarations []*walker.Declaration
}

// DeclFunc to handle declaration
type DeclFunc func(*walker.Declaration) error

// AnnotationFunc to handle annotation
type AnnotationFunc func(decl *walker.Declaration, ann *walker.Annotation) error

// CreateBuildContext to create PrebuildContext
func CreateBuildContext(d *ProjectDescriptor) (pc *BuildContext, err error) {
	var (
		projInfo     *ProjectInfo
		declarations []*walker.Declaration
	)
	if projInfo, err = ReadProject(typenv.Layout.App); err != nil {
		return
	}
	if declarations, err = walker.Walk(projInfo.Files); err != nil {
		return
	}
	return &BuildContext{
		ProjectDescriptor: d,
		Declarations:      declarations,
		ProjectInfo:       projInfo,
	}, nil
}

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
