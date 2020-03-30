# Typical Go

[![Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public.](https://www.repostatus.org/badges/latest/wip.svg)](https://www.repostatus.org/#wip)
![Go-Workflow](https://github.com/typical-go/typical-go/workflows/Go/badge.svg)

A Build Tool (+ Framework) for Golang. <https://typical-go.github.io/>

## Descriptor File

Define descriptor in `typical/descriptor.go` folder
```go
var Descriptor = typcore.Descriptor{
    Name:    "configuration-with-invocation",
    Version: "1.0.0",

    App: typapp.EntryPoint(server.Main), 

    BuildTool: typbuildtool.
        BuildSequences(
            typbuildtool.StandardBuild(),
        ),

    ConfigManager: typcfg.
        Configures(
            server.Configuration(), 
        ),
}
```

- `BuildTool` is definition of build-tool for the project. Use `./typicalw` to run the build
- `App` is definition of the application. Use `./typicalw run` to run the application
- `ConfigManager` is configuration for the project. Note: This is subject to change on next version.


## Build-Tool Wrapper

`typicalw` is your best friend. It will download, compile and run the actual build-tool for your daily development task.

```bash
./typicalw
```

```
NAME:
   configuration-with-invocation - Build-Tool

USAGE:
   build-tool [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
   build, b  Build the binary
   test, t   Run the testing
   run, r    Run the binary
   clean, c  Clean the project from generated file during build time
   release   Release the distribution
   mock      Generate mock class
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false) 
```

## Typical Tmp

Typical-tmp is an important folder that contains the build tool mechanism. By default, it is located in `.typical-tmp` and can be changed by hacking/editing the `typicalw` script.

Since the typical-go project is still undergoing development, maybe there is some breaking change and deleting typical-tmp can solve the issue since it will be healed by itself.

## Examples

- [x] [Hello World](https://github.com/typical-go/typical-go/tree/master/examples/hello-world)
- [x] [Get Config From Descriptor](https://github.com/typical-go/typical-go/tree/master/examples/get-config-from-descriptor)
- [x] [Configuration With Invocation](https://github.com/typical-go/typical-go/tree/master/examples/configuration-with-invocation)
- [x] [Simple Additional Task](https://github.com/typical-go/typical-go/tree/master/examples/simple-additional-task)
- [x] [Provide Constructor](https://github.com/typical-go/typical-go/tree/master/examples/provide-constructor)
- [x] [Generate Mock](https://github.com/typical-go/typical-go/tree/master/examples/generate-mock)
- [ ] [Generate Readme](https://github.com/typical-go/typical-go/tree/master/examples/generate-readme)
- [x] [Generate Docker-Compose](https://github.com/typical-go/typical-go/tree/master/examples/generate-docker-compose)
- [x] [Serve React Demo](https://github.com/typical-go/typical-go/tree/master/examples/serve-react-demo)



## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details




