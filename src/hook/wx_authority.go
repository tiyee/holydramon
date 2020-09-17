package hook

import (
	"fmt"
	"github.com/tiyee/holydramon/src/engine"
	"github.com/tiyee/holydramon/src/service"
)

func WxAuthority(ctx *engine.Context) {

	signature := ctx.QueryArgs().Peek("signature")
	timestamp := ctx.QueryArgs().Peek("timestamp")
	nonce := ctx.QueryArgs().Peek("nonce")
	fmt.Println(signature, timestamp, nonce)
	if false == service.CheckSignature(string(signature), string(timestamp), string(nonce), engine.ImmutableConfig.Wx.Token) {
		ctx.Logger().Error("wx_authority fail")
		ctx.SuccessString("text/html", "")
		ctx.Abort()
	}
}
