/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package logger

import (
	"github.com/otmc-sw/logger/internal"
)

type Level = internal.Level

const (
	TraceLevel = internal.TraceLevel
	DebugLevel = internal.DebugLevel
	InfoLevel  = internal.InfoLevel
	WarnLevel  = internal.WarnLevel
	ErrorLevel = internal.ErrorLevel
	CritLevel  = internal.CritLevel
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
