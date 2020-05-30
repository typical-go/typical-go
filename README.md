# Typical Go

[![Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public.](https://www.repostatus.org/badges/latest/wip.svg)](https://www.repostatus.org/#wip)
![Go-Workflow](https://github.com/typical-go/typical-go/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/typical-go/typical-go)](https://goreportcard.com/report/github.com/typical-go/typical-go)
[![codebeat badge](https://codebeat.co/badges/a8b3c7a6-c42a-480a-acb4-68ece12f36b8)](https://codebeat.co/projects/github-com-typical-go-typical-go-master)
[![codecov](https://codecov.io/gh/typical-go/typical-go/branch/master/graph/badge.svg)](https://codecov.io/gh/typical-go/typical-go)

A Build Tool (+ Framework) for Golang. <https://typical-go.github.io/>


## Use Cases

- Framework for Build-Tool  
  Typical-Go provides levels of abstraction to develop your own build-tool. 
- Build-Tool as a framework (BAAF)  
  It is a concept where both build-tool and application utilize the same definition. We no longer see build-tool as a separate beast with the application but rather part of the same living organism. 


## Wrapper

Wrapper responsible to download, compile and run both build-tool and application through simple bash script called `typicalw`

```bash
./typicalw
```

```
Typical Build

Usage:

  ./typicalw <command> [argument]

The commands are:

  test, t     Test the project
  run, r      Run the project in local environment
  publish, p  Publish the project
  clean, c    Clean the project
  help, h     Shows a list of commands or help for one command

Use "./typicalw help <topic>" for more information about that topic
```

## Descriptor

The unique about Typical-Go is it use go-based descriptor file rather than DSL which is making it easier to understand and maintain. 

It should be defined at `typical/descriptor.go` with variable name `Descriptor`
```go 
var Descriptor = typgo.Descriptor{
	Name:    "typical-go",
	Version: "0.9.55",

	EntryPoint: wrapper.Main,
	Layouts:    []string{"wrapper", "pkg"},

	Test:    &typgo.StdTest{},
	Compile: &typgo.StdCompile{},
	Run:     &typgo.StdRun{},
	Release: &typgo.Github{Owner: "typical-go", RepoName: "typical-go"},

	Utility: typgo.NewUtility(taskExamples), // Test all the examples
}

```
## Annotation

Typical-Go support java-like annotation (expect the parameter in JSON format) for code-generation purpose.

## Dependency Injection

Typical-Go encourage dependency-injection using [dig](https://github.com/uber-go/dig) and annotation. See the [example](https://github.com/typical-go/typical-go/tree/master/examples/provide-constructor).

```go
// OpenConn open new database connection
// @constructor
func OpenConn() *sql.DB{
}
```

```go
// CloseConn close the database connection
// @destructor
func CloseConn(obj Object){
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


## Typical Tmp

Typical-tmp is an important folder that contains the build-tool mechanisms. By default, it is located in `.typical-tmp` and can be changed by hacking/editing the `typicalw` script.


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
