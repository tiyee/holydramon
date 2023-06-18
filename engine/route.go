package engine

type HandlerFunc func(ctx *Context)
type Route struct {
	Method       string
	Path         string
	BeforeHooks  []HandlerFunc
	AfterHooks   []HandlerFunc
	HandlerFuncs []HandlerFunc
}
type Routes map[string]*Route
type IRouter interface {
	POST(path string, fn ...HandlerFunc)
	GET(path string, fn ...HandlerFunc)
	PUT(path string, fn ...HandlerFunc)
	DELETE(path string, fn ...HandlerFunc)
	OPTIONS(path string, fn ...HandlerFunc)
	Use(pos HookPos, matcher IMatcher, fn HandlerFunc)
	SetRest(path string, rest any)
}
type IHttpGET interface {
	GET(ctx *Context)
}
type IHttpPOST interface {
	POST(ctx *Context)
}
type IHttpPUT interface {
	PUT(ctx *Context)
}
type IHttpDELETE interface {
	DELETE(ctx *Context)
}
type IHttpOPTIONS interface {
	OPTIONS(ctx *Context)
}
type IHttpPATCH interface {
	PATCH(ctx *Context)
}
type IHttpTRACE interface {
	TRACE(ctx *Context)
}
type IHttpHEAD interface {
	HEAD(ctx *Context)
}
type IHttpCONNECT interface {
	CONNECT(ctx *Context)
}
