/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package formatter

import (
	"fmt"
	"strconv"

	"github.com/otmc-sw/logger/core"
)

type PrettyFormatter struct {
	colorize   bool
	timeFormat string
}

func NewPrettyFormatter(colorize bool, timeFormat string) core.Formatter {
	return &PrettyFormatter{colorize: colorize, timeFormat: timeFormat}
}

func (f *PrettyFormatter) Format(entry core.Entry) string {
	timeFormat := f.timeFormat
	if timeFormat == "" {
		timeFormat = "2006-01-02 15:04:05.000 -07:00"
	}
	timestamp := entry.Time.Format(timeFormat)
	levelStr := entry.Level.String()
	message := entry.Message

	if f.colorize {
		levelStr = core.ColorLevel(levelStr)
		message = core.ColorMessage(entry.Level.String(), message)
	}

	var formatted string
	if entry.Function != "" || entry.File != "" || entry.Line != 0 {
		formatted = fmt.Sprintf(
			"%s %20.20s() %15.15s:%-5d |%s| %s",
			timestamp,
			entry.Function,
			entry.File,
			entry.Line,
			levelStr,
			message,
		)
	} else {
		formatted = fmt.Sprintf(
			"%s |%s| %s",
			timestamp,
			levelStr,
			message,
		)
	}

	return formatted + "\n"
}

func (f *PrettyFormatter) FormatRequest(req core.Request) string {
	timeFormat := f.timeFormat
	if timeFormat == "" {
		timeFormat = "15:04:05.000"
	}
	timestamp := req.Time.Format(timeFormat)
	latency := req.Latency.String()
	method := req.Method
	statusCodeStr := strconv.Itoa(req.StatusCode)

	if f.colorize {
		timestamp = core.ColorTime(timestamp)
		latency = core.ColorLatency(latency)
		method = core.ColorMethod(method)
		statusCodeStr = core.ColorStatusCode(req.StatusCode)
	}

	return fmt.Sprintf(
		"%s |%s| %12s | %s | %-15s | %s\n",
		timestamp,
		method,
		latency,
		statusCodeStr,
		req.ClientIP,
		req.Path,
	)
}
