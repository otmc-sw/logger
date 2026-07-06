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
	cfg := DefaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	var fmt internal.Formatter
	if cfg.JSON {
		fmt = formatter.NewJSONFormatter(cfg.TimeFormat)
	} else {
		fmt = formatter.NewPrettyFormatter(cfg.Console, cfg.TimeFormat)
	}

	var writers []internal.Writer
	if cfg.Console {
		writers = append(writers, internal.NewConsoleWriter(nil))
	}
	if cfg.File && cfg.Filename != "" {
		writers = append(writers, internal.NewRotateWriter(
			cfg.Filename, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge, cfg.Compress,
		))
	}

	var writer internal.Writer
	if len(writers) > 0 {
		writer = internal.NewMultiWriter(writers...)
	}

	return &stdLogger{core: internal.NewCore(cfg.Level, cfg.Caller, fmt, writer)}
}

func (l *stdLogger) Trace(f string, a ...any) { l.core.Log(internal.TraceLevel, 4, f, a...) }
func (l *stdLogger) Debug(f string, a ...any) { l.core.Log(internal.DebugLevel, 4, f, a...) }
func (l *stdLogger) Info(f string, a ...any)  { l.core.Log(internal.InfoLevel, 4, f, a...) }
func (l *stdLogger) Warn(f string, a ...any)  { l.core.Log(internal.WarnLevel, 4, f, a...) }
func (l *stdLogger) Error(f string, a ...any) { l.core.Log(internal.ErrorLevel, 4, f, a...) }
func (l *stdLogger) Crit(f string, a ...any)  { l.core.Log(internal.CritLevel, 4, f, a...) }

func (l *stdLogger) Request(method, path string, statusCode int, latency time.Duration, clientIP string) {
	l.core.LogRequest(internal.Request{
		Time:       time.Now(),
		Method:     method,
		Path:       path,
		StatusCode: statusCode,
		Latency:    latency,
		ClientIP:   clientIP,
	})
}

func (l *stdLogger) Sync() error { return l.core.Sync() }
