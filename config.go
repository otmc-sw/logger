/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package logger

// Config represents the logger configuration
type Config struct {
	// Level is the minimum log level to output
	Level Level

	// Console enables console output
	Console bool

	// File enables file output
	File bool

	// Filename is the path to the log file
	Filename string

	// JSON enables JSON formatting
	JSON bool

	// Caller enables caller information (function, file, line)
	Caller bool

	// Stacktrace enables stacktrace for errors
	Stacktrace bool

	// MaxSize is the maximum size in megabytes of a log file before it gets rotated
	MaxSize int

	// MaxBackups is the maximum number of old log files to retain
	MaxBackups int

	// MaxAge is the maximum number of days to retain old log files
	MaxAge int

	// Compress determines if the rotated log files should be compressed using gzip
	Compress bool
}

// DefaultConfig returns the default configuration
func DefaultConfig() Config {
	return Config{
		Level:      InfoLevel,
		Console:    true,
		File:       false,
		JSON:       false,
		Caller:     true,
		Stacktrace: false,
		MaxSize:    100, // 100 MB
		MaxBackups: 3,
		MaxAge:     30, // 30 days
		Compress:   true,
	}
}
