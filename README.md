[![Release](https://img.shields.io/github/release/typical-go/typical-go/all.svg)](https://github.com/typical-go/typical-go/releases/latest)
[![Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public.](https://www.repostatus.org/badges/latest/wip.svg)](https://www.repostatus.org/#wip)
![Go-Workflow](https://github.com/typical-go/typical-go/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/typical-go/typical-go)](https://goreportcard.com/report/github.com/typical-go/typical-go)
[![codebeat badge](https://codebeat.co/badges/a8b3c7a6-c42a-480a-acb4-68ece12f36b8)](https://codebeat.co/projects/github-com-typical-go-typical-go-master)
[![BCH compliance](https://bettercodehub.com/edge/badge/typical-go/typical-go?branch=master)](https://bettercodehub.com/)
[![codecov](https://codecov.io/gh/typical-go/typical-go/branch/master/graph/badge.svg)](https://codecov.io/gh/typical-go/typical-go)

# Typical Go

Build Automation Tool For Golang
- Alternative for [GNU Make](https://www.gnu.org/software/make/manual/make.html) (a.k.a makefile)
- Framework-based Build Tool (No DSL)
- Supporting Java-like annotation for code generation

## Install

```
$ go install github.com/typical-go/typical-go
```

## Setup

Setup a new project
```bash
$ typical-go setup -new -go-mod -project-pkg=github.com/typical-go/typical-go/my-project
```
- `-new` generate simple app and typical-build source
- `-go-mod` initiate go.mod
- `-project-pkg` name of project package


## Run 

It is recommended to run via wrapper [`typicalw`](typicalw) 
```bash
$ ./typicalw
```

If the wrapper is missing, you can generate it using `setup` command
```bash
$ typical-go setup
```

## Typical-Build

Typical-Build contain project descriptor and the build logic. By default, it is located in [`tools/typical-build`](tools/typical-build/typical-build) and can be changed in wrapper script.

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
      &typgo.CompileProject{},
      // run
      &typgo.RunCmd{
         Before: typgo.BuildSysRuns{"compile"},
         Action: &typgo.RunProject{},
      },
      // clean
      &typgo.CleanProject{},
      // ping
      &typgo.Command{
         Name: "ping",
         Action: typgo.NewAction(
            func(c *typgo.Context) error {
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

## Typical-Tmp

Typical-Tmp is temporary folder that contain downloaded file and other build-mechanism. By default, it is located in `.typical-tmp` and can be changed in wrapper script.

## Annotation

Typical-Go support java-like annotation (except the parameter in [StructTag](https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go) format) for code-generation purpose. [Learn more](pkg/typannot)

```go
// @mytag (key1:"val1" key2:"val2")
func myFunc(){
}
```

## Examples

Typical-Go using itself as build-tool which is an excellent example. For other examples:
- [x] [Hello World](https://github.com/typical-go/typical-go/tree/master/examples/hello-world)
- [x] [Use Dependency Injection](https://github.com/typical-go/typical-go/tree/master/examples/use-dependency-injection)
- [x] [Mock Command](https://github.com/typical-go/typical-go/tree/master/examples/mock-command)
- [x] [Custom Build-Tool](https://github.com/typical-go/typical-go/tree/master/examples/custom-build-tool)
- [x] [Custom Command](https://github.com/typical-go/typical-go/tree/master/examples/custom-command)

## See Also

- [`pkg/typapp`](pkg/typapp): Typical Application Framework
- [`pkg/typannot`](pkg/typannot): Annotation for Code Generation
- [`pkg/typmock`](pkg/typmock): Mock by Annotation
- [`pkg/typrls`](pkg/typmock): Project Releaser
- [Typical-Rest-Server](https://github.com/typical-go/typical-rest-server): Rest Server Implementation


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
