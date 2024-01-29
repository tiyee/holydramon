package engine

type HttpMethod string

const (
	MethodGet     HttpMethod = "GET"
	MethodHead    HttpMethod = "HEAD"
	MethodPost    HttpMethod = "POST"
	MethodPut     HttpMethod = "PUT"
	MethodPatch   HttpMethod = "PATCH" // RFC 5789
	MethodDelete  HttpMethod = "DELETE"
	MethodConnect HttpMethod = "CONNECT"
	MethodOptions HttpMethod = "OPTIONS"
	MethodTrace   HttpMethod = "TRACE"
)

type Guard struct {
	method      HttpMethod
	Pos         HookPos
	HandlerFunc HandlerFunc
}

func Cors(fn HandlerFunc) Guard {
	return Guard{
		method:      MethodOptions,
		Pos:         PosAhead,
		HandlerFunc: fn,
	}
}
