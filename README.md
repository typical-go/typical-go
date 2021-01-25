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

## Getting Started

1. Install typical-go (Optional, only needed to setup new project)
   ```
   $ go get -u github.com/typical-go/typical-go
   ```
2. Setup new project
   ```
   $ typical-go setup -new -go-mod -project-pkg=[PACKAGE_NAME]
   ```
   - `-new` generate simple app and typical-build source
   - `-go-mod` initiate go.mod
   - `-project-pkg` name of project package

3. Generate wrapper for existing project
   ```
   $ typical-go setup
   ```

4. Run the wrapper [`typicalw`](typicalw). If typical-go not exist then it is automatically downloaded.
   ```
   $ ./typicalw
   ```

Check [examples/my-project](https://github.com/typical-go/typical-go/tree/master/examples/my-project) for what generated new project look like

## Wrapper

The wrapper is bash script to download, build and run the build-tool. 

Any downloaded and required file will be store in temporary folder which is located in `.typical-tmp`. Temporary folder is recommended to be deleted after updating typical-go version.

You can hack some parameter accordingly in wrapper script.
```bash
PROJECT_PKG="github.com/typical-go/typical-go"
BUILD_TOOL="tools/typical-build"
TYPTMP=.typical-tmp
TYPGO=$TYPTMP/bin/typical-go
TYPGO_SRC=github.com/typical-go/typical-go
```

## Project Descriptor

By default, project descriptor is located in [`tools/typical-build/typical-build.go`](tools/typical-build/typical-build.go) which contain project detail and task list.

```go
var descriptor = typgo.Descriptor{
   ProjectName:    "application-name",
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
         Before: typgo.TaskNames{"build"},
      },
   },
}
```

The main function must call `typgo.Start()` to compile the descriptor struct to the actual build-tool.  
```go
func main() {
	typgo.Start(&descriptor)
}
```

Check [examples/custom-build-tool](https://github.com/typical-go/typical-go/tree/master/examples/custom-build-tool) for example simple custom build-tool if you need to develop your own custom-build-tool without typical-go framework.

## Build Tasks

- Custom task with golang code
   ```go
   pingTask := &typgo.Task{
      Name:  "ping",
      Usage: "print pong",
      Action: typgo.NewAction(func(c *typgo.Context) error {
         fmt.Println("pong")
         return nil
      }),
   }
   ```

- Custom task to call bash command
   ```go
   helpTask := &typgo.Task{
      Name:  "help",
      Usage: "print help",
      Action: &typgo.Bash{
         Name:   "go",
         Args:   []string{"help"},
         Stdout: os.Stdout,
      },
   },
   ```

- Custom task to call multiple bash command and golang code
   ```go
   infoTask := &typgo.Task{
      Name:  "info",
      Usage: "print info",
      Action: typgo.NewAction(func(c *typgo.Context) error {
         fmt.Println("print the info:")
         c.ExecuteBash("go version")
         c.ExecuteBash("git version")
         return nil
      }),
   },
   ```

- Custom task to call other task
   ```go
   allTask := &typgo.Task{
      Name:   "all",
      Usage:  "run all custom task",
      Action: typgo.TaskNames{"ping", "info", "help"},
   },
   ```

## Annotation

Typical-Go support java-like annotation (except the parameter in [StructTag](https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go) format) for code-generation purpose. [Learn more](pkg/typast)

```go
// @mytag (key1:"val1" key2:"val2")
func myFunc(){
}
```


## Learning from Examples

Typical-Go using itself as build-tool which is an excellent example. For other examples:
- [hello-world](https://github.com/typical-go/typical-go/tree/master/examples/hello-world)
- [typapp-sample](https://github.com/typical-go/typical-go/tree/master/examples/typapp-sample)
- [typmock-sample](https://github.com/typical-go/typical-go/tree/master/examples/typmock-sample)
- [custom-build-tool](https://github.com/typical-go/typical-go/tree/master/examples/custom-build-tool)
- [custom-task](https://github.com/typical-go/typical-go/tree/master/examples/custom-task)
- [my-project](https://github.com/typical-go/typical-go/tree/master/examples/my-project): generated from setup command
- [Typical-Rest-Server](https://github.com/typical-go/typical-rest-server): Rest Server Implementation

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
