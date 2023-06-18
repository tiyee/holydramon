package main

import (
	"github.com/tiyee/holydramon/components/log"
	"github.com/tiyee/holydramon/engine"
	"github.com/tiyee/holydramon/router"
)

func main() {
	e := engine.New(func(eng *engine.Engine) {
		eng.SetAddr(":3003")
	})
	log.InitLogger([]string{"./logs/output.log"}, []string{"./logs/error.log"})
	router.LoadRouter(e)
	if err := e.Run(); err != nil {
		panic(err)
	}
}
