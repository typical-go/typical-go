package typdep

// Providable mean the type contain Provide method
type Providable interface {
	Provide(*Container) error
}

// Provide the constructor
func Provide(di *Container, constructors ...Providable) (err error) {
	for _, constructor := range constructors {
		if err = constructor.Provide(di); err != nil {
			return
		}
	}
	return
}
