# Provide Constructor

Example typical-go project to demonstrate how to provide constructor to dependency injection

Add `CtorAnnot` and `DtorAnnotation` in annotators
```go
var descriptor = typgo.Descriptor{
	Cmds: []typgo.Cmd{
		// annotate
		&typast.AnnotateProject{
			Annotators: []typast.Annotator{
				&typapp.CtorAnnot{}, // constructor
				&typapp.DtorAnnotation{}, // destructor
			},
		},

		// ...
	},
}
```

## Constructor

Put `@ctor` annotation in constructor function comment
```go
// HelloWorld text
// @ctor
func HelloWorld() string {
	return "Hello World"
}

// HelloTypical text
// @ctor (name: "typical")
func HelloTypical() string {
	return "Hello Typical"
}
```




