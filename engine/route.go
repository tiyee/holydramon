package engine

type HandlerFunc func(ctx *Context)
type Route struct {
	Method       string
	Methods      int
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
	Rest(path string, rest any)
	Union(methods []string, path string, fn ...HandlerFunc)
	SetGuard(gd Guard)
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
