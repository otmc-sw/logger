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

// PrettyFormatter formats log entries in a pretty console format
type PrettyFormatter struct {
	colorize bool
}

// NewPrettyFormatter creates a new pretty formatter
func NewPrettyFormatter(colorize bool) internal.Formatter {
	return &PrettyFormatter{colorize: colorize}
}

// Format formats a log entry
func (f *PrettyFormatter) Format(entry internal.Entry) string {
	timestamp := entry.Time.Format("2006-01-02 15:04:05.000 -07:00")
	levelStr := entry.Level.String()

	if f.colorize {
		levelStr = internal.ColorLevel(levelStr)
	}

	return fmt.Sprintf(
		"%s %20.20s() %15.15s:%-5d |%s| %s\n",
		timestamp,
		entry.Function,
		entry.File,
		entry.Line,
		levelStr,
		entry.Message,
	)
}
