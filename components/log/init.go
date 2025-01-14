package log

import (
	"fmt"
	log "go.uber.org/zap"
	logcore "go.uber.org/zap/zapcore"
)

type (
	Field         = logcore.Field
	Logger        = log.Logger
	Option        = log.Option
	SugaredLogger = log.SugaredLogger
)

var logger *log.Logger

func InitLogger(out, err []string) {
	logger = GetLogger(out, err, false, true, true)
	defer func() {
		if err := logger.Sync(); err != nil {
			fmt.Println(err.Error())
		}
	}()
	logger.Info("log initialize")
}
func InitCustomerLogger(out, err []string, development, disableStacktrace, disableCaller bool) {
	logger = GetLogger(out, err, development, disableStacktrace, disableCaller)
	defer func() {
		if err := logger.Sync(); err != nil {
			fmt.Println(err.Error())
		}
	}()
	logger.Info("Customer log initialize")
}
func GetLogger(outputPath, errorPath []string, development, disableStacktrace, disableCaller bool) *log.Logger {
	encoderConfig := logcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     logcore.DefaultLineEnding,
		EncodeLevel:    logcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     logcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: logcore.SecondsDurationEncoder,
		EncodeCaller:   logcore.FullCallerEncoder, // 全路径编码器
	}

	// 设置日志级别
	atom := log.NewAtomicLevelAt(log.DebugLevel)
	config := log.Config{
		Level:             atom,
		Development:       development,
		DisableStacktrace: disableStacktrace,
		DisableCaller:     disableCaller,
		Encoding:          "json",
		EncoderConfig:     encoderConfig,
		InitialFields:     map[string]interface{}{},
		OutputPaths:       outputPath,
		ErrorOutputPaths:  errorPath,
	}

	return log.Must(config.Build())
}
