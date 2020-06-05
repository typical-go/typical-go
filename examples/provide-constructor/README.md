# Provide Constructor

Example typical-go project to demonstrate how to provide constructor to dependency injection

Put `DependencyInject` in prebuild descriptor
```go
// Descriptor of sample
var Descriptor = typgo.Descriptor{
	// ...
	Prebuild: &typgo.DependencyInjection{},
}

```

Put `@ctor` annotation in constructor function comment
```go
// HelloWorld text
// @ctor
func HelloWorld() string {
	return "Hello World"
}

// HelloTypical text
// @ctor {"name": "typical"}
func HelloTypical() string {
	return "Hello Typical"
}
```

Put `@dtor` annotation to function that called after application close
```go
// Close the application
// @dtor
func Close() {
	fmt.Println("close the app")
}
```

`precond_DO_NOT_EDIT.go` will be generated in `typical/` folder to provide the constructor
```go
func init() {
	typgo.Provide(
		typgo.NewConstructor("", helloworld.HelloWorld),
		typgo.NewConstructor("typical", helloworld.HelloTypical),
	)
}
```



