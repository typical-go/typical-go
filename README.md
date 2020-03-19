# Typical Go

![stability-unstable](https://img.shields.io/badge/stability-unstable-yellow.svg)

A Framework (+ Build Tool) for Golang. <https://typical-go.github.io/>

## About

### Descriptor File

Create descriptor in `typical` folder
```go
var Descriptor = typcore.Descriptor{
	Name:    "configuration-with-invocation",
	Version: "0.0.1",

    // The Application
	App: typapp.
		Create(serverApp), 

    // The Build Tool
	BuildTool: typbuildtool.
		Create(
			typbuildtool.StandardBuild(),
		),

    // The Configuration Manager
	ConfigManager: typcfg.
		Create(
			serverApp, 
		),
}
```

### Build-Tool Wrapper

`.typicalw` is your best friend. It will download, compile and run the actual build-tool for your daily development task.


## Examples

- [Hello World](https://github.com/typical-go/typical-go/tree/master/examples/hello-world)
- [Get Config From Descriptor](https://github.com/typical-go/typical-go/tree/master/examples/get-config-from-descriptor)
- [Configuration With Invocation](https://github.com/typical-go/typical-go/tree/master/examples/configuration-with-invocation)
- [Simple Additional Task](https://github.com/typical-go/typical-go/tree/master/examples/simple-additional-task)
- [Provide Constructor](https://github.com/typical-go/typical-go/tree/master/examples/provide-constructor)
- [Serve React Demo](https://github.com/typical-go/typical-go/tree/master/examples/serve-react-demo)



## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details




