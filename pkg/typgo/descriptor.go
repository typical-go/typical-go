package typgo

type (
	// Descriptor describe the project
	Descriptor struct {
		ProjectName    string // By default is same with project folder. Only allowed characters(a-z,A-Z), underscore or dash.
		ProjectVersion string // By default it is 0.0.1
		EnvLoader      EnvLoader
		Tasks          []Tasker
	}
)
