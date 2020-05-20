# Configuration with invocation

Example typical-go project to demonstrate how to set the configuration

Append the configurer to the project descriptor `typical/descriptor.go`
```go
// Descriptor of sample
var Descriptor = typgo.Descriptor{
	// ...

	Configurer: server.Configuration(),
}
```

For multiple configuration
```go
// Descriptor of sample
var Descriptor = typgo.Descriptor{
	// ...

	Configurer: typgo.Configurers{
		server.Configuration(),
		// More configurer...
	}
}
```