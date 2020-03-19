# Generate Mock

Example typical-go project to demonstrate how to generate mock

Put `[mock]` annotation in interface comment
```go
// Greeter responsible to greet [mock]
type Greeter interface {
	Greet() string
}
```

Run the mock task to generate mock at `mock_` + package name. 
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