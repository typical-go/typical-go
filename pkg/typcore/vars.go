package typcore

const (
	// Version of Typical-Go
	Version = "0.9.43"

	// DefaultProjectPackageVar is full path for DefaultProjectPackage variable for ldflgs purporse
	DefaultProjectPackageVar = "github.com/typical-go/typical-go/pkg/typcore.DefaultProjectPackage"

	// DefaultTypicalTmpVar is full path for DefaultTypicalTmpVar variable for ldflgs purporse
	DefaultTypicalTmpVar = "github.com/typical-go/typical-go/pkg/typcore.DefaultTypicalTmp"
)

var (
	// DefaultProjectPackage is default value for ProjectPackage
	DefaultProjectPackage string // NOTE: supply by ldflags

	// DefaultTypicalTmp is default value of typical tmp
	DefaultTypicalTmp string // NOTE: supply by ldflags
)
