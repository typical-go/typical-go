# Custom Build Tool

Feel free to create custom build-tool without restriction in `tools/typical-build`. 
```go
func main() {
	ctx := context.Background()
	output := "bin/custom-build-tool"
	mainPackage := "./cmd/custom-build-tool"
	execkit.Run(ctx, &execkit.GoBuild{MainPackage: mainPackage, Output: output})
	execkit.Run(ctx, &execkit.Command{Name: output, Stdout: os.Stdout, Stderr: os.Stderr})
}
```


You also can modify `.typicalw` wrapper to set build-tool src, temporary folder and project package.

```bash
#!/bin/bash

set -e

TYPTMP=.typical-tmp                            # temporary file location 
TYPGO=$TYPTMP/bin/typical-go                   # typical-go output
TYPGO_SRC=github.com/typical-go/typical-go     # typical-go source
BUILDTOOL_SRC=tools/typical-build              # build-tool source
PROJECT_PKG=github.com/typical-go/typical-go/examples/custom-command  # project package

if ! [ -s $TYPGO ]; then
	echo "Build typical-go"
	go build -o $TYPGO $TYPGO_SRC
fi

$TYPGO run \
	-src=$BUILDTOOL_SRC \
	-project-pkg=$PROJECT_PKG \
	-typical-tmp=$TYPTMP \
	$@
```

Remove typical tmp to reset the build-tool binary
```bash
rm -rf .typical-tmp
```