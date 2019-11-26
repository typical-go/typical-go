package runn

import (
	"fmt"
	"reflect"
)

// Execute all statement
func Execute(stmts ...interface{}) error {
	for i, stmt := range stmts {
		switch stmt.(type) {
		case Runner:
			return stmt.(Runner).Run()
		case func() error:
			return stmt.(func() error)()
		default:
			return fmt.Errorf("Statement-%d: Invalid: %s", i, reflect.TypeOf(stmt))
		}
	}
	return nil
}
