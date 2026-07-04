/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package logger

import (
	"github.com/otmc-sw/logger/formatter"
	"github.com/otmc-sw/logger/internal"
)

var global Logger = New()

// Logger is the public logging interface
type Logger interface {
	Trace(format string, args ...any)
	Debug(format string, args ...any)
	Info(format string, args ...any)
	Warn(format string, args ...any)
	Error(format string, args ...any)
	Crit(format string, args ...any)
	Sync() error
}

// stdLogger is the standard logger implementation
type stdLogger struct {
	core *internal.Core
}

// New creates a new logger with the given options
func New(opts ...Option) Logger {
	config := DefaultConfig()
	for _, opt := range opts {
		opt(&config)
	}

	// Create formatter
	var fmt internal.Formatter
	if config.JSON {
		fmt = formatter.NewJSONFormatter()
	} else {
		fmt = formatter.NewPrettyFormatter(config.Console)
	}

	// Create writer
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

// Init initializes the global logger
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

// Trace logs a trace message
func (l *stdLogger) Trace(format string, args ...any) {
	l.core.Log(internal.TraceLevel, format, args...)
}

// Debug logs a debug message
func (l *stdLogger) Debug(format string, args ...any) {
	l.core.Log(internal.DebugLevel, format, args...)
}

// Info logs an info message
func (l *stdLogger) Info(format string, args ...any) {
	l.core.Log(internal.InfoLevel, format, args...)
}

// Warn logs a warning message
func (l *stdLogger) Warn(format string, args ...any) {
	l.core.Log(internal.WarnLevel, format, args...)
}

// Error logs an error message
func (l *stdLogger) Error(format string, args ...any) {
	l.core.Log(internal.ErrorLevel, format, args...)
}

// Crit logs a critical message and exits
func (l *stdLogger) Crit(format string, args ...any) {
	l.core.Log(internal.CritLevel, format, args...)
}

// Sync flushes the logger
func (l *stdLogger) Sync() error {
	return l.core.Sync()
}
