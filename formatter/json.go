/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package formatter

import (
	"encoding/json"
	"time"

	"github.com/otmc-sw/logger/internal"
)

// JSONFormatter formats log entries as JSON
type JSONFormatter struct{}

// NewJSONFormatter creates a new JSON formatter
func NewJSONFormatter() internal.Formatter {
	return &JSONFormatter{}
}

// Format formats a log entry
func (f *JSONFormatter) Format(entry internal.Entry) string {
	data := map[string]interface{}{
		"time":     entry.Time.Format(time.RFC3339Nano),
		"level":    entry.Level.String(),
		"message":  entry.Message,
		"function": entry.Function,
		"file":     entry.File,
		"line":     entry.Line,
	}

	b, _ := json.Marshal(data)
	return string(b) + "\n"
}

// FormatRequest formats an HTTP request entry
func (f *JSONFormatter) FormatRequest(req internal.Request) string {
	data := map[string]interface{}{
		"time":       req.Time.Format(time.RFC3339Nano),
		"method":     req.Method,
		"path":       req.Path,
		"status":     req.StatusCode,
		"latency":    req.Latency.String(),
		"client_ip":  req.ClientIP,
	}

	b, _ := json.Marshal(data)
	return string(b) + "\n"
}
