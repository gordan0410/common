package log

import (
	"context"
	"github.com/gordan0410/common/enum"
	"os"
	"runtime/debug"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type LogField struct {
	ProjectName string
	Version     string
	Level       int
}

var (
	once         sync.Once
	singleLogger *CommonLogger
)

type CommonLogger struct {
	zerolog.Logger
}

// NewLogger 初始化 logger，只會執行一次
func NewLogger(field LogField) *CommonLogger {
	once.Do(func() {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
		level := zerolog.Level(field.Level)
		logger := log.Level(level).Output(os.Stdout).With().Caller().Str("projectName", field.ProjectName).Str("version", field.Version).Logger()
		singleLogger = &CommonLogger{logger}
	})
	return singleLogger
}

func (commonLogger CommonLogger) WithCtx(ctx context.Context) *CommonLogger {
	tid, _ := ctx.Value(enum.TraceId.ToString()).(string)
	stid, _ := ctx.Value(enum.SubTraceId.ToString()).(string)
	return &CommonLogger{commonLogger.With().Str(enum.TraceId.ToString(), tid).Str(enum.SubTraceId.ToString(), stid).Logger()}
}

func (commonLogger CommonLogger) Error() *zerolog.Event {
	return commonLogger.Logger.Error().Bytes("stack_trace", debug.Stack())
}

func (commonLogger CommonLogger) Panic() *zerolog.Event {
	return commonLogger.Logger.Panic().Bytes("stack_trace", debug.Stack())
}

// 以下為 logger 的方法，這些方法會使用 singleton logger

func WithCtx(ctx context.Context) *CommonLogger {
	NewLogger(LogField{ProjectName: "", Version: ""})

	tid, _ := ctx.Value(enum.TraceId.ToString()).(string)
	stid, _ := ctx.Value(enum.SubTraceId.ToString()).(string)
	logger := singleLogger.With().Str(enum.TraceId.ToString(), tid).Str(enum.SubTraceId.ToString(), stid).Logger()
	return &CommonLogger{logger}
}

func With() zerolog.Context {
	NewLogger(LogField{ProjectName: "", Version: ""})
	return singleLogger.With()
}

func Info() *zerolog.Event {
	NewLogger(LogField{ProjectName: "", Version: ""})
	return singleLogger.Info()
}

func Error() *zerolog.Event {
	NewLogger(LogField{ProjectName: "", Version: ""})
	return singleLogger.Error().Bytes("stack_trace", debug.Stack())
}

func Warn() *zerolog.Event {
	NewLogger(LogField{ProjectName: "", Version: ""})
	return singleLogger.Warn()
}

func Debug() *zerolog.Event {
	NewLogger(LogField{ProjectName: "", Version: ""})
	return singleLogger.Debug()
}

func Trace() *zerolog.Event {
	NewLogger(LogField{ProjectName: "", Version: ""})
	return singleLogger.Trace()
}

func Panic() *zerolog.Event {
	NewLogger(LogField{ProjectName: "", Version: ""})

	return singleLogger.Panic().Bytes("stack_trace", debug.Stack())
}
