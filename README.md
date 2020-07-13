# Typical Go

[![Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public.](https://www.repostatus.org/badges/latest/wip.svg)](https://www.repostatus.org/#wip)
![Go-Workflow](https://github.com/typical-go/typical-go/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/typical-go/typical-go)](https://goreportcard.com/report/github.com/typical-go/typical-go)
[![codebeat badge](https://codebeat.co/badges/a8b3c7a6-c42a-480a-acb4-68ece12f36b8)](https://codebeat.co/projects/github-com-typical-go-typical-go-master)
[![codecov](https://codecov.io/gh/typical-go/typical-go/branch/master/graph/badge.svg)](https://codecov.io/gh/typical-go/typical-go)

A Build Tool (+ Framework) for Golang. <https://typical-go.github.io/>


## Install

```
$ go install github.com/typical-go/typical-go
```

## Usage

Run typical-go binary to build the build-system and run it.  
```
$ typical-go run
```

The build-system output view:
```
Typical Build

Usage:

  ./typicalw <command> [argument]

The commands are:

  test, t      Test the project
  compile, c   Compile the project
  run, r       Run the project in local environment
  release      Release the project
  clean        Clean the project
  examples, e  Test all example
  help, h      Shows a list of commands or help for one command

Use "./typicalw help <topic>" for more information about that topic
```

Check help for argument documentation
```
$ typical-go help run
```
```
NAME:
   typical-go run - Run build-tool for project in current working directory

USAGE:
   typical-go run [command options] [arguments...]

OPTIONS:
   --src value          Build-tool source (default: "tools/typical-build")
   --project-pkg value  Project package name. Same with module package in go.mod by default
   --typical-tmp value  Temporary directory location to save builds-related files (default: ".typical-tmp")
   --create:wrapper     Create wrapper script (default: false)
```

## Typical Build

Typical Build is golang program that manage build and task for current project. By default located in `tools/typical-build`

```go
package main

import (
	"log"

	"github.com/typical-go/typical-go/pkg/typgo"
)

var (
	// Descriptor of sample
	descriptor = typgo.Descriptor{
		Name:    "hello-world",
		Version: "1.0.0",

		Commands: typgo.Commands{
			&typgo.CompileCmd{
				Action: &typgo.StdCompile{},
			},
			&typgo.RunCmd{
				Action: &typgo.StdRun{},
			},
			&typgo.TestCmd{
				Action: &typgo.StdTest{},
			},
			&typgo.CleanCmd{
				Action: &typgo.StdClean{},
			},
		},
	}
)

func main() {
	if err := typgo.Run(&descriptor); err != nil {
		log.Fatal(err)
	}
}
```

## Wrapper 

The wrapper that invoke download typical-go and execute it. This is the recommendation way to use typical-go.
```
$ typical-go -create:wrapper
```

Run wrapper script
```
$ ./typicalw
```
```go 
Typical Build

Usage:

  ./typicalw <command> [argument]

The commands are:

  compile, c  Compile the project
  run, r      Run the project in local environment
  clean       Clean the project
  help, h     Shows a list of commands or help for one command

Use "./typicalw help <topic>" for more information about that topic
```

## Annotation

Typical-Go support java-like annotation (except the parameter in JSON format) for code-generation purpose.

## Dependency Injection

Typical-Go encourage dependency-injection using [dig](https://github.com/uber-go/dig) and annotation. See the [example](https://github.com/typical-go/typical-go/tree/master/examples/provide-constructor).

```go
// OpenConn open new database connection
// @ctor
func OpenConn() *sql.DB{
}
```

```go
// CloseConn close the database connection
// @dtor
func CloseConn(db *sql.DB){
}
```

## Mock

Typical-Go encourge mocking using [gomock](https://github.com/golang/mock) and annotation. See the [example](https://github.com/typical-go/typical-go/tree/master/examples/generate-mock).

```go
type(
   // Reader responsible to read
   // @mock
   Reader interface{
      Read() error
   }
)
```

## Examples

This repo contain both library, examples and wrapper source-code. The wrapper itself using Typical-Go as its build-tool which is an excellent example.

List of example:
- [x] [Hello World](https://github.com/typical-go/typical-go/tree/master/examples/hello-world)
- [x] [Configuration With Invocation](https://github.com/typical-go/typical-go/tree/master/examples/configuration-with-invocation)
- [x] [Simple Additional Task](https://github.com/typical-go/typical-go/tree/master/examples/simple-additional-task)
- [x] [Provide Constructor](https://github.com/typical-go/typical-go/tree/master/examples/provide-constructor)
- [x] [Generate Mock](https://github.com/typical-go/typical-go/tree/master/examples/generate-mock)
- [x] [Generate Docker-Compose](https://github.com/typical-go/typical-go/tree/master/examples/generate-docker-compose)
- [x] [Serve React Demo](https://github.com/typical-go/typical-go/tree/master/examples/serve-react-demo)


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
