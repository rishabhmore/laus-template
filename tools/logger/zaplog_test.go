package logger

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestInfo(t *testing.T) {
	// define our args struct
	type args struct {
		ctx context.Context
		msg string
	}
	tests := map[string]args{
		"test info": {
			ctx: context.Background(),
			msg: "This is an info log",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			observedZapCore, observedLogs := observer.New(zap.InfoLevel)
			observedLogger := zap.New(observedZapCore).Sugar()
			_ = SetLogger(observedLogger)
			Info(tt.ctx, tt.msg)
			assert.Equal(t, 1, observedLogs.Len())
			log := observedLogs.All()[0]
			assert.Equal(t, fmt.Sprintf("\nRequest-ID: <nil>\n[%s]\n", tt.msg), log.Message)
			assert.Equal(t, zapcore.Level(0), log.Level)
		})
	}
}

func TestDebug(t *testing.T) {
	// define our args struct
	type args struct {
		ctx context.Context
		msg string
	}
	tests := map[string]args{
		"test debug": {
			ctx: context.Background(),
			msg: "This is a debug log",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			observedZapCore, observedLogs := observer.New(zap.DebugLevel)
			observedLogger := zap.New(observedZapCore).Sugar()
			_ = SetLogger(observedLogger)
			Debug(tt.ctx, tt.msg)
			assert.Equal(t, 1, observedLogs.Len())
			log := observedLogs.All()[0]
			assert.Equal(t, fmt.Sprintf("\nRequest-ID: <nil>\n[%s]\n", tt.msg), log.Message)
			assert.Equal(t, zapcore.Level(-1), log.Level)
		})
	}
}

func TestError(t *testing.T) {
	// define our args struct
	type args struct {
		ctx context.Context
		msg string
	}

	tests := map[string]struct {
		message string
		args    args
	}{
		"test error": {
			message: "This is an error log",
			args: args{
				ctx: context.Background(),
				msg: "some args",
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			observedZapCore, observedLogs := observer.New(zap.ErrorLevel)
			observedLogger := zap.New(observedZapCore).Sugar()
			_ = SetLogger(observedLogger)

			// Test our zapper error method
			Error(tt.args.ctx, tt.message, tt.args.msg)
			assert.Equal(t, 1, observedLogs.Len())

			log := observedLogs.All()[0]
			// validate if our error message logged is the same as we expect
			assert.Equal(t, fmt.Sprintf("\nRequest-ID: <nil>\n%s\n[%s]\n", tt.message, tt.args.msg), log.Message)
			assert.Equal(t, zapcore.ErrorLevel, log.Level)
		})
	}
}

func Test_customLevelEncoder(t *testing.T) {
	tests := map[string]struct {
		level zapcore.Level
		want  interface{} // output of encoding InfoLevel
	}{
		"TestDebugLevel":      {zapcore.DebugLevel, "DEBUG"},
		"TestInfoLevel":       {zapcore.InfoLevel, "INFO"},
		"TestWarnLevel":       {zapcore.WarnLevel, "WARNING"},
		"TestErrorLevel":      {zapcore.ErrorLevel, "ERROR"},
		"TestDebugPanicLevel": {zapcore.DPanicLevel, "CRITICAL"},
		"TestPanicLevel":      {zapcore.PanicLevel, "ALERT"},
		"TestFatalLevel":      {zapcore.FatalLevel, "EMERGENCY"},
	}

	for name, tt := range tests {
		assertAppended(
			t,
			tt.want,
			func(arr zapcore.ArrayEncoder) { customLevelEncoder(tt.level, arr) },
			"Unexpected output serializing InfoLevel with %q.", name,
		)
	}
}

func assertAppended(t testing.TB, expected interface{}, f func(zapcore.ArrayEncoder), msgAndArgs ...interface{}) {
	mem := zapcore.NewMapObjectEncoder()
	err := mem.AddArray("k", zapcore.ArrayMarshalerFunc(func(arr zapcore.ArrayEncoder) error {
		f(arr)
		return nil
	}))
	if err != nil {
		fmt.Printf("assertAppended:: Failed to add to Array")
	}
	arr := mem.Fields["k"].([]interface{})
	assert.Equal(t, 1, len(arr), "Expected to append exactly one element to array.")
	assert.Equal(t, expected, arr[0], msgAndArgs...)
}
