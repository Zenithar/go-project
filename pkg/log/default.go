package log

import (
	"context"
	l "log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var defaultFactory Factory

// -----------------------------------------------------------------------------

func init() {
	defaultLogger, err := zap.NewProduction()
	if err != nil {
		l.Fatalln(err)
	}

	defaultFactory = NewFactory(defaultLogger)
}

// SetLogger defines the default package logger
func SetLogger(instance Factory) {
	defaultFactory = instance
}

// -----------------------------------------------------------------------------

// Bg delegates a no-context logger
func Bg() Logger {
	return defaultFactory.Bg()
}

// For delegates a context logger
func For(ctx context.Context) Logger {
	return defaultFactory.For(ctx)
}

// With return a sub logger factory
func With(fields ...zapcore.Field) Factory {
	return defaultFactory.With(fields...)
}
