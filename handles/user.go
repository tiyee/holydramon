package handles

import "github.com/tiyee/holydramon/engine"

type User struct {
}

// just check if it implements interface
var _ engine.IHttpGET = User{}
var _ engine.IHttpPOST = User{}

func (u User) GET(c *engine.Context) {
	c.String(200, "get")
}
func (u User) POST(c *engine.Context) {
	c.String(200, "get")
}
