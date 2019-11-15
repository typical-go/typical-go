package app

// Module of application
func Module() interface{} {
	return applicationModule{}
}

type applicationModule struct {
}

func (m applicationModule) Run() interface{} {
	return start
}
