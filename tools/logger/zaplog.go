package logger

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Key struct to hold the value for Key ID
type ContextKey struct {
	Name string
}

// The Key ID for the Request ID provided from a http request
var RequestIdCtxKey = &ContextKey{echo.HeaderXRequestID}

// Internal method to get Request Id from a context
func getRequestID(c context.Context) string {
	return fmt.Sprintf("%v", c.Value(RequestIdCtxKey))
}

// Internal method to get a formatted string of a Request Id from a context
func getFormattedRequestID(c context.Context) string {
	return fmt.Sprintf("\nRequest-ID: %s", getRequestID(c))
}

// A global instance of our zaplog with custom encoding configs
var Logger = initializeLogger()

// Initialize our Zap Logger here
func initializeLogger() *zap.SugaredLogger {
	loggerCfg := &zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Encoding:         "json",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	plain, err := loggerCfg.Build(zap.AddStacktrace(zap.DPanicLevel))
	if err != nil {
		// If we failed to build a logger instance according to our config
		// then we will simply return a no-operation logger instance
		plain = zap.NewNop()
	}
	// otherwise, return a sugared logger
	return plain.Sugar()
}

// Sets a sugared logger to the current global logger instance
func SetLogger(logger *zap.SugaredLogger) *zap.SugaredLogger {
	Logger = logger
	return Logger
}

func Info(c context.Context, message string, args ...interface{}) {
	Logger.Info(getFormattedRequestID(c), "\n", message, "\n", args, "\n")
}

func Debug(c context.Context, message string, args ...interface{}) {
	Logger.Debug(getFormattedRequestID(c), "\n", message, "\n", args, "\n")
}

func Error(c context.Context, message string, args ...interface{}) {
	Logger.Error(getFormattedRequestID(c), "\n", message, "\n", args, "\n")
}

// The logger encoder config for zapcore
var encoderConfig = zapcore.EncoderConfig{
	TimeKey:        "time",
	LevelKey:       "severity",
	NameKey:        "logger",
	CallerKey:      "caller",
	MessageKey:     "message",
	StacktraceKey:  "stacktrace",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    customLevelEncoder,
	EncodeTime:     zapcore.RFC3339TimeEncoder,
	EncodeDuration: zapcore.MillisDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

// A custom log level encoder to provide log responses
func customLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch l {
	case zapcore.DebugLevel:
		enc.AppendString("DEBUG")
	case zapcore.InfoLevel:
		enc.AppendString("INFO")
	case zapcore.WarnLevel:
		enc.AppendString("WARNING")
	case zapcore.ErrorLevel:
		enc.AppendString("ERROR")
	case zapcore.DPanicLevel:
		enc.AppendString("CRITICAL")
	case zapcore.PanicLevel:
		enc.AppendString("ALERT")
	case zapcore.FatalLevel:
		enc.AppendString("EMERGENCY")
	}
}
