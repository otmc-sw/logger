/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package logger

import (
	"time"
)

// Level represents the severity level of a log entry
type Level int

const (
	// TraceLevel is the most verbose level
	TraceLevel Level = iota
	// DebugLevel is for debug information
	DebugLevel
	// InfoLevel is for general informational messages
	InfoLevel
	// WarnLevel is for warning messages
	WarnLevel
	// ErrorLevel is for error messages
	ErrorLevel
	// FatalLevel is for fatal errors (exits after logging)
	FatalLevel
	// PanicLevel is for panic conditions (panics after logging)
	PanicLevel
)

// String returns the string representation of the level
func (l Level) String() string {
	switch l {
	case TraceLevel:
		return "TRACE"
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	case PanicLevel:
		return "PANIC"
	default:
		return "UNKNOWN"
	}
}

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

// Entry represents a log entry
type Entry struct {
	Time     time.Time
	Level    Level
	Function string
	File     string
	Line     int
	Message  string
}
