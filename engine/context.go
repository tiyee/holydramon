package engine

import (
	"context"
	"encoding/json"

	"github.com/tiyee/holydramon/helps"
	"github.com/tiyee/holydramon/schema"
	"github.com/tiyee/holydramon/vo"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Context struct {
	ctx         context.Context
	w           http.ResponseWriter
	r           *http.Request
	handlers    []HandlerFunc
	index       int8
	userData    map[string]any
	queryParsed bool
	queryMap    url.Values
	domain      string
}

const abortIndex int8 = math.MaxInt8 >> 1

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{w: w, r: r}
}

func (c *Context) Request() *http.Request {

	return c.r
}
func (c *Context) Response() http.ResponseWriter {
	return c.w
}
func (c *Context) Ctx() context.Context {
	return c.ctx
}
func (c *Context) JSONArgs(v schema.ISchema) error {
	if err := helps.JSONArgs(c.Request().Body, v); err != nil {
		return err
	}
	v.Hook()
	return nil

}
func (c *Context) String(code int, text string) {
	c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
	c.Response().WriteHeader(code)
	c.Response().Write([]byte(text))

}
func (c *Context) AjaxJson(code int, message string, data any) {
	ret := vo.Base{
		Error:   code,
		Message: message,
		Msg:     message,
		Data:    data,
	}
	c.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Response().WriteHeader(200)

	if bs, err := json.Marshal(ret); err == nil {
		c.Response().Write(bs)
	}

}
func (c *Context) AjaxError(code int, message string, data any) {
	c.AjaxJson(code, message, data)
}
func (c *Context) AjaxSuccess(message string, data any) {
	c.AjaxJson(0, message, data)
}
func (c *Context) JSON(httpCode int, data any) {
	c.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
	c.w.WriteHeader(httpCode)
	if bs, err := json.Marshal(data); err == nil {
		c.Response().Write(bs)
	}
}
func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

// IsAborted returns true if the current context was aborted.
func (c *Context) IsAborted() bool {
	return c.index >= abortIndex
}
func (c *Context) Abort() {
	c.index = abortIndex
}
func (c *Context) reset(w http.ResponseWriter, r *http.Request) {
	c.r = r
	c.w = w
	c.handlers = c.handlers[:0]
	c.index = 0
	c.queryParsed = false
	c.userData = map[string]any{}
	c.ctx = context.Background()
}
func (c *Context) Error(message string, code int) {
	c.w.WriteHeader(code)
	c.w.Write([]byte(message))
}
func (c *Context) NotFound() {
	c.Error("not found", 404)
}
func (c *Context) Redirect(code int, url string) {
	http.Redirect(c.w, c.r, url, code)
	c.Abort()
	return
}
func (c *Context) SetUserValue(key string, value interface{}) {
	c.userData[key] = value
}
func (c *Context) UserValue(key string) any {
	if v, exist := c.userData[key]; exist {
		return v
	} else {
		return nil
	}
}
func (c *Context) SetDomain(domain string) {
	c.domain = domain
}
func (c *Context) SetCookie(key string, value []byte, expired time.Duration) {
	cookie := &http.Cookie{
		Domain:   c.domain,
		HttpOnly: true,
		Path:     "/",
		Name:     key,
		Expires:  time.Now().Add(expired),
		Secure:   true,
		Value:    string(value),
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(c.w, cookie)
}
func (c *Context) SetCookieV2(cookie *http.Cookie) {
	http.SetCookie(c.w, cookie)
}
func (c *Context) SetAuthorization(s string) {
	c.Response().Header().Set("Authorization", s)
}
func (c *Context) Authorization() string {
	return c.Request().Header.Get("Authorization")
}
func (c *Context) Cookie(key string) string {
	if cookie, err := c.Request().Cookie(key); err == nil {
		return cookie.Value
	} else {
		return ""
	}
}
func (c *Context) initQueryCache() {
	if !c.queryParsed {
		c.queryMap = c.r.URL.Query()
		c.queryParsed = true
	}
}

func (c *Context) Query(key string) string {
	c.initQueryCache()
	return c.queryMap.Get(key)
}
func (c *Context) QueryInt(key string, missing int) int {
	return int(c.QueryInt64(key,int64(missing)))
}
func (c *Context) QueryInt64(key string, missing int64) int64 {
	c.initQueryCache()
	s := c.queryMap.Get(key)
	if len(s) == 0 {
		return missing
	}
	if n, err := strconv.ParseInt(s, 10, 64); err == nil {
		return n
	}
	return missing
}
func (c *Context) QueryArray(key string) []string {
	c.initQueryCache()
	if array, exist := c.queryMap[key]; exist {
		return array
	} else {
		return []string{}
	}
}
func (c *Context) QueryMap() url.Values {
	c.initQueryCache()
	return c.queryMap

}
func (c *Context) PostForm(key string) string {
	return c.r.PostFormValue(key)
}
func (c *Context) AbortWithStatus(code int) {
	c.w.WriteHeader(code)
	c.Abort()
}
