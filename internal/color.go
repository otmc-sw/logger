/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package internal

import (
	"strings"
)

// ANSI color codes
const (
	colorReset      = "\033[0m"
	colorWhiteBg    = "\033[47m"
	colorCyanBg     = "\033[46m"
	colorBlueBg     = "\033[44m"
	colorYellowBg   = "\033[43m"
	colorRedBg      = "\033[41m"
	colorLightRedBg = "\033[101m"
	colorRed        = "\033[31m"
	colorYellow     = "\033[33m"
)

// ColorLevel returns the colored level string
func ColorLevel(level string) string {
	switch level {
	case "TRACE":
		return colorWhiteBg + " TRACE " + colorReset
	case "DEBUG":
		return colorCyanBg + " DEBUG " + colorReset
	case "INFO":
		return colorBlueBg + " INFO  " + colorReset
	case "WARN":
		return colorYellowBg + " WARN  " + colorReset
	case "ERROR":
		return colorRedBg + " ERROR " + colorReset
	case "CRIT":
		return colorLightRedBg + " CRIT  " + colorReset
	default:
		return "UNKNOWN"
	}
}

// StripColorCodes removes ANSI color codes from a string
func StripColorCodes(s string) string {
	s = strings.ReplaceAll(s, colorReset, "")
	s = strings.ReplaceAll(s, colorWhiteBg, "")
	s = strings.ReplaceAll(s, colorCyanBg, "")
	s = strings.ReplaceAll(s, colorBlueBg, "")
	s = strings.ReplaceAll(s, colorYellowBg, "")
	s = strings.ReplaceAll(s, colorRedBg, "")
	s = strings.ReplaceAll(s, colorLightRedBg, "")
	s = strings.ReplaceAll(s, colorRed, "")
	s = strings.ReplaceAll(s, colorYellow, "")
	return s
}

// ColorMessage returns the colored message string based on level
func ColorMessage(level string, message string) string {
	switch level {
	case "WARN":
		return colorYellow + message + colorReset
	case "ERROR":
		return colorRed + message + colorReset
	case "CRIT":
		return colorRed + message + colorReset
	default:
		return message
	}
}