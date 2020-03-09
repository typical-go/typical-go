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