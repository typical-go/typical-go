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

## Setup New Project

1. Install typical-go
   ```
   $ go get -u github.com/typical-go/typical-go
   ```

2. Setup new project
   ```bash
   $ typical-go setup -new -go-mod -project-pkg=github.com/typical-go/typical-go/my-project
   ```
   - `-new` generate simple app and typical-build source
   - `-go-mod` initiate go.mod
   - `-project-pkg` name of project package

3. Generate wrapper for existing project
   ```bash
   $ typical-go setup
   ```

## Run in Current Project

It is recommended to run via wrapper [`typicalw`](typicalw). Typical-Go is automatically downloaded when not exist.
```bash
$ ./typicalw
```

## Project Descriptor

By default, project descriptor is located in [`tools/typical-build`](tools/typical-build/typical-build.go) which contain project detail and task list. It can be changed in wrapper script.

```go
package main

import (
   "fmt"

   "github.com/typical-go/typical-go/pkg/typgo"
)

var descriptor = typgo.Descriptor{
   ProjectName:    "custom-task",
   ProjectVersion: "1.0.0",

   Tasks: []typgo.Tasker{
      // test
      &typgo.GoTest{
         Args:     []string{"-timeout=30s"},
         Includes: []string{"internal/*", "pkg/*"},
      },
      // compile
      &typgo.GoBuild{},
      // run
      &typgo.RunBinary{
         Before: typgo.BuildCmdRuns{"build"},
         Action: &typgo.RunProject{},
      },
      // ping
      &typgo.Task{
         Name: "ping",
         Action: typgo.NewAction(
            func(c *typgo.Context) error {
               fmt.Println("pong")
               return nil
            },
         ),
      },
   },
}

func main() {
   typgo.Start(&descriptor)
}
```

## Temporary Folder

By default, temporary folder is located in `.typical-tmp` which contain downloaded file and other build-mechanism. It can be changed in wrapper script. Please remove temporary folder when update the typical-go version.

## Annotation Support

Typical-Go support java-like annotation (except the parameter in [StructTag](https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go) format) for code-generation purpose. [Learn more](pkg/typast)

```go
// @mytag (key1:"val1" key2:"val2")
func myFunc(){
}
```

## Learning from Examples

Typical-Go using itself as build-tool which is an excellent example. For other examples:
- [x] [hello-world](https://github.com/typical-go/typical-go/tree/master/examples/hello-world)
- [x] [typapp-sample](https://github.com/typical-go/typical-go/tree/master/examples/typapp-sample)
- [x] [typmock-sample](https://github.com/typical-go/typical-go/tree/master/examples/typmock-sample)
- [x] [custom-build-tool](https://github.com/typical-go/typical-go/tree/master/examples/custom-build-tool)
- [x] [custom-task](https://github.com/typical-go/typical-go/tree/master/examples/custom-task)

## See Also
- [Typical-Rest-Server](https://github.com/typical-go/typical-rest-server): Rest Server Implementation


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
