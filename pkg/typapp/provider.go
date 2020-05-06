package typapp

var (
	_ Provider = (*Constructor)(nil)
	_ Provider = (*Providers)(nil)
)

type (
	// Provider responsible to provide constructor
	Provider interface {
		Constructors() []*Constructor
	}

	// Providers is list of provider
	Providers []Provider

	// Constructor details
	Constructor struct {
		Name string
		Fn   interface{}
	}
)

// Constructors is list of constructor
func (c *Constructor) Constructors() []*Constructor {
	return []*Constructor{c}
}

// Constructors is list of constructor
func (p Providers) Constructors() (ctors []*Constructor) {
	for _, provider := range p {
		ctors = append(ctors, provider.Constructors()...)
	}

	return
}
