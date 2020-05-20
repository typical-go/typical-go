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
// Descriptor of sample
var Descriptor = typgo.Descriptor{
	Name:    "hello-world",
	Version: "1.0.0",

	EntryPoint: helloworld.Main,

	Build: &typgo.StdBuild{},
}

```