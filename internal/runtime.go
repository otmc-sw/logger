/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package internal

import (
	"path/filepath"
	"runtime"
	"strings"
)

// CallerInfo holds information about the caller
type CallerInfo struct {
	Function string
	File     string
	Line     int
}

// GetCaller retrieves caller information from the runtime
func GetCaller(skip int) CallerInfo {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return CallerInfo{
			Function: "Unknown",
			File:     "Unknown",
			Line:     0,
		}
	}

	fn := "Unknown"
	if f := runtime.FuncForPC(pc); f != nil {
		parts := strings.Split(f.Name(), ".")
		fn = parts[len(parts)-1]
	}

	return CallerInfo{
		Function: fn,
		File:     filepath.Base(file),
		Line:     line,
	}
}
