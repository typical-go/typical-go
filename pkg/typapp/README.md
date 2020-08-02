# typapp

Typical Application framework. Similar with [uber fx](https://github.com/uber-go/fx) except provide global state for dependency-injection.

```go
// append contructor definition
typapp.AppendCtor(&typapp.Constructor{
    Fn: func() string {
        return "World"
    },
})

// append destructor definition
typapp.AppendDtor(&typapp.Destructor{
    Fn: func() {
        fmt.Println("clean something")
    },
})

// start the application
typapp.Start(func(name string) {
    fmt.Printf("Hello %s\n", name)
})

// Output: Hello World
// clean something
```


## Annotation

Pairing with typical-go, use `@ctor` to add contrustor
```go
// GetName return name value
// @ctor
func GetName() string{
    return "World"
}

// CleanSomething clean something
// @dtor
func CleanSomething() {
}

func main(){
    typapp.Start(func(name string){
        fmt.Printf("Hello %s\n", name)
    })
}
```


## Named Contructor

```go
// GetName return name value
// @ctor (name:"world")
func GetName() string{
    return "World"
}

type start struct{
    dig.In
    Name string `name:"world"`
}

func main(){
    typapp.Start(func(s start){
        fmt.Printf("Hello %s\n", s.Name)
    })
}

```


## Learn More

- [dig](https://github.com/uber-go/dig): A reflection based dependency injection toolkit for Go.