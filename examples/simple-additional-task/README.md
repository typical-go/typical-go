# Additional Build-Tool Task

Example typical-go project to demonstrate how to add additional build-tool task

Create task function 
```go
func taskPrintContext(b *typgo.BuildTool) []*cli.Command {
	return []*cli.Command{
		{
			Name:    "context",
			Aliases: []string{"ctx"},
			Usage:   "Print context as json",
			Action: func(cliCtx *cli.Context) (err error) {
				var b []byte
				if b, err = json.MarshalIndent(b, "", "    "); err != nil {
					return
				}
				fmt.Println(string(b))
				return
			},
		},
	}
}
```

Add commander to descriptor `typical/descriptor`
```go
var Descriptor = typgo.Descriptor{
	// ... 
	
	Utility: typgo.NewUtility(taskPrintContext), // Add custom task
}
```

Execute the task
```bash
./typicalw ctx
```