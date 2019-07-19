package typicore

// Statement statement
type Statement interface {
	Run() error
}
