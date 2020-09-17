package hook

import (
	"fmt"
	"github.com/tiyee/holydramon/src/engine"
)

func Authority(ctx *engine.Context) {
	fmt.Println("hook authority")
}
func Logger(ctx *engine.Context) {
	fmt.Println("hook logger")
}
