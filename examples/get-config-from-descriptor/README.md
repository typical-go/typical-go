# Configuration with Config Store

Example typical-go project to demonstrate how to get the configuration from descriptor

Register configuration to descriptor in `WithConfigurer(...)` method
```go
// Descriptor of sample
var Descriptor = typcore.Descriptor{
	// ...

	ConfigManager: typcfg.
		Create(
			typcfg.NewConfiguration(server.ConfigName, &server.Config{}), // register serverApp configurer
		),
}
```

Use `RetrieveConfig()` to get the config
```go
// Main function to run the server
func Main(d *typcore.Descriptor) (err error) {
	var spec interface{}
	if spec, err = d.RetrieveConfig(ConfigName); err != nil {
		return
	}

	// type assertion to Config type
	cfg := spec.(*Config)

	fmt.Printf("Get Config From Descriptor -- Serve http at %s\n", cfg.Address)
	return http.ListenAndServe(cfg.Address, &handler{})
}
```