package engine

import (
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var ImmutableConfig *Config
var ImmutableComponents *Components

type HandlerFunc func(ctx *Context)
type Route struct {
	Method      string
	Path        string
	BeforeHooks []HandlerFunc
	AfterHooks  []HandlerFunc
	HandlerFunc HandlerFunc
}
type Routes map[string]*Route
type Engine struct {
	routes     Routes
	context    *Context
	logger     *zap.Logger
	config     *Config
	hooks      []*Hook
	components *Components
}

func New() *Engine {
	return &Engine{
		routes: make(Routes, 0),
	}
}
func (e *Engine) InitConfig(tomlPath *string) error {
	var config Config
	if _, err := toml.DecodeFile(*tomlPath, &config); err == nil {
		e.config = &config
		ImmutableConfig = &config
		return nil
	} else {
		return err
	}

}
func (e *Engine) InitLogger() error {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
	}

	// 设置日志级别
	atom := zap.NewAtomicLevelAt(zap.DebugLevel)

	config := zap.Config{
		Level:             atom,
		Development:       true,
		DisableStacktrace: true,
		DisableCaller:     false,
		Encoding:          "json",
		EncoderConfig:     encoderConfig,
		InitialFields:     map[string]interface{}{"serviceName": "holydramon"},
		OutputPaths:       []string{e.Config().System.LogPath},
		ErrorOutputPaths:  []string{"./logs/error.log"},
	}
	if logger, err := config.Build(); err == nil {
		ticker := time.NewTicker(time.Second * 10)
		go func() {
			for _ = range ticker.C {
				//fmt.Printf("ticked at %v \n", time.Now())
				if err := logger.Sync(); err != nil {
					fmt.Println(err.Error())
					break
				}
			}
		}()
		e.logger = logger
		return nil
	} else {
		return err
	}

}
func (e *Engine) Logger() *zap.Logger {
	return e.logger
}
func (e *Engine) Config() *Config {
	return e.config
}
func (e *Engine) Components() *Components {
	return e.components
}
func (e *Engine) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	methodS := string(ctx.Method())
	pathS := string(ctx.Path())
	key := methodS + pathS
	if route, exist := e.routes[key]; exist {
		context := Context{
			RequestCtx: ctx,
			engine:     e,
			state:      0,
		}
		for _, fn := range route.BeforeHooks {
			if context.state == CtxStateContinue {
				fn(&context)
			}
		}
		if context.state == CtxStateContinue {
			route.HandlerFunc(&context)
		}
		for _, fn := range route.AfterHooks {
			if context.state == CtxStateContinue {
				fn(&context)
			}
		}

	} else {
		ctx.NotFound()
	}
}

func (e *Engine) setRoute(method string, path string, fn HandlerFunc) {
	key := method + path
	e.routes[key] = &Route{
		Method:      method,
		Path:        path,
		HandlerFunc: fn,
		BeforeHooks: make([]HandlerFunc, 0),
		AfterHooks:  make([]HandlerFunc, 0),
	}
}
func (e *Engine) Run() (err error) {
	e.dispatch()
	return fasthttp.ListenAndServe(e.Config().System.Addr, e.HandleFastHTTP)
}
func (e *Engine) GET(path string, fn HandlerFunc) {
	e.setRoute("GET", path, fn)
}
func (e *Engine) POST(path string, fn HandlerFunc) {
	e.setRoute("POST", path, fn)
}
func (e *Engine) AddHook(hook *Hook) {
	e.hooks = append(e.hooks, hook)
}
func (e *Engine) dispatch() {
	for _, hk := range e.hooks {
		for _, rt := range e.routes {
			if hk.Match(rt) {
				if hk.Pos == PosBefore {
					rt.BeforeHooks = append(rt.BeforeHooks, hk.HandlerFunc)
				}
				if hk.Pos == PosAfter {
					rt.AfterHooks = append(rt.AfterHooks, hk.HandlerFunc)
				}
			}
		}
	}
}
