package typcore

// DefaultLayout is default project layout
var DefaultLayout = ProjectLayout{
	Cmd:     "cmd",
	Bin:     "bin",
	Temp:    ".typical-tmp",
	Mock:    "mock",
	Release: "release",
}

// ProjectLayout is reflect folder structure of the project
type ProjectLayout struct {
	Bin     string
	Cmd     string
	Temp    string // TODO: temp folder is not part project layout as it is constant for all typical-go
	Mock    string // TODO: mock folder is not part project layout but rather mock generator
	Release string // TODO: consider release folder as project layout
}
