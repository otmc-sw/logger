/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package logger

import (
	"fmt"
)

// TextFormatter formats log entries in simple text format
type TextFormatter struct{}

// NewTextFormatter creates a new text formatter
func NewTextFormatter() Formatter {
	return &TextFormatter{}
}

// Format formats a log entry
func (f *TextFormatter) Format(entry Entry) string {
	return fmt.Sprintf("%s %s\n", entry.Level.String(), entry.Message)
}
