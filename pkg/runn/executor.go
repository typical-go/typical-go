package runn

import "github.com/typical-go/typical-rest-server/EXPERIMENTAL/errkit"

// Executor do the code statement execution
type Executor struct {
	StopWhenError bool
}

// Execute all statement
func (e Executor) Execute(stmts ...interface{}) (err error) {
	var errs errkit.Errors
	for _, stmt := range stmts {
		switch stmt.(type) {
		case error:
			stmtErr := stmt.(error)
			if e.StopWhenError {
				return stmtErr
			}
			errs.Add(stmtErr)
		case Runner:
			runner := stmt.(Runner)
			runErr := runner.Run()
			if runErr != nil {
				if e.StopWhenError {
					return runErr
				}
				errs.Add(runErr)
			}
		}
	}
	if len(errs) > 0 {
		err = errs
	}
	return
}
