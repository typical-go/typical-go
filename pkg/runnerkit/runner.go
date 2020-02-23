package runnerkit

import (
	"context"
	"fmt"
	"reflect"
)

// Runner contain run function
type Runner interface {
	Run(context.Context) error
}

// Run all runnerkit
func Run(ctx context.Context, stmts ...interface{}) (err error) {
	for i, stmt := range stmts {
		switch stmt.(type) {
		case Runner:
			if err = stmt.(Runner).Run(ctx); err != nil {
				return
			}
		case func(context.Context) error:
			if err = stmt.(func(context.Context) error)(ctx); err != nil {
				return
			}
		default:
			return fmt.Errorf("Statement-%d: Invalid: %s", i, reflect.TypeOf(stmt))
		}
	}
	return nil
}
