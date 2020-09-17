package main

import (
	"github.com/tiyee/holydramon/src/api"
	"github.com/tiyee/holydramon/src/engine"
)

func initRouter(e *engine.Engine) error {
	e.GET("/", api.Portal)
	e.GET("/wx", api.Wx)
	e.GET("/test", api.Test)
	return nil
}
