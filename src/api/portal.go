package api

import (
	"github.com/tiyee/holydramon/src/engine"
	"go.uber.org/zap"
)

func Portal(c *engine.Context) {
	c.Logger().Error("test", zap.String("my_key", "xxxxxxxxxxxx"))
	c.Success("text/html", []byte("111111"))
	c.Abort()
}
func Wx(c *engine.Context) {
	c.Logger().Error("wx", zap.String("22222", "xxxxxxxxxxxx"))
	c.Success("text/html", []byte("wx call"))
	return
}
func Test(c *engine.Context) {
	c.Logger().Error("test", zap.String("22222", "xxxxxxxxxxxx"))
	c.Success("text/html", []byte("test"))
	return
}
