/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package formatter

import (
	"fmt"

	"github.com/otmc-sw/logger/internal"
)

// TextFormatter formats log entries in simple text format
type TextFormatter struct{}

// NewTextFormatter creates a new text formatter
func NewTextFormatter() internal.Formatter {
	return &TextFormatter{}
}

// Format formats a log entry
func (f *TextFormatter) Format(entry internal.Entry) string {
	return fmt.Sprintf("%s %s\n", entry.Level.String(), entry.Message)
}

// FormatRequest formats an HTTP request entry
func (f *TextFormatter) FormatRequest(req internal.Request) string {
	return fmt.Sprintf("%s %s %d %s\n", req.Method, req.Path, req.StatusCode, req.Latency)
}
