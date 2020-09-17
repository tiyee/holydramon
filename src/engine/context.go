package engine

import (
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)
const CtxStateContinue=0
const CtxStateAbort=99
type Context struct {
	*fasthttp.RequestCtx
	engine *Engine
	state  int
}

func (c *Context) Engine() *Engine {

	return c.engine
}

// return customer logger, being different from *fasthttp.RequestCtx.log
func (c *Context) Logger() *zap.Logger {
	return c.engine.logger
}
func (c *Context) Continue() {
	c.state = 0
}
func (c *Context) Abort() {
	c.state = 99
}
