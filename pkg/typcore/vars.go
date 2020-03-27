package typcore

const (
	// Version of Typical-Go
	Version = "0.9.46"

	// DefaultProjectPackageVar is full path for DefaultProjectPackage variable for ldflgs purpose
	DefaultProjectPackageVar = "github.com/typical-go/typical-go/pkg/typcore.DefaultProjectPackage"

	// DefaultTypicalTmpVar is full path for DefaultTypicalTmpVar variable for ldflgs purpose
	DefaultTypicalTmpVar = "github.com/typical-go/typical-go/pkg/typcore.DefaultTypicalTmp"
)

var (
	// DefaultProjectPackage is default value for ProjectPackage which is supply by ldflags
	DefaultProjectPackage string

	// DefaultTypicalTmp is default value of typical tmp which is supply by ldflags
	DefaultTypicalTmp string
)
