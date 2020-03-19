# Configuration with Config Store

Example typical-go project to demonstrate how to get the configuration from descriptor

Set configuration for the application or module by implement `func Configure() *typcore.Configuration`
```go
// Configure the application
func (a *App) Configure() *typcore.Configuration {
	return typcore.NewConfiguration(a.ConfigName, &Config{})
}
```

Register it to descriptor in `WithConfigurer(...)` method
```go
var Descriptor = typcore.Descriptor{
    // ... 

	ConfigManager: typcfg.
		Create(
			serverApp, // register serverApp configurer
		),
}
```

Use `RetrieveConfig()` to get the config
```go
// Run server
func (a *App) Run(d *typcore.Descriptor) (err error) {
	var spec interface{}
	if spec, err = d.RetrieveConfig(a.ConfigName); err != nil {
		return
	}

	// type assertion to Config type
	cfg := spec.(*Config)

	fmt.Printf("Configuration With Config Store -- Serve http at %s\n", cfg.Address)
	return http.ListenAndServe(cfg.Address, a)
}
```