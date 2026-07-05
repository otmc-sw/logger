/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package internal

import (
	"fmt"
	"strings"
)

// ANSI color codes
const (
	// Reset
	ColorReset = "\033[0m"

	// =========================
	// Foreground Colors
	// =========================
	ColorBlack  = "\033[30m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"

	// Bright Foreground Colors
	ColorBrightBlack  = "\033[90m" // Gray
	ColorBrightRed    = "\033[91m"
	ColorBrightGreen  = "\033[92m"
	ColorBrightYellow = "\033[93m"
	ColorBrightBlue   = "\033[94m"
	ColorBrightCyan   = "\033[96m"
	ColorBrightWhite  = "\033[97m"

	// =========================
	// Background Colors
	// =========================
	ColorBlackBg  = "\033[40m"
	ColorRedBg    = "\033[41m"
	ColorGreenBg  = "\033[42m"
	ColorYellowBg = "\033[43m"
	ColorBlueBg   = "\033[44m"
	ColorCyanBg   = "\033[46m"
	ColorWhiteBg  = "\033[47m"

	// Bright Background Colors
	ColorBrightBlackBg  = "\033[100m" // Gray
	ColorBrightRedBg    = "\033[101m"
	ColorBrightGreenBg  = "\033[102m"
	ColorBrightYellowBg = "\033[103m"
	ColorBrightBlueBg   = "\033[104m"
	ColorBrightCyanBg   = "\033[106m"
	ColorBrightWhiteBg  = "\033[107m"
)

// ColorLevel returns the colored level string
func ColorLevel(level string) string {
	switch level {
	case "TRACE":
		return ColorWhiteBg + " TRACE " + ColorReset
	case "DEBUG":
		return ColorCyanBg + " DEBUG " + ColorReset
	case "INFO":
		return ColorBlueBg + " INFO  " + ColorReset
	case "WARN":
		return ColorYellowBg + " WARN  " + ColorReset
	case "ERROR":
		return ColorRedBg + " ERROR " + ColorReset
	case "CRIT":
		return ColorBrightRedBg + " CRIT  " + ColorReset
	default:
		return "UNKNOWN"
	}
}

// StripColorCodes removes ANSI color codes from a string
func StripColorCodes(s string) string {
	s = strings.ReplaceAll(s, ColorReset, "")
	s = strings.ReplaceAll(s, ColorWhiteBg, "")
	s = strings.ReplaceAll(s, ColorCyanBg, "")
	s = strings.ReplaceAll(s, ColorBlueBg, "")
	s = strings.ReplaceAll(s, ColorYellowBg, "")
	s = strings.ReplaceAll(s, ColorRedBg, "")
	s = strings.ReplaceAll(s, ColorBrightRedBg, "")
	s = strings.ReplaceAll(s, ColorRed, "")
	s = strings.ReplaceAll(s, ColorYellow, "")
	s = strings.ReplaceAll(s, ColorGreen, "")
	s = strings.ReplaceAll(s, ColorWhite, "")
	s = strings.ReplaceAll(s, ColorBlack, "")
	return s
}

// ColorMessage returns the colored message string based on level
func ColorMessage(level string, message string) string {
	switch level {
	case "WARN":
		return ColorYellow + message + ColorReset
	case "ERROR":
		return ColorRed + message + ColorReset
	case "CRIT":
		return ColorBrightRed + message + ColorReset
	default:
		return message
	}
}

// ColorMethod returns the colored HTTP method with background color
func ColorMethod(method string) string {
	formattedMethod := fmt.Sprintf("%-7s", method)
	switch method {
	case "GET":
		return ColorBlueBg + formattedMethod + ColorReset
	case "POST":
		return ColorGreenBg + formattedMethod + ColorReset
	case "PUT":
		return ColorYellowBg + formattedMethod + ColorReset
	case "DELETE":
		return ColorRedBg + formattedMethod + ColorReset
	case "PATCH":
		return ColorCyanBg + formattedMethod + ColorReset
	case "OPTIONS":
		return ColorWhiteBg + formattedMethod + ColorReset
	case "HEAD":
		return ColorWhiteBg + formattedMethod + ColorReset
	default:
		return ColorWhiteBg + formattedMethod + ColorReset
	}
}

// ColorStatusCode returns the colored HTTP status code
func ColorStatusCode(code int) string {
	switch {
	case code >= 500:
		return ColorRed + fmt.Sprintf("%3d", code) + ColorReset
	case code >= 400:
		return ColorRed + fmt.Sprintf("%3d", code) + ColorReset
	case code >= 300:
		return ColorYellow + fmt.Sprintf("%3d", code) + ColorReset
	case code >= 200:
		return ColorGreen + fmt.Sprintf("%3d", code) + ColorReset
	default:
		return fmt.Sprintf("%3d", code)
	}
}

// ColorTime returns the colored time string
func ColorTime(timeStr string) string {
	return ColorCyan + fmt.Sprintf("%6s", timeStr) + ColorReset
}

// ColorLatency returns the colored latency string
func ColorLatency(latency string) string {
	return ColorBlue + fmt.Sprintf("%6s", latency) + ColorReset
}
