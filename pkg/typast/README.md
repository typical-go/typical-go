# typast

Java-like annotation in golang for code generation purpose

## Usage

Add `AnnotateProject` in project descriptor
```go
var descriptor = typgo.Descriptor{
    // ...

	Cmds: []typgo.Cmd{
        // ...

		// annotate
		&typast.AnnotateProject{
			Annotators: []typast.Annotator{
				// TODO: add annotator
			},
        },
	},
}
```

```bash
$ ./typicalw annotate
```

## Specification

Same with java except the parameter in struct tag format
```go
// @mytag (key1:"val1" key2:"val2")
func myFunc(){
}
```

## Create Annotator

*WIP*