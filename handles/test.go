package handles

import "github.com/tiyee/holydramon/engine"

func Test(c *engine.Context) {
	c.String(200, "test")
}
