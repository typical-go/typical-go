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
func (e Executor) Execute(stmts ...interface{}) error {
	var errs coll.Errors
	for i, stmt := range stmts {
		var err error
		switch stmt.(type) {
		case Runner:
			err = stmt.(Runner).Run()
		case func() error:
			err = stmt.(func() error)()
		default:
			return fmt.Errorf("Statement-%d: Invalid: %s", i, reflect.TypeOf(stmt))
		}
		if err != nil {
			if e.StopWhenError {
				return err
			}
			errs.Append(err)
		}

	}
	return errs.ToError()
}
