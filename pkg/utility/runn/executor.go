package runn

import (
	"fmt"
	"reflect"

	"github.com/typical-go/typical-go/pkg/utility/coll"
)

// Executor do the code statement execution
type Executor struct {
	StopWhenError bool
}

// Execute all statement
func (e Executor) Execute(stmts ...interface{}) (err error) {
	var errs coll.Errors
	for i, stmt := range stmts {
		switch stmt.(type) {
		case Runner:
			runner := stmt.(Runner)
			runErr := runner.Run()
			if runErr != nil {
				if e.StopWhenError {
					return runErr
				}
				errs.Add(runErr)
			}
		default:
			err = fmt.Errorf("Statement-%d: Invalid: %s", i, reflect.TypeOf(stmt))
			return
		}

	}
	if len(errs) > 0 {
		err = errs
	}
	return
}
