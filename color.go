/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package logger

import (
	"strings"
)

// ANSI color codes
const (
	colorReset     = "\033[0m"
	colorGray      = "\033[90m"
	colorBlue      = "\033[34m"
	colorGreen     = "\033[32m"
	colorYellow    = "\033[33m"
	colorRed       = "\033[31m"
	colorBrightRed = "\033[91m"
	colorRedBg     = "\033[41m"
)

// colorLevel returns the colored level string
func colorLevel(level string) string {
	switch level {
	case "TRACE":
		return colorGray + " TRACE " + colorReset
	case "DEBUG":
		return colorBlue + " DEBUG " + colorReset
	case "INFO":
		return colorGreen + " INFO  " + colorReset
	case "WARN":
		return colorYellow + " WARN  " + colorReset
	case "ERROR":
		return colorRed + " ERROR " + colorReset
	case "FATAL":
		return colorBrightRed + " FATAL " + colorReset
	case "PANIC":
		return colorRedBg + " PANIC " + colorReset
	default:
		return "UNKNOWN"
	}
}

// stripColorCodes removes ANSI color codes from a string
func stripColorCodes(s string) string {
	s = strings.ReplaceAll(s, colorReset, "")
	s = strings.ReplaceAll(s, colorGray, "")
	s = strings.ReplaceAll(s, colorBlue, "")
	s = strings.ReplaceAll(s, colorGreen, "")
	s = strings.ReplaceAll(s, colorYellow, "")
	s = strings.ReplaceAll(s, colorRed, "")
	s = strings.ReplaceAll(s, colorBrightRed, "")
	s = strings.ReplaceAll(s, colorRedBg, "")
	return s
}
