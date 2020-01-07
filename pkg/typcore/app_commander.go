package typcore

import (
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
)

// AppCommander responsible to return commands for App
type AppCommander interface {
	AppCommands(*AppContext) []*cli.Command
}

// IsAppCommander return true if object implementation of AppCLI
func IsAppCommander(obj interface{}) (ok bool) {
	_, ok = obj.(AppCommander)
	return
}

// AppContext is context of app
type AppContext struct {
	*ProjectDescriptor
}

// NewAppContext return new instance of AppContext
func NewAppContext(desc *ProjectDescriptor) *AppContext {
	return &AppContext{
		ProjectDescriptor: desc,
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
		if err = di.Provide(func() *cli.Context {
			return c
		}); err != nil {
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
	if err = provide(di, a.Constructors...); err != nil {
		return
	}
	if err = invoke(di, a.App.Prepare()...); err != nil {
		return
	}
	runner := common.Application{
		StartFn: func() error { return di.Invoke(fn) },
		StopFn:  func() error { return invoke(di, a.App.Destroy()...) },
	}
	for _, err := range runner.Run() {
		log.Error(err.Error())
	}
	return
}
