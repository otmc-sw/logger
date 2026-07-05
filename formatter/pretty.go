/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package formatter

import (
	"fmt"
	"strconv"

	"github.com/otmc-sw/logger/internal"
)

type PrettyFormatter struct {
	colorize bool
}

func NewPrettyFormatter(colorize bool) internal.Formatter {
	return &PrettyFormatter{colorize: colorize}
}

func (f *PrettyFormatter) Format(entry internal.Entry) string {
	timestamp := entry.Time.Format("2006-01-02 15:04:05.000 -07:00")
	levelStr := entry.Level.String()
	message := entry.Message

	if f.colorize {
		levelStr = internal.ColorLevel(levelStr)
		message = internal.ColorMessage(entry.Level.String(), message)
	}

	return fmt.Sprintf(
		"%s %20.20s() %15.15s:%-5d |%s| %s\n",
		timestamp,
		entry.Function,
		entry.File,
		entry.Line,
		levelStr,
		message,
	)
}

func (f *PrettyFormatter) FormatRequest(req internal.Request) string {
	timestamp := req.Time.Format("15:04:05.000")
	latency := req.Latency.String()
	method := req.Method
	statusCodeStr := strconv.Itoa(req.StatusCode)

	if f.colorize {
		timestamp = internal.ColorTime(timestamp)
		latency = internal.ColorLatency(latency)
		method = internal.ColorMethod(method)
		statusCodeStr = internal.ColorStatusCode(req.StatusCode)
	}

	return fmt.Sprintf(
		"%s |%s| %s | %s | %-15s | %s\n",
		timestamp,
		method,
		latency,
		statusCodeStr,
		req.ClientIP,
		req.Path,
	)
}
