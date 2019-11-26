package runn

import (
	"fmt"
	"reflect"
)

// Execute all statement
func Execute(stmts ...interface{}) (err error) {
	for i, stmt := range stmts {
		switch stmt.(type) {
		case Runner:
			if err = stmt.(Runner).Run(); err != nil {
				return
			}
		case func() error:
			if err = stmt.(func() error)(); err != nil {
				return
			}
		default:
			return fmt.Errorf("Statement-%d: Invalid: %s", i, reflect.TypeOf(stmt))
		}
	}
	return nil
}
