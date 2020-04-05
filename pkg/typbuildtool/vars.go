package typbuildtool

const (
	// DefaultReleaseFolder is default value for release folder location
	DefaultReleaseFolder = "release"

	// DefaultBinFolder is default value for bin folder
	DefaultBinFolder = "bin"

	// DefaultCmdFolder is default value for cmd folder
	DefaultCmdFolder = "cmd"

	// DefaultConfigFile is default config file path
	DefaultConfigFile = ".env"

	// DefaultEnablePrecondition is default precondition flag
	DefaultEnablePrecondition = true

	// Linux as releast target
	Linux ReleaseTarget = "linux/amd64"

	// Darwin as releast target
	Darwin ReleaseTarget = "darwin/amd64"
)
