package typrls

type (
	// Publisher responsible to publish (from release folder)
	Publisher interface {
		Publish(*Context) error
	}
	// Publishers for composite publish
	Publishers []Publisher
	// PublishFn release function
	NewPublisher func(*Context) error
)

//
// NewPublisher
//

var _ Publisher = (NewPublisher)(nil)

func (n NewPublisher) Publish(c *Context) error {
	return n(c)
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
