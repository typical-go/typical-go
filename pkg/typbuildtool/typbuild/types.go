package typbuild

// Builder reponsible to build
type Builder interface {
	Build(c *Context) (bin string, err error)
}

// Cleaner responsible to clean the project
type Cleaner interface {
	Clean(*Context) error
}

// Prebuilder responsible to prebuild
type Prebuilder interface {
	Prebuild(c *Context) error
}

// Tester responsible to test the project
type Tester interface {
	Test(*Context) error
}

// Runner responsible to run the project
type Runner interface {
	Run(*RunContext) error
}

// Mocker responsible to mock
type Mocker interface {
	Mock(*Context) error
}
