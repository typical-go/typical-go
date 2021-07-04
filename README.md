[![Release](https://img.shields.io/github/release/typical-go/typical-go/all.svg)](https://github.com/typical-go/typical-go/releases/latest)
[![Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public.](https://www.repostatus.org/badges/latest/wip.svg)](https://www.repostatus.org/#wip)
![Go-Workflow](https://github.com/typical-go/typical-go/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/typical-go/typical-go)](https://goreportcard.com/report/github.com/typical-go/typical-go)
[![codebeat badge](https://codebeat.co/badges/a8b3c7a6-c42a-480a-acb4-68ece12f36b8)](https://codebeat.co/projects/github-com-typical-go-typical-go-master)
[![BCH compliance](https://bettercodehub.com/edge/badge/typical-go/typical-go?branch=master)](https://bettercodehub.com/)
[![codecov](https://codecov.io/gh/typical-go/typical-go/branch/master/graph/badge.svg)](https://codecov.io/gh/typical-go/typical-go)
# Typical Go

Build Automation Tool For Golang
- Manage build tasks &mdash; *alternative for [makefile](https://www.gnu.org/software/make/manual/make.html)*
- Framework-based Build Tool &mdash; *no DSL to be learned, write build task in Go*
- Wrapper Script  &mdash; *single script to prepare and run the build-tool*
- Supporting java-like annotation for code generation &mdash; *alternative for [go-generate](https://blog.golang.org/generate)*

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

4. Run the project
   ```
   $ typical-go run
   ```
   Or via wrapper (recommendation)
   ```
   $ ./typicalw 
   ```

Check [examples/my-project](https://github.com/typical-go/typical-go/tree/master/examples/my-project) for what generated new project look like

## Wrapper Script

[`typicalw`](typicalw) is a bash script to prepare and run the build-tool. 
```
$ ./typicalw
```

You can hack the parameters accordingly
```bash
PROJECT_PKG="github.com/typical-go/typical-go"
BUILD_TOOL="tools/typical-build"
TYPTMP=.typical-tmp
TYPGO=$TYPTMP/bin/typical-go
TYPGO_SRC=github.com/typical-go/typical-go
```

Any downloaded or required file will be saved in temporary folder which is located in `.typical-tmp` in project directory including typical-go itself. Its mean you don't need to install typical-go manually and the project always use designed version. 

To update typical-go to new version
```
$ go get -u go get -u github.com/typical-go/typical-go
$ rm -rf .typical-tmp
```

## Project Descriptor

By default, project descriptor is located in [`tools/typical-build/typical-build.go`](tools/typical-build/typical-build.go) which contain project detail and task list.

```go
var descriptor = typgo.Descriptor{
   ProjectName:    "application-name",
   ProjectVersion: "1.0.0",

   Environment: typgo.DotEnv(".env"),

   Tasks: []typgo.Tasker{
      // test
      &typgo.GoTest{
         Includes: []string{"internal/*", "pkg/*"},
      },
      // build
      &typgo.GoBuild{},
      // run
      &typgo.RunBinary{
         Before: typgo.TaskNames{"build"},
      },
   },
}
```

The descriptor file is regular golang file that will be compiled by typical-go, so main function should be defined.
```go
func main() {
	typgo.Start(&descriptor)
}
```

It is possible to use other custom build-tool framework, check [examples/custom-build-tool](https://github.com/typical-go/typical-go/tree/master/examples/custom-build-tool) for example.

## Wiki

Check [wiki](https://github.com/typical-go/typical-go/wiki) for more documentation
- [App Framework](https://github.com/typical-go/typical-go/wiki/App-Framework)
- [Standard Build Task](https://github.com/typical-go/typical-go/wiki/Standard-Build-Tasks)
- [Custom Build Tasks](https://github.com/typical-go/typical-go/wiki/Custom-Build-Tasks)
- [Generate Code](https://github.com/typical-go/typical-go/wiki/Generate-Code)
- [Generate Mock](https://github.com/typical-go/typical-go/wiki/Generate-Mock)

## Examples

Typical-Go using itself as build-tool which is an excellent example. 

For other examples:
- [hello-world](https://github.com/typical-go/typical-go/tree/master/examples/hello-world)
- [typapp-sample](https://github.com/typical-go/typical-go/tree/master/examples/typapp-sample)
- [typmock-sample](https://github.com/typical-go/typical-go/tree/master/examples/typmock-sample)
- [custom-build-tool](https://github.com/typical-go/typical-go/tree/master/examples/custom-build-tool)
- [custom-task](https://github.com/typical-go/typical-go/tree/master/examples/custom-task)
- [my-project](https://github.com/typical-go/typical-go/tree/master/examples/my-project): generated from setup command
- [Typical-Rest-Server](https://github.com/typical-go/typical-rest-server): Rest Server Implementation

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
