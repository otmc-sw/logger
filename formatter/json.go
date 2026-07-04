/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package formatter

import (
	"encoding/json"
	"time"

	"github.com/otmc-sw/logger"
)

// JSONFormatter formats log entries as JSON
type JSONFormatter struct{}

// NewJSONFormatter creates a new JSON formatter
func NewJSONFormatter() *JSONFormatter {
	return &JSONFormatter{}
}

// Format formats a log entry
func (f *JSONFormatter) Format(entry logger.Entry) string {
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
