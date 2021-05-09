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
- Supporting java-like annotation for code generation purpose &mdash; *alternative for [go-generate](https://blog.golang.org/generate)*

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



## Build Tasks

Software development contain many build tasks like compile, test, run (locally), create the release, generate code, database migration, etc. You can add go-based task in descriptor

```go
var descriptor = typgo.Descriptor{
   Tasks: []typgo.Tasker{
      // add tasks
   },
}
```
## Standard Build Tasks

Standard build task for golang project

```go
var descriptor = typgo.Descriptor{
   Tasks: []typgo.Tasker{
      gobuild,
      gotest,
      gorun,
   },
}
```

### Compile Project

Compile the project using [go build](https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies).
```
$ ./typicalw build
$ ./typicalw b 
$ ./typicalw build [extraArguments...]
```

The default compilation is main package in `cmd/PROJECT_NAME` and output in `bin/PROJECT_NAME`
```go
gobuild := &typgo.GoBuild{}
```

With custom parameter
```go
gobuild := &typgo.GoBuild{
   MainPackage: "cmd/PROJECT_NAME",
   Output:      "bin/PROJECT_NAME",
   Ldflags: typgo.BuildVars{
      "github.com/typical-go/typical-go/pkg/typgo.ProjectName":    "PROJECT_NAME",
      "github.com/typical-go/typical-go/pkg/typgo.ProjectVersion": "v0.0.1",
   },
},
```


### Test Project

Test the project using [go-test](https://golang.org/cmd/go/#hdr-Test_packages) 
```
$ ./typicalw test
$ ./typicalw t 
$ ./typicalw t -coverprofile=cover.out
$ ./typicalw test [extraArguments...]
```

It support [glob pattern](https://en.wikipedia.org/wiki/Glob_(programming)) to include/exclude target package. 
```go
gotest := &typgo.GoTest{
   Includes: []string{"internal/app/**", "pkg/**"},
   Excludes: []string{"internal/app/model"},
}
```

With arguments
```go
gotest := &typgo.GoTest{
   Timeout:  60 * time.Second,
   NoCover:  false,
   Verbose:  false,
   Includes: []string{"internal/app/**", "pkg/**"},
   Excludes: []string{"internal/app/model"},
}
```

### Run Project

Run the project 
```
$ ./typicalw run
$ ./typicalw r
$ ./typicalw r [extraArguments...]
```

Execute annotate and compile before run
```go
run := &typgo.RunBinary{
   Before: typgo.TaskNames{"generate", "build"},
}
```


## Custom Build Tasks


### Calling Other Tasks

```go
callOtherTask := &typgo.Task{
   Name:   "other-tasks",
   Usage:  "call other-tasks",
   Action: typgo.TaskNames{"ping", "info", "help"},
},
```

### Calling Bash Command

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

### In-code Implemention

Simple example
```go
pingTask := &typgo.Task{
   Name:  "ping",
   Usage: "print pong",
   Action: typgo.NewAction(func(c *typgo.Context) error {
      c.Info("pong")
      return nil
   }),
}
```

Call bash command
```go
infoTask := &typgo.Task{
   Name:  "info",
   Usage: "print info",
   Action: typgo.NewAction(func(c *typgo.Context) error {
      c.Info("print the info:")
      c.ExecuteBash("go version")
      c.ExecuteBash("git version")
      return nil
   }),
},
```

### Sub Task

```go
databaseTool := &typgo.Task{
   Name:    "database",
   Aliases: []string{"db"},
   Usage:   "database tool",
   SubTasks: []*typgo.Task{
      {
         Name:  "create",
         Usage: "create database",
         Action: typgo.NewAction(...),
      },
      {
         Name:  "drop",
         Usage: "drop database",
         Action: typgo.NewAction(...),
      },
      {
         Name:  "migrate",
         Usage: "migrate database",
         Action: typgo.NewAction(...),
      },
      {
         Name:  "seed",
         Usage: "seed database",
         Action: typgo.NewAction(...),
      },
   },
},
```

### Tasker Interface

Parameterized task by implemented `typgo.Tasker` 
```go
type greetTask struct {
	person string
}

var _ typgo.Tasker = (*greetTask)(nil)
var _ typgo.Action = (*greetTask)(nil)

func (g *greetTask) Task() *typgo.Task {
	return &typgo.Task{
		Name:   "greet",
		Usage:  "greet person",
		Action: g,
	}
}

func (g *greetTask) Execute(c *typgo.Context) error {
	info.Infof("Hello %s\n", g.person)
	return nil
}
```

## Generate Code

`typgen` is similar with [`go generate`](https://blog.golang.org/generate) except it provide in-code implementation with java-like annotation support

Add generate task in [project-descriptor](#project-descriptor)
```go
var descriptor = typgo.Descriptor{
	
	Tasks: []typgo.Tasker{
		&typgen.Generator{
			Processor: typgen.Processors{
				// add annotation director
			},
		},
	
	},
}

```

The directive is similar with java except the parameter in [StructTag](https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go) format
```go
// @mytag (key1:"val1" key2:"val2")
func myFunc(){
}
```

Trigger generate task
```
$ ./typicalw generate
```



## Dependency Injection

[typapp](pkg/typapp) package is application framework with dependency injection and graceful shutdown. It using reflection-based dependency injection ([dig](https://github.com/uber-go/dig)). 

It is similar with [fx](https://github.com/uber-go/fx) except encourage global state. 
```go
typapp.Reset() // make sure constructor and container is empty (optional)
typapp.Provide("", func() string { return "world" })

err := typapp.Invoke(func(text string) {
   fmt.Printf("hello %s\n", text)
})
if err != nil {
   log.Fatal(err)
}

// Output: hello world
```

### Ctor Annotation

Use `@ctor` to add constructor and import the side-effect to initiate provide constructor
```go
import _ "PROJECT_PACKAGE/internal/generated/ctor" // import the side-effect

// @ctor
func GetName() string{ return "World" }

// @ctor (name:"other")
func GetOtherName() string{ return "Other" }

func main(){
    typapp.Invoke(func(name string){
        fmt.Printf("Hello %s\n", name)
    })
}
```

Add annotation processor in [generator](#generate-code) and trigger generate
```go
generator := &typgen.Generator{
   Processor: typgen.Processors{
      &typapp.CtorAnnot{}, // add CtorAnnot
   },
}
```

### Named Values

Using `dig.In` for named values (https://godoc.org/go.uber.org/dig#hdr-Named_Values)
```go
typapp.Provide("t1", func() string { return "hello" }) // provide same type
typapp.Provide("t2", func() string { return "world" }) // provide same type

type param struct {
   dig.In
   Text1 string `name:"t1"`
   Text2 string `name:"t2"`
}

printHello := func(p param) {
   fmt.Printf("%s %s\n", p.Text1, p.Text2)
}

if err := typapp.Invoke(printHello); err != nil {
   log.Fatal(err)
}

// Output: hello world
```


### Gracefully Stop

Use `StartApp()` to support gracefully stop
```go
typapp.Provide("", func() string { return "world" })

startFn := func(text string) { fmt.Printf("hello %s\n", text) }
stopFn := func() { fmt.Println("bye2") }

if err := typapp.StartApp(startFn, stopFn); err != nil {
   log.Fatal(err)
}

// Output: hello world
// bye2
```



## Generate mock

Generate mock using [gomock](https://github.com/golang/mock) with annotation
```
$ ./typicalw mock
```

Add generate mock task
```go
genMock := &typmock.GoMock{
   Sources: []string{"internal"},
}
```

Add `@mock` annotation to the interface
```go
// Reader responsible to read
// @mock
type Reader interface{
    Read()
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
