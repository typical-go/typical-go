# Typical Go

[![Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public.](https://www.repostatus.org/badges/latest/wip.svg)](https://www.repostatus.org/#wip)
![Go-Workflow](https://github.com/typical-go/typical-go/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/typical-go/typical-go)](https://goreportcard.com/report/github.com/typical-go/typical-go)

A Build Tool (+ Framework) for Golang. <https://typical-go.github.io/>

## Introduction

Typical-Go provides levels of abstraction for build/compile the (golang) project. The unique about Typical-Go is it use go-based descriptor file rather than DSL which is making it easier to understand and maintain.

You can use Typical-Go as:
- Framework to create custom build-tool
- Golang build-tool
- Build-Tool as a framework 

## Build-Tool As A Framework (BAAF)

Build-Tool as a framework (BAAF) is a concept where both build-tool and application utilize the same definition/settings/descriptor. We no longer see build-tool as a separate beast with the application but rather part of the same living organism. This is only possible because the app, build-tool, and descriptor speak with the same tongue.

## Descriptor File

Typically, the descriptor defined in `typical/descriptor.go` 
```go

var Descriptor = typcore.Descriptor{
	Name:        "typical-rest-server",                                       // name of the project
	Description: "Example of typical and scalable RESTful API Server for Go", // description of the project
	Version:     "0.8.25",                                                    // version of the project

	App: typapp.EntryPoint(server.Main, "server").
		Imports(
			server.Configuration(), 
			typredis.Module(),    // create and destroy redis connection
			typpostgres.Module(), // create and destroy postgres db connection
		),

	BuildTool: typbuildtool.
		BuildSequences(
			typbuildtool.StandardBuild(),
			typbuildtool.Github("typical-go", "typical-rest-server"), // Create release to Github
		).
		Utilities(
			typpostgres.Utility(), // create database, drop, migrate, seed, etc.
			typredis.Utility(),    // redis console

			// Generate dockercompose and spin up docker
			typdocker.Compose(
				typpostgres.DockerRecipeV3(),
				typredis.DockerRecipeV3(),
			),
		),
}
```

- `BuildTool` is definition of build-tool for the project. Use `./typicalw` to run the build-tool
- `App` is definition of the application. Use `./typicalw run` to run the application


## Build-Tool Wrapper

`typicalw` is your best friend. It will download, compile and run the actual build-tool for your day-to-day development.

```bash
./typicalw
```

```
NAME:
   typical-rest-server - Build-Tool

USAGE:
   build-tool [global options] command [command options] [arguments...]

VERSION:
   0.8.25

DESCRIPTION:
   Example of typical and scalable RESTful API Server for Go

COMMANDS:
   build, b      Build the binary
   test, t       Run the testing
   run, r        Run the binary
   clean, c      Clean the project from generated file during build time
   release       Release the distribution
   mock          Generate mock class
   postgres, pg  Postgres Utility
   redis         Redis utility
   docker        Docker utility
   help, h       Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

## Typical Tmp

Typical-tmp is an important folder that contains the build-tool mechanisms. By default, it is located in `.typical-tmp` and can be changed by hacking/editing the `typicalw` script.

Since the typical-go project is still undergoing development, maybe there is some breaking change and deleting typical-tmp can solve the issue since it will be healed by itself.

## Examples

- [x] [Hello World](https://github.com/typical-go/typical-go/tree/master/examples/hello-world)
- [x] [Configuration With Invocation](https://github.com/typical-go/typical-go/tree/master/examples/configuration-with-invocation)
- [x] [Simple Additional Task](https://github.com/typical-go/typical-go/tree/master/examples/simple-additional-task)
- [x] [Provide Constructor](https://github.com/typical-go/typical-go/tree/master/examples/provide-constructor)
- [x] [Generate Mock](https://github.com/typical-go/typical-go/tree/master/examples/generate-mock)
- [x] [Generate Docker-Compose](https://github.com/typical-go/typical-go/tree/master/examples/generate-docker-compose)
- [x] [Serve React Demo](https://github.com/typical-go/typical-go/tree/master/examples/serve-react-demo)



## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details




