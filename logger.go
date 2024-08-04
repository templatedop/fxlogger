package fxlogger

import (
	//"context"
	"context"
	"fmt"
	//"gotemplate/logger"
	"os"

	"github.com/rs/zerolog"
)

const (
	Service = "service"
)

type Logger struct {
	logger *zerolog.Logger
	//ctx   context.Context
}

func (l *Logger) ToZerolog() *zerolog.Logger {
	return l.logger
}

func (l *Logger) ContextLogger(ctx context.Context) *Logger {

	return &Logger{zerolog.Ctx(ctx)}
}

func (l *Logger) CallerIncluded() *Logger {
	lo := l.logger.With().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + 3).Logger()
	return &Logger{logger: &lo}
	//l.logger = l.logger.With().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + 3).Logger().Level(level)

}

func (l *Logger) Debug(message interface{}, args ...interface{}) {

	l.msg(zerolog.DebugLevel, message, args...)
}

func (l *Logger) Info(message string, args ...interface{}) {

	l.log1(zerolog.InfoLevel, message, args...)
}

// Warn -.
func (l *Logger) Warn(message string, args ...interface{}) {
	l.log1(zerolog.WarnLevel, message, args...)
}

// Error -.
func (l *Logger) Error(message interface{}, args ...interface{}) {
	l.msg(zerolog.ErrorLevel, message, args...)
}

// Fatal -.
func (l *Logger) Fatal(message interface{}, args ...interface{}) {
	l.msg(zerolog.FatalLevel, message, args...)

	os.Exit(1)
}

func (l *Logger) msg(level zerolog.Level, message interface{}, args ...interface{}) {
	lw := l.CallerIncluded()
	switch msg := message.(type) {
	case error:
		lw.log1(level, msg.Error(), args...)
	case string:
		lw.log1(level, msg, args...)
	default:
		lw.log1(zerolog.InfoLevel, fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
	}
}

func (l *Logger) log1(level zerolog.Level, message string, args ...interface{}) {

	loggers := l.logger.WithLevel(level)
	if len(args) == 0 {
		loggers.Msg(message)
	} else {
		loggers.Msgf(message, args...)
	}
}

func (l *Logger) ChainableDebug() *zerolog.Event {
	//debugLogger := l.logger.With().CallerWithSkipFrameCount(0).Logger()
	return l.logger.Debug()
}

func (l *Logger) ChainableInfo() *zerolog.Event {
	//InfoLogger := l.logger.With().CallerWithSkipFrameCount(0).Logger()
	return l.logger.Info()
}

func (l *Logger) ChainableWarn() *zerolog.Event {
	return l.logger.Warn()
}

func (l *Logger) ChainableError() *zerolog.Event {
	return l.logger.Error()
}

func (l *Logger) FromZerolog(logger *zerolog.Logger) *Logger {
	return &Logger{logger}
}