/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package formatter

import (
	"encoding/json"

	"github.com/otmc-sw/logger/internal"
)

type JSONFormatter struct {
	timeFormat string
}

func NewJSONFormatter(timeFormat string) internal.Formatter {
	return &JSONFormatter{timeFormat: timeFormat}
}

func (f *JSONFormatter) Format(entry internal.Entry) string {
	data := map[string]interface{}{
		"time":     entry.Time.Format(f.timeFormat),
		"level":    entry.Level.String(),
		"message":  entry.Message,
		"function": entry.Function,
		"file":     entry.File,
		"line":     entry.Line,
	}

	b, _ := json.Marshal(data)
	return string(b) + "\n"
}

func (f *JSONFormatter) FormatRequest(req internal.Request) string {
	data := map[string]interface{}{
		"time":      req.Time.Format(f.timeFormat),
		"method":    req.Method,
		"path":      req.Path,
		"status":    req.StatusCode,
		"latency":   req.Latency.String(),
		"client_ip": req.ClientIP,
	}

	b, _ := json.Marshal(data)
	return string(b) + "\n"
}
