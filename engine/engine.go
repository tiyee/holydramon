package engine

import (
	"errors"
	"fmt"
	"github.com/tiyee/holydramon/components/log"
	"net/http"
	"strings"
	"sync"
)

type Engine struct {
	routes Routes
	addr   string
	pool   sync.Pool
	hooks  []Hook
	guards []Guard
}
type IEnginOpt func(e *Engine)

func New(opts ...IEnginOpt) *Engine {
	engine := &Engine{
		routes: make(map[string]*Route, 0),
		addr:   ":3003",
		hooks:  make([]Hook, 0),
		guards: make([]Guard, 0),
	}
	engine.pool.New = func() any {
		return engine.allocateContext()
	}
	for _, opt := range opts {
		opt(engine)
	}
	return engine

}
func (e *Engine) SetAddr(addr string) {
	e.addr = addr
}
func (e *Engine) allocateContext() *Context {

	return &Context{r: nil, w: nil, handlers: make([]HandlerFunc, 0), index: 0}
}
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodS := r.Method
	pathS := r.URL.Path
	key := methodS + pathS
	context := e.pool.Get().(*Context)
	context.reset(w, r)
	if route, exist := e.routes[key]; exist {
		for _, fn := range route.BeforeHooks {
			context.handlers = append(context.handlers, fn)
		}
		context.handlers = append(context.handlers, route.HandlerFuncs...)
		for _, fn := range route.AfterHooks {
			context.handlers = append(context.handlers, fn)
		}
		defer func() {
			if err := recover(); err != nil {
				log.Error("recover", log.String("error", fmt.Sprintf("%v", err)))
				context.Error("internal error", 500)
			}
		}()
		defer func() {
			if err := log.Sync(); err != nil {
				fmt.Println(err.Error())
			}
		}()
		context.index = -1
		context.Next()
		e.pool.Put(context)
	} else {
		context.NotFound()
	}
}
func (e *Engine) setRoute(method string, path string, fn ...HandlerFunc) {
	path = "/" + strings.Trim(path, "/")
	key := method + path
	e.routes[key] = &Route{
		Method:       method,
		Path:         path,
		HandlerFuncs: fn,
		BeforeHooks:  make([]HandlerFunc, 0),
		AfterHooks:   make([]HandlerFunc, 0),
	}
}
func (e *Engine) SetGuard(gd Guard) {
	e.guards = append(e.guards, gd)
}
func (e *Engine) Run() (err error) {
	e.dispatch()
	if len(e.addr) < 3 {
		return errors.New("empty addr")
	}
	if len(e.routes) == 0 {
		return errors.New("empty routes")
	}
	return http.ListenAndServe(e.addr, e)

}
func (e *Engine) GET(path string, fn ...HandlerFunc) {
	e.setRoute(http.MethodGet, path, fn...)
}
func (e *Engine) POST(path string, fn ...HandlerFunc) {
	e.setRoute(http.MethodPost, path, fn...)
}
func (e *Engine) PUT(path string, fn ...HandlerFunc) {
	e.setRoute(http.MethodPut, path, fn...)
}
func (e *Engine) DELETE(path string, fn ...HandlerFunc) {

	e.setRoute(http.MethodDelete, path, fn...)
}
func (e *Engine) OPTIONS(path string, fn ...HandlerFunc) {
	e.setRoute(http.MethodOptions, path, fn...)
}
func (e *Engine) PATCH(path string, fn ...HandlerFunc) {
	e.setRoute(http.MethodPatch, path, fn...)
}
func (e *Engine) Rest(path string, rest any) {
	e.SetRest(path, rest)
}
func (e *Engine) Union(methods []string, path string, fn ...HandlerFunc) {
	for _, method := range methods {
		e.setRoute(method, path, fn...)
	}
}
func (e *Engine) SetRest(path string, rest any) {
	flag := 0
	if hd, ok := rest.(IHttpGET); ok {
		e.setRoute(http.MethodGet, path, hd.GET)
		flag += 1
	}
	if hd, ok := rest.(IHttpPOST); ok {
		e.setRoute(http.MethodPost, path, hd.POST)
		flag += 2
	}
	if hd, ok := rest.(IHttpPUT); ok {
		e.setRoute(http.MethodPut, path, hd.PUT)
		flag += 4
	}
	if hd, ok := rest.(IHttpDELETE); ok {
		e.setRoute(http.MethodDelete, path, hd.DELETE)
		flag += 8
	}
	if hd, ok := rest.(IHttpOPTIONS); ok {
		e.setRoute(http.MethodOptions, path, hd.OPTIONS)
		flag += 16
	}
	if hd, ok := rest.(IHttpPATCH); ok {
		e.setRoute(http.MethodPatch, path, hd.PATCH)
		flag += 32
	}
	if hd, ok := rest.(IHttpHEAD); ok {
		e.setRoute(http.MethodHead, path, hd.HEAD)
		flag += 64
	}
	if hd, ok := rest.(IHttpTRACE); ok {
		e.setRoute(http.MethodTrace, path, hd.TRACE)
		flag += 128
	}
	if hd, ok := rest.(IHttpCONNECT); ok {
		e.setRoute(http.MethodConnect, path, hd.CONNECT)
		flag += 256
	}
	if flag == 0 {
		panic("empty handle in path: " + path)
	}
}

func (e *Engine) dispatch() {
	pathSets := make(map[string]struct{})
	for _, route := range e.routes {
		pathSets[route.Path+route.Method] = struct{}{}
	}

	for _, guard := range e.guards {
		for _, route := range e.routes {
			//如果方法本身存在，就当中间件，否则就当一个路由
			if _, exist := pathSets[route.Path+string(guard.Method)]; exist {
				e.Use(guard.Pos, Identical(route.Path), guard.HandlerFunc)
			} else {
				e.setRoute(string(guard.Method), route.Path, guard.HandlerFunc)
			}
		}

	}
	for _, route := range e.routes {
		for _, hook := range e.hooks {
			if hook.matcher.Match(route.Method, route.Path) {
				if hook.Pos == PosAhead {
					route.BeforeHooks = append(route.BeforeHooks, hook.HandlerFunc)
				}
				if hook.Pos == PosBehind {
					route.AfterHooks = append(route.AfterHooks, hook.HandlerFunc)
				}
			}
		}
	}
}
