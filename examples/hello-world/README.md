# Hello World

Simple hello-world application

Implement of `typcore.App`
```go
// Run app
func (*App) Run(d *typcore.Descriptor) error {
	fmt.Println("Hello World")
	return nil
}
```

Setup the typical descriptor
```go
// Descriptor of sample
var Descriptor = typcore.Descriptor{
	Name:    "hello-world",
	Version: "0.0.1",

	App: helloworld.New(), // the application

	BuildTool: typbuildtool.
		Create(
			typbuildtool.StandardBuild(), // standard build module
		),
}

```