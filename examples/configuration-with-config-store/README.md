# Configuration with invocation

Example typical-go project to demonstrate how to set the configuration

Append the configurer to the project descriptor `typical/descriptor.go`
```go
var Descriptor = typcore.Descriptor{
	// ...

	App: typapp.New(serverApp), // wrap serverApp with Typical App

	Configuration: typcfg.New().
		AppendConfigurer(
			serverApp, // Append configurer for the this project
		),
}

```

Example of configurer implementation
```go
func (m *Module) Configure(loader typcfg.Loader) *typcfg.Detail {
	return &typcfg.Detail{
		Prefix: m.Prefix,
		Spec:   &config.Config{},
		Constructor: typdep.NewConstructor(
			func() (cfg config.Config, err error) {
				err = loader.Load(m.Prefix, &cfg)
				return
			}),
	}
}
```

Create the invocation and function with config as its parameter
```go
// EntryPoint of application
func (a *App) EntryPoint() *typdep.Invocation {
	return typdep.NewInvocation(a.start)
}

func (a *App) start(cfg config.Config) error {
	return http.ListenAndServe(cfg.Address, a)
}
```