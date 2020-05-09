# Hello World

Simple hello-world application

Create function with `*typgo.Descriptor` as function parameter
```go
func Main(d *typgo.Descriptor) (err error) {
	fmt.Println("Hello World")
	return
}
```

Setup the typical descriptor
```go
var Descriptor = typgo.Descriptor{
	Name:    "hello-world",
	Version: "1.0.0",

	App: typgo.NewApp(helloworld.Main), // the application

	BuildTool: typgo.
		BuildSequences(
			typgo.StandardBuild(), // standard build module
		),
}
```