package hooks

import (
	"github.com/tiyee/holydramon/components/jwt"
	"github.com/tiyee/holydramon/engine"
)

func Authorize(c *engine.Context) {

	token := c.Authorization()
	if len(token) < 10 {
		c.AjaxError(403, "请先登录", 1)
		c.Abort()
		return
	}
	j := jwt.New[*jwt.User]([]byte("12345"))
	user := &jwt.User{}
	if err := j.Decode([]byte(token), user); err == nil {
		c.SetUserValue("jwt", user)
	} else {
		c.AjaxError(403, "请先登录", 2)
		c.Abort()
	}

}
func AuthorizeAdmin(c *engine.Context) {

	ck := c.Cookie("my_cookie")
	if len(ck) < 32 {
		c.AjaxError(405, "请先登录", 1)
		c.Abort()
		return
	}
	j := jwt.New[*jwt.User]([]byte("12345"))
	user := &jwt.User{}
	if err := j.Decode([]byte(ck), user); err == nil {
		c.SetUserValue("jwt", user)
	} else {
		c.AjaxError(403, "请先登录", 2)
		c.Abort()
	}
}
