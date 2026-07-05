/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package logger

import (
	"time"
)

// Trace logs a trace message using the global logger
func Trace(format string, args ...any) {
	global.Trace(format, args...)
}

// Debug logs a debug message using the global logger
func Debug(format string, args ...any) {
	global.Debug(format, args...)
}

// Info logs an info message using the global logger
func Info(format string, args ...any) {
	global.Info(format, args...)
}

// Warn logs a warning message using the global logger
func Warn(format string, args ...any) {
	global.Warn(format, args...)
}

// Error logs an error message using the global logger
func Error(format string, args ...any) {
	global.Error(format, args...)
}

// Crit logs a critical message using the global logger and exits
func Crit(format string, args ...any) {
	global.Crit(format, args...)
}

// Request logs an HTTP request using the global logger
func Request(method, path string, statusCode int, latency time.Duration, clientIP string) {
	global.Request(method, path, statusCode, latency, clientIP)
}

// Sync flushes the global logger
func Sync() error {
	return global.Sync()
}

// SetLevel sets the log level of the global logger
func SetLevel(level Level) {
	if l, ok := global.(*stdLogger); ok {
		l.core.SetLevel(level)
	}
}
