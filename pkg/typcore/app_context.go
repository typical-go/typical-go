package typcore

import (
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
)

// AppContext is context of app
type AppContext struct {
	*Descriptor
}

// NewAppContext return new instance of AppContext
func NewAppContext(desc *Descriptor) *AppContext {
	return &AppContext{
		Descriptor: desc,
	}
}

// ActionFunc to return ActionFunc to invoke function fn
func (a *AppContext) ActionFunc(fn interface{}) func(*cli.Context) error {
	return func(c *cli.Context) (err error) {
		return a.Invoke(c, fn)
	}
}

// Invoke function with Dependency Injection
func (a *AppContext) Invoke(c *cli.Context, fn interface{}) (err error) {
	di := dig.New()
	if c != nil {
		if err = di.Provide(func() *cli.Context { return c }); err != nil {
			return
		}
	}
	if a.Configuration != nil {
		if err = provide(di, a.Configuration.Provide()...); err != nil {
			return
		}
	}
	if err = provide(di, a.App.Provide()...); err != nil {
		return
	}
	if err = provide(di, a.Constructors()...); err != nil {
		return
	}
	if err = invoke(di, a.App.Prepare()...); err != nil {
		return
	}
	startFn := func() error {
		return di.Invoke(fn)
	}
	stopFn := func() error {
		return invoke(di, a.App.Destroy()...)
	}
	errs := common.NewApplication(startFn).
		WithStopFn(stopFn).
		Run()
	for _, err := range errs {
		log.Error(err.Error())
	}
	return
}
