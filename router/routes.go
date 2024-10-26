package router

import (
	"github.com/tiyee/holydramon/engine"
	"github.com/tiyee/holydramon/handles"
	"github.com/tiyee/holydramon/hooks"
)

func LoadRouter(r engine.IRouter) {

	r.GET("/test", handles.Test)
	r.Rest("/user", handles.User{})
	r.Use(engine.PosAhead, engine.Prefix("/wx", []string{}), hooks.Authorize)
	r.SetGuard(engine.Cors(hooks.Authorize))

}
