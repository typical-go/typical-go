# Configuration with invocation

Example typical-go project to demonstrate how to set the configuration

Append the configurer to the project descriptor `typical/descriptor.go`
```go
var Descriptor = typcore.Descriptor{
	// ...

	App: typapp.Create(serverApp), // wrap serverApp with Typical App

	ConfigManager: typcfg.
		Create(
			serverApp, // Append configurer for the this project
		),
}

```

Example of configurer implementation
```go
// Configure the application
func (a *App) Configure() *typcore.Configuration {
	return typcore.NewConfiguration(a.ConfigName, &Config{})
}
```

Create the invocation and function with config as its parameter
```go
// EntryPoint of application
func (a *App) EntryPoint() *typdep.Invocation {
	return typdep.NewInvocation(a.start)
}

func (a *App) start(cfg *Config) error {
	fmt.Printf("Configuration With Invocation -- Serve http at %s\n", cfg.Address)
	return http.ListenAndServe(cfg.Address, a)
}

```