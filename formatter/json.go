/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package formatter

import (
	"encoding/json"
	"time"

	"github.com/otmc-sw/logger/core"
)

type JSONFormatter struct {
	timeFormat string
}

func NewJSONFormatter(timeFormat string) core.Formatter {
	return &JSONFormatter{timeFormat: timeFormat}
}

func (f *JSONFormatter) Format(entry core.Entry) string {
	timeFormat := f.timeFormat
	if timeFormat == "" {
		timeFormat = time.RFC3339Nano
	}
	data := map[string]interface{}{
		"time":     entry.Time.Format(timeFormat),
		"level":    entry.Level.String(),
		"message":  entry.Message,
		"function": entry.Function,
		"file":     entry.File,
		"line":     entry.Line,
	}

	b, _ := json.Marshal(data)
	return string(b) + "\n"
}

func (f *JSONFormatter) FormatRequest(req core.Request) string {
	timeFormat := f.timeFormat
	if timeFormat == "" {
		timeFormat = time.RFC3339Nano
	}
	data := map[string]interface{}{
		"time":      req.Time.Format(timeFormat),
		"method":    req.Method,
		"path":      req.Path,
		"status":    req.StatusCode,
		"latency":   req.Latency.String(),
		"client_ip": req.ClientIP,
	}

	b, _ := json.Marshal(data)
	return string(b) + "\n"
}
