# Typical Go

[![Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public.](https://www.repostatus.org/badges/latest/wip.svg)](https://www.repostatus.org/#wip)
![Go-Workflow](https://github.com/typical-go/typical-go/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/typical-go/typical-go)](https://goreportcard.com/report/github.com/typical-go/typical-go)
[![codebeat badge](https://codebeat.co/badges/a8b3c7a6-c42a-480a-acb4-68ece12f36b8)](https://codebeat.co/projects/github-com-typical-go-typical-go-master)

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

## Application

`App` in descriptor define the application. `./typicalw run` run the application based on this.

`typapp` package is common golang application geared with dependency-injection and configuration. 
- `EntryPoint` contain main function and source folder. 
- `Imports` to put configurations, constructor, destructor or preparation into the application

You can make your own application implementation by implmenent `typcore.App` interface

## Build Tool

`BuildTool` in descriptor define the build-tool. `./typicalw` run the build-tool based on this.

`typbuildtool` package is common build-tool with build-sequence and utilities.
- `BuildSequence` is sequence of build process (check [Build Life-Cycle](#build-life-cycle) section)
- `Utilities` custom task for development

## Build Life-Cycle

Each build-sequence contain either precondition, run, test, release or publish. 

```
        +--------------------+       
        |    Precondition    |       
        +--------------------+       
            /            \           
           /              \          
          /                \         
         /         +----------------+
        /          |      Test      |
+-------------+    +----------------+
|     Run     |             |        
+-------------+             |        
                   +----------------+
                   |     Release    |
                   +----------------+
                            |        
                            |        
                   +----------------+
                   |     Publish    |
                   +----------------+
```

| # | Phase | Description | Command | 
|---|-------|-------------|---------|
| 1 | Precondition | Setup the project; most likely generate file that required for application | `./typicalw` |
| 2 | Test |  Test the project | `./typicalw test` |
| 3 | Run | Run the project for local environment | `./typicalw run` |
| 4 | Release | Execute before publish the project | n/a |
| 5 | Publish | Publish the project | `./typicalw publish`  |



## Wrapper

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
   test, t          Test the project
   run, r           Run the project in local environment
   publish, p       Publish the project
   clean, c         Clean the project
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




