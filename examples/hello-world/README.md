# Hello World

Simple hello-world application

Create function with `*typcore.Descriptor` as function parameter
```go
func Main(d *typcore.Descriptor) (err error) {
	fmt.Println("Hello World")
	return
}
```

Setup the typical descriptor
```go
var Descriptor = typcore.Descriptor{
	Name:    "hello-world",
	Version: "1.0.0",

	App: typcore.NewApp(helloworld.Main), // the application

	BuildTool: typcore.
		BuildSequences(
			typcore.StandardBuild(), // standard build module
		),
}
```