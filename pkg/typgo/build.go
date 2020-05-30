package typgo

const (
	CompilePhase Phase = iota
	RunPhase
)

type (
	// Phase of build process
	Phase int

	// Build responsible to execute build process
	Build interface {
		Execute(*Context, Phase) (bool, error)
	}

	// Builds is array of build
	Builds []Build
)

var _ Build = (Builds)(nil)

func (d Phase) String() string {
	return [...]string{
		"compile_phase",
		"run_phase",
	}[d]
}

// Execute build
func (b Builds) Execute(ctx *Context, phase Phase) (bool, error) {
	var ok bool
	for _, build := range b {
		ok1, err := build.Execute(ctx, phase)
		if err != nil {
			return ok1, err
		}
		ok = ok || ok1
	}
	return ok, nil
}
