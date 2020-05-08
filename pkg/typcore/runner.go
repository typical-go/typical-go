package typcore

var (
	_ Runner = (*SimpleRunner)(nil)
)

type (
	// Runner responsible to run the application
	Runner interface {
		Run(*Descriptor) error
	}

	// SimpleRunner is simple implementation of App
	SimpleRunner struct {
		fn func(*Descriptor) error
	}
)

// Run fn as a runner
func Run(fn func(*Descriptor) error) *SimpleRunner {
	return &SimpleRunner{fn: fn}
}

// Run the simple app
func (a *SimpleRunner) Run(d *Descriptor) (err error) {
	return a.fn(d)
}
