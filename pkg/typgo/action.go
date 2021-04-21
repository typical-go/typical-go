package typgo

type (
	// Action responsible to execute process
	Action interface {
		Execute(*Context) error
	}
	// Actions for composite execution
	Actions []Action
	// ExecuteFn is execution function
	NewAction func(*Context) error
)

//
// NewAction
//

var _ Action = (NewAction)(nil)

// Execute action
func (a NewAction) Execute(c *Context) error {
	return a(c)
}

//
// Actions
//

var _ Action = (Actions)(nil)

// Execute actions
func (a Actions) Execute(c *Context) error {
	for _, action := range a {
		if err := action.Execute(c); err != nil {
			return err
		}
	}
	return nil
}
