# Configuration with invocation

Example typical-go project to demonstrate how to set the configuration

Append the configurer to the project descriptor `typical/descriptor.go`
```go
// Descriptor of sample
var Descriptor = typgo.Descriptor{
	// ...

	App: typapp.EntryPoint(server.Main, "server").
		Imports(
			server.Configuration(), // Append configurer for the this project
		),
}
```
