/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package logger

import (
	"time"

	"github.com/otmc-sw/logger/formatter"
	"github.com/otmc-sw/logger/internal"
)

var global Logger = New()

type Logger interface {
	Trace(format string, args ...any)
	Debug(format string, args ...any)
	Info(format string, args ...any)
	Warn(format string, args ...any)
	Error(format string, args ...any)
	Crit(format string, args ...any)
	Request(method, path string, statusCode int, latency time.Duration, clientIP string)
	Sync() error
}

type stdLogger struct {
	core *internal.Core
}

func New(opts ...Option) Logger {
	config := DefaultConfig()
	for _, opt := range opts {
		opt(&config)
	}

	var fmt internal.Formatter
	if config.JSON {
		fmt = formatter.NewJSONFormatter()
	} else {
		fmt = formatter.NewPrettyFormatter(config.Console)
	}

	var writer internal.Writer
	var writers []internal.Writer

	if config.Console {
		writers = append(writers, internal.NewConsoleWriter(nil))
	}

	if config.File && config.Filename != "" {
		rotateWriter := internal.NewRotateWriter(
			config.Filename,
			config.MaxSize,
			config.MaxBackups,
			config.MaxAge,
			config.Compress,
		)
		writers = append(writers, rotateWriter)
	}

	if len(writers) > 0 {
		if len(writers) == 1 {
			writer = writers[0]
		} else {
			writer = internal.NewMultiWriter(writers...)
		}
	}

	core := internal.NewCore(config.Level, config.Caller, fmt, writer)

	return &stdLogger{core: core}
}

func Init(config Config) {
	global = New(
		WithLevel(config.Level),
		WithConsole(config.Console),
		WithFile(config.Filename),
		WithJSON(config.JSON),
		WithCaller(config.Caller),
		WithMaxSize(config.MaxSize),
		WithMaxBackups(config.MaxBackups),
		WithMaxAge(config.MaxAge),
		WithCompress(config.Compress),
	)
}

func (l *stdLogger) Trace(format string, args ...any) {
	l.core.Log(internal.TraceLevel, 4, format, args...)
}

func (l *stdLogger) Debug(format string, args ...any) {
	l.core.Log(internal.DebugLevel, 4, format, args...)
}

func (l *stdLogger) Info(format string, args ...any) {
	l.core.Log(internal.InfoLevel, 4, format, args...)
}

func (l *stdLogger) Warn(format string, args ...any) {
	l.core.Log(internal.WarnLevel, 4, format, args...)
}

func (l *stdLogger) Error(format string, args ...any) {
	l.core.Log(internal.ErrorLevel, 4, format, args...)
}

func (l *stdLogger) Crit(format string, args ...any) {
	l.core.Log(internal.CritLevel, 4, format, args...)
}

func (l *stdLogger) Request(method, path string, statusCode int, latency time.Duration, clientIP string) {
	req := internal.Request{
		Time:       time.Now(),
		Method:     method,
		Path:       path,
		StatusCode: statusCode,
		Latency:    latency,
		ClientIP:   clientIP,
	}

	l.core.LogRequest(req)
}

func (l *stdLogger) Sync() error {
	return l.core.Sync()
}
