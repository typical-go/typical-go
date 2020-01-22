package typcore

import (
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/common"

	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
)

// BuildContext is context for build tool
type BuildContext struct {
	*ProjectDescriptor
}

// NewBuildContext to return new instance of BuildContext
func NewBuildContext(desc *ProjectDescriptor) *BuildContext {
	return &BuildContext{
		ProjectDescriptor: desc,
	}
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
