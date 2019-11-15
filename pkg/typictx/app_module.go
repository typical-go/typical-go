package typictx

// AppModule is application module
type AppModule interface {
	Run() interface{}
}

// NewAppModule return new instance of AppModule
func NewAppModule(fn interface{}) AppModule {
	return &appModule{fn: fn}
}

type appModule struct {
	fn interface{}
}

func (a *appModule) Run() interface{} {
	return a.fn
}
