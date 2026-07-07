/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package logger

import (
	"time"

	"github.com/otmc-sw/logger/core"
)

type Formatter = core.Formatter
type Hook = core.Hook

type Level = core.Level

const (
	TraceLevel = core.TraceLevel
	DebugLevel = core.DebugLevel
	InfoLevel  = core.InfoLevel
	WarnLevel  = core.WarnLevel
	ErrorLevel = core.ErrorLevel
	CritLevel  = core.CritLevel
)

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
	case "CRIT":
		return CritLevel
	default:
		return InfoLevel
	}
}

type Config struct {
	Level      Level
	Console    bool
	File       bool
	Filename   string
	JSON       bool
	Caller     bool
	Stacktrace bool
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
	TimeFormat string
}

func DefaultConfig() Config {
	return Config{
		Level:      InfoLevel,
		Console:    true,
		File:       false,
		JSON:       false,
		Caller:     true,
		Stacktrace: false,
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     90,
		Compress:   true,
		TimeFormat: "2006-01-02 15:04:05.000 -07:00",
	}
}

func Trace(format string, args ...any) {
	global.core.Log(core.TraceLevel, 3, format, args...)
}

func Debug(format string, args ...any) {
	global.core.Log(core.DebugLevel, 3, format, args...)
}

func Info(format string, args ...any) {
	global.core.Log(core.InfoLevel, 3, format, args...)
}

func Warn(format string, args ...any) {
	global.core.Log(core.WarnLevel, 3, format, args...)
}

func Error(format string, args ...any) {
	global.core.Log(core.ErrorLevel, 3, format, args...)
}

func Crit(format string, args ...any) {
	global.core.Log(core.CritLevel, 3, format, args...)
}

func Request(method, path string, statusCode int, latency time.Duration, clientIP string) {
	global.Request(method, path, statusCode, latency, clientIP)
}

func Sync() error {
	return global.Sync()
}

func SetLevel(level Level) {
	global.SetLevel(level)
}

func Configure(opts ...Option) {
	global.Configure(opts...)
}

func GetConfig() Config {
	return global.Config()
}

func Update(cfg Config) {
	global.Update(cfg)
}
