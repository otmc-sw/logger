/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package formatter

import (
	"fmt"

	"github.com/otmc-sw/logger/core"
)

type TextFormatter struct{}

func NewTextFormatter() core.Formatter {
	return &TextFormatter{}
}

func (f *TextFormatter) Format(entry core.Entry) string {
	return fmt.Sprintf("%s %s\n", entry.Level.String(), entry.Message)
}

func (f *TextFormatter) FormatRequest(req core.Request) string {
	return fmt.Sprintf("%s %s %d %s\n", req.Method, req.Path, req.StatusCode, req.Latency)
}
