# Provide Constructor

Example typical-go project to demonstrate how to provide constructor to dependency injection

Put `@constructor` annotation in constructor function comment
```go
// HelloWorld text
// @constructor
func HelloWorld() string {
	return "Hello World"
}

// HelloTypical text
// @constructor {"name": "typical"}
func HelloTypical() string {
	return "Hello Typical"
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



