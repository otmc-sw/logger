/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package logger

import (
	"github.com/otmc-sw/logger/internal"
)

// Level is the log level type
type Level = internal.Level

const (
	TraceLevel = internal.TraceLevel
	DebugLevel = internal.DebugLevel
	InfoLevel  = internal.InfoLevel
	WarnLevel  = internal.WarnLevel
	ErrorLevel = internal.ErrorLevel
	FatalLevel = internal.FatalLevel
	PanicLevel = internal.PanicLevel
)

// ParseLevel parses a string level and returns the Level constant
func ParseLevel(level string) Level {
	switch level {
	case "TRACE":
		return TraceLevel
	case "DEBUG":
		return DebugLevel
	case "INFO":
		return InfoLevel
	case "WARN":
		return WarnLevel
	case "ERROR":
		return ErrorLevel
	case "FATAL":
		return FatalLevel
	case "PANIC":
		return PanicLevel
	default:
		return InfoLevel
	}
}
