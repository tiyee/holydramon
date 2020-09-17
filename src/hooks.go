package main

import (
	"github.com/tiyee/holydramon/src/engine"
	"github.com/tiyee/holydramon/src/hook"
)

func initHooks(e *engine.Engine)error  {
	e.AddHook(engine.NewHook(engine.PosBefore,hook.Authority,nil,nil).Include("/*",2))
	e.AddHook(engine.NewHook(engine.PosAfter,hook.Logger,nil,nil).Include("/*",2))
	e.AddHook(engine.NewHook(engine.PosBefore,hook.WxAuthority,nil,nil).Include("/wx",1))
	return nil
}
