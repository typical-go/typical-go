package typrls

type (
	// Publisher responsible to publish (from release folder)
	Publisher interface {
		Publish(*Context) error
	}
	// Publishers for composite publish
	Publishers []Publisher
	// PublishFn release function
	PublishFn     func(*Context) error
	publisherImpl struct {
		fn PublishFn
	}
)

//
// publisherImpl
//

// NewPublisher return new instance of Releaser
func NewPublisher(fn PublishFn) Publisher {
	return &publisherImpl{fn: fn}
}

func (r *publisherImpl) Publish(c *Context) error {
	return r.fn(c)
}

//
// Publishers
//

var _ Publisher = (Publishers)(nil)

// Publish the release
func (p Publishers) Publish(c *Context) (err error) {
	for _, publisher := range p {
		if err = publisher.Publish(c); err != nil {
			return
		}
	}
	return
}
