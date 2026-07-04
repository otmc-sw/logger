/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package logger

var global Logger = New()

// Logger is the public logging interface
type Logger interface {
	Trace(format string, args ...any)
	Debug(format string, args ...any)
	Info(format string, args ...any)
	Warn(format string, args ...any)
	Error(format string, args ...any)
	Fatal(format string, args ...any)
	Panic(format string, args ...any)
	Sync() error
}

// stdLogger is the standard logger implementation
type stdLogger struct {
	core *Core
}

// New creates a new logger with the given options
func New(opts ...Option) Logger {
	config := DefaultConfig()
	for _, opt := range opts {
		opt(&config)
	}

	// Create formatter
	var formatter Formatter
	if config.JSON {
		formatter = NewJSONFormatter()
	} else {
		formatter = NewPrettyFormatter(config.Console)
	}

	// Create writer
	var writer Writer
	var writers []Writer

	if config.Console {
		writers = append(writers, NewConsoleWriter(nil))
	}

	if config.File && config.Filename != "" {
		rotateWriter := NewRotateWriter(
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
			writer = NewMultiWriter(writers...)
		}
	}

	core := NewCore(config.Level, config.Caller, formatter, writer)

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
	l.core.Log(TraceLevel, format, args...)
}

// Debug logs a debug message
func (l *stdLogger) Debug(format string, args ...any) {
	l.core.Log(DebugLevel, format, args...)
}

// Info logs an info message
func (l *stdLogger) Info(format string, args ...any) {
	l.core.Log(InfoLevel, format, args...)
}

// Warn logs a warning message
func (l *stdLogger) Warn(format string, args ...any) {
	l.core.Log(WarnLevel, format, args...)
}

// Error logs an error message
func (l *stdLogger) Error(format string, args ...any) {
	l.core.Log(ErrorLevel, format, args...)
}

// Fatal logs a fatal message and exits
func (l *stdLogger) Fatal(format string, args ...any) {
	l.core.Log(FatalLevel, format, args...)
}

// Panic logs a panic message and panics
func (l *stdLogger) Panic(format string, args ...any) {
	l.core.Log(PanicLevel, format, args...)
}

// Sync flushes the logger
func (l *stdLogger) Sync() error {
	return l.core.Sync()
}
