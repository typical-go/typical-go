package typcore

const (
	// DefaultProjectPkgVar is full path for DefaultProjectPkg variable for ldflgs purpose
	DefaultProjectPkgVar = "github.com/typical-go/typical-go/pkg/typcore.DefaultProjectPkg"

	// DefaultTypicalTmpVar is full path for DefaultTypicalTmpVar variable for ldflgs purpose
	DefaultTypicalTmpVar = "github.com/typical-go/typical-go/pkg/typcore.DefaultTypicalTmp"
)

var (
	// DefaultProjectPkg is default value for ProjectPackage which is supply by ldflags
	DefaultProjectPkg string

	// DefaultTypicalTmp is default value of typical tmp which is supply by ldflags
	DefaultTypicalTmp string
)
