/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package logger

import (
	"time"
)

func Trace(format string, args ...any) {
	global.Trace(format, args...)
}

func Debug(format string, args ...any) {
	global.Debug(format, args...)
}

func Info(format string, args ...any) {
	global.Info(format, args...)
}

func Warn(format string, args ...any) {
	global.Warn(format, args...)
}

func Error(format string, args ...any) {
	global.Error(format, args...)
}

func Crit(format string, args ...any) {
	global.Crit(format, args...)
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