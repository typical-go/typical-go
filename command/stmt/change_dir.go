package stmt

import "os"

type ChangeDirectory struct {
	ProjectPath string
}

func (c ChangeDirectory) Run() error {
	return os.Chdir(c.ProjectPath)
}
