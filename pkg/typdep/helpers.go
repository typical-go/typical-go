package typdep

// InvokeAll the invocation
func InvokeAll(di *Container, invocations ...*Invocation) (err error) {
	for _, invocation := range invocations {
		if err = invocation.Invoke(di); err != nil {
			return
		}
	}
	return
}

// ProvideAll the constructor
func ProvideAll(di *Container, constructors ...*Constructor) (err error) {
	for _, constructor := range constructors {
		if err = constructor.Provide(di); err != nil {
			return
		}
	}
	return
}
