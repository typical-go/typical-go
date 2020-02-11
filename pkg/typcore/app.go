package typcore

// App is interface of app
type App interface {
	Run(*AppContext) error
}

// AppContext is context of app
type AppContext struct {
	*Descriptor
}
