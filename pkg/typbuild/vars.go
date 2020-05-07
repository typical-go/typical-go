package typbuild

var (
	// DefaultReleaseFolder is default value for release folder location
	DefaultReleaseFolder = "release"

	// DefaultBinFolder is default value for bin folder
	DefaultBinFolder = "bin"

	// DefaultCmdFolder is default value for cmd folder
	DefaultCmdFolder = "cmd"

	// DefaultConfigFile is default config file path
	// TODO: move typcfg package
	DefaultConfigFile = ".env"

	// Linux as releast target
	Linux ReleaseTarget = "linux/amd64"

	// Darwin as releast target
	Darwin ReleaseTarget = "darwin/amd64"
)

var (
	// ProjectPkg only available in BuildTool scope. The BuildTool must be compiled by wrapper.
	ProjectPkg string

	// TypicalTmp only available in BuildTool scope. The BuildTool must be compiled by wrapper.
	TypicalTmp string
)
