# Provide Constructor

Example typical-go project to demonstrate how to provide constructor to dependency injection

Put `[constructor]` annotation in constructor function comment
```go
// NewGreeter return new insteance of Greeter [constructor]
func NewGreeter() *Greeter {
	return &Greeter{}
}
```

`init_constructor_do_not_edit.go` will be generated in `typical/` folder
```go
func init() {
	typapp.AppendConstructor(
		typdep.NewConstructor(helloworld.NewGreeter),
	)
}
```

Create the invocation and function with config as its parameter
```go
// EntryPoint of application
func (a *App) EntryPoint() *typdep.Invocation {
	return typdep.NewInvocation(start)
}

func start(greeter *Greeter) {
	greeter.Greet()
}
```