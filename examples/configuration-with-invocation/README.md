# Configuration with invocation

Example typical-go project to demonstrate how to set the configuration

Add `ConfigManager` in `typical/descriptor.go`
```go
// Descriptor of sample
var Descriptor = typgo.Descriptor{
	// ...

	Prebuild: &typgo.ConfigManager{
		Configs: []*typgo.Configuration{
			{Name: "SERVER", Spec: &server.Config{}},
		},
	},
}
```
