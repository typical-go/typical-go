package typcore

import (
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/common"

	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
)

// NewContext return new instance of Context
func NewContext(d *ProjectDescriptor) *Context {
	return &Context{
		ProjectDescriptor: d,
	}
}

// Context of application
type Context struct {
	*ProjectDescriptor
}

// Action to return action function that required config and object only
func (c *Context) Action(obj, fn interface{}) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) (err error) {
		di := dig.New()
		if err = di.Provide(func() *cli.Context {
			return ctx
		}); err != nil {
			return
		}
		if c.Configuration != nil {
			if err = provide(di, c.Configuration.Provide()...); err != nil {
				return
			}
		}
		if err := common.NewApplication(func() error {
			return di.Invoke(fn)
		}).Run(); err != nil {
			log.Error(err.Error())
		}
		return
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
