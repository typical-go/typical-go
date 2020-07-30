# Typical Go

[![Release](https://img.shields.io/github/release/typical-go/typical-go/all.svg)](https://github.com/typical-go/typical-go/releases/latest)
[![Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public.](https://www.repostatus.org/badges/latest/wip.svg)](https://www.repostatus.org/#wip)
![Go-Workflow](https://github.com/typical-go/typical-go/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/typical-go/typical-go)](https://goreportcard.com/report/github.com/typical-go/typical-go)
[![codebeat badge](https://codebeat.co/badges/a8b3c7a6-c42a-480a-acb4-68ece12f36b8)](https://codebeat.co/projects/github-com-typical-go-typical-go-master)
[![BCH compliance](https://bettercodehub.com/edge/badge/typical-go/typical-go?branch=master)](https://bettercodehub.com/)
[![codecov](https://codecov.io/gh/typical-go/typical-go/branch/master/graph/badge.svg)](https://codecov.io/gh/typical-go/typical-go)

Build Automation For Golang
- Alternative for [GNU Make](https://www.gnu.org/software/make/manual/make.html) (a.k.a makefile)
- Framework-based Build Tool (No DSL)
- Supporting Java-like annotation for code generation

## Install

```
$ go install github.com/typical-go/typical-go
```

## Usage

Run build-tool for project in working directory
```
$ typical-go run
```
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

## Wrapper 

The wrapper that invoke download typical-go and execute it. This is the recommendation way to use typical-go.
```
$ ./typicalw

```

To create new wrapper
```
$ typical-go -create:wrapper
```

## Typical Build

Typical Build is a golang program that manage builds and tasks for the project. By default , the source located in `tools/typical-build`

```go
package main

import (
   "fmt"

   "github.com/typical-go/typical-go/pkg/typgo"
)

var descriptor = typgo.Descriptor{
   ProjectName:    "custom-command",
   ProjectVersion: "1.0.0",

   Cmds: []typgo.Cmd{

      // compile
      &typgo.CompileCmd{
         Action: &typgo.StdCompile{},
      },

      // run
      &typgo.RunCmd{
         Before: typgo.BuildSysRuns{"compile"},
         Action: &typgo.StdRun{},
      },

      // clean
      &typgo.CleanCmd{
         Action: &typgo.StdClean{},
      },

      // ping
      &typgo.Command{
         Name: "ping",
         Action: typgo.NewAction(func(c *typgo.Context) error {
            fmt.Println("pong")
            return nil
         }),
      },
   },
}

func main() {
   typgo.Start(&descriptor)
}
```

## Annotation

Typical-Go support java-like annotation (except the parameter in [StructTag](https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go) format) for code-generation purpose.

```go
// @mytag (key1:"val1" key2:"val2")
func myFunc(){
}
```

## Dependency Injection

Typical-Go encourage dependency-injection using [dig](https://github.com/uber-go/dig) and annotations (`@ctor` for constructor and `@dtor` for destructor). See the [example](https://github.com/typical-go/typical-go/tree/master/examples/use-dependency-injection).

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

## Mocking

Typical-Go encourage mocking using [gomock](https://github.com/golang/mock) and annotation(`@mock`). See the [example](https://github.com/typical-go/typical-go/tree/master/examples/mock-command).

```go
type(
   // Reader responsible to read
   // @mock
   Reader interface{
      Read() error
   }
)
```

## Application Config

Typical-go encourage [application config with environment variables](https://12factor.net/config) using [envconfig](https://github.com/kelseyhightower/envconfig) and annotation (`@cfg`). See the [example](https://github.com/typical-go/typical-go/tree/master/examples/use-config).

```go
type (
   // ServerCfg configuration
   // @cfg (prefix:"SERVER")
   ServerCfg struct {
      Address string `envconfig:"ADDRESS" default:":8080" required:"true"`
   }
)
```

## Examples

Typical-Go using itself as build-tool which is an excellent example. For other examples:
- [x] [Hello World](https://github.com/typical-go/typical-go/tree/master/examples/hello-world)
- [x] [Use Config](https://github.com/typical-go/typical-go/tree/master/examples/use-config)
- [x] [Use Dependency Injection](https://github.com/typical-go/typical-go/tree/master/examples/use-dependency-injection)
- [x] [Mock Command](https://github.com/typical-go/typical-go/tree/master/examples/mock-command)
- [x] [Docker Command](https://github.com/typical-go/typical-go/tree/master/examples/docker-command)
- [x] [Custom Build-Tool](https://github.com/typical-go/typical-go/tree/master/examples/custom-build-tool)
- [x] [Custom Command](https://github.com/typical-go/typical-go/tree/master/examples/custom-command)


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
