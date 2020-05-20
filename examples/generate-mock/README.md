# Generate Mock

Example typical-go project to demonstrate how to generate mock

Put `@mock` annotation in interface comment
```go
// Greeter responsible to greet 
// @mock
type Greeter interface {
	Greet() string
}
```

Add typmock utility to descriptor
```go
var Descriptor = typgo.Descriptor{
	// ...

	Utility: typmock.Utility(),
}
```

Generate mock
```bash
./typicalw mock
```

Use the mock class
```go
greeter := mock_helloworld.NewMockGreeter(ctrl)
greeter.EXPECT().Greet().Return("some-word")
```

Run the test
```bash
./typicalw test
```

## Learn More

Typical-Go using [gomock](https://github.com/golang/mock) in behind