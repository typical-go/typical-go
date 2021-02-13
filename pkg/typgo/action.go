package typgo

type (
	// Action responsible to execute process
	Action interface {
		Execute(*Context) error
	}
	// Actions for composite execution
	Actions []Action
	// ExecuteFn is execution function
	ExecuteFn  func(*Context) error
	actionImpl struct {
		fn ExecuteFn
	}
)

//
// actionImpl
//

// NewAction return new instance of Action
func NewAction(fn ExecuteFn) Action {
	return &actionImpl{fn: fn}
}

// Execute action
func (a *actionImpl) Execute(c *Context) error {
	return a.fn(c)
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
