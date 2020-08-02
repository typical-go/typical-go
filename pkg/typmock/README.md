# typmock

Generate mock using [gomock](https://github.com/golang/mock) with annotation

```go
// Reader responsible to read
// @mock
type Reader interface{
    Read()
}
```

Add `typmock.MockCmd` to typical-build
```go
var descriptor = typgo.Descriptor{

    // ...

	Cmds: []typgo.Cmd{

        // ...
        
		// mock
		&typmock.MockCmd{},
	},
}
```

Generate gomock
```bash
./typicalw mock
```