# Provide Constructor

Example typical-go project to demonstrate how to provide constructor to dependency injection


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




