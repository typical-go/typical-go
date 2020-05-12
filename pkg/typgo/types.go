package typgo

// Cleaner responsible to clean the project
type Cleaner interface {
	Clean(*Context) error
}

// Tester responsible to test the project
type Tester interface {
	Test(*Context) error
}

// Releaser responsible to release
type Releaser interface {
	Release(*Context) (err error)
}

// Publisher responsible to publish the release to external source
type Publisher interface {
	Publish(*Context) error
}

// Runner responsible to run the project in local environment
type Runner interface {
	Run(c *Context) error
}
