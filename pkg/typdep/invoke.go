package typdep

// Invokable means the type contain Invoke method
type Invokable interface {
	Invoke(*Container) error
}

// Invoke the invocation
func Invoke(di *Container, invocations ...Invokable) (err error) {
	for _, invocation := range invocations {
		if err = invocation.Invoke(di); err != nil {
			return
		}
	}
	return
}
