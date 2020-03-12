package typdep

// Invoke the invocation
func Invoke(di *Container, invocations ...*Invocation) (err error) {
	for _, invocation := range invocations {
		if err = invocation.Invoke(di); err != nil {
			return
		}
	}
	return
}

// Provide the constructor
func Provide(di *Container, constructors ...*Constructor) (err error) {
	for _, constructor := range constructors {
		if err = constructor.Provide(di); err != nil {
			return
		}
	}
	return
}
