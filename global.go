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

func Close() error {
	return global.Close()
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
