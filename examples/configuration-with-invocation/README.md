# Hello World With Config

Example typical-go project to demonstrate how to set the configuration

Append the configurer to the project descriptor `typical/descriptor.go`
```go
var Descriptor = typcore.Descriptor{
    // ...

	Configuration: typcfg.New().
		AppendConfigurer(
			hello, // Append configurer for the this project
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