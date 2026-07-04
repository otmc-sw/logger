/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package logger

// Option is a function that configures the logger
type Option func(*Config)

// WithLevel sets the log level
func WithLevel(level Level) Option {
	return func(c *Config) {
		c.Level = level
	}
}

// WithConsole enables or disables console output
func WithConsole(enabled bool) Option {
	return func(c *Config) {
		c.Console = enabled
	}
}

// WithFile enables file output with the specified filename
func WithFile(filename string) Option {
	return func(c *Config) {
		c.File = true
		c.Filename = filename
	}
}

// WithJSON enables JSON formatting
func WithJSON(enabled bool) Option {
	return func(c *Config) {
		c.JSON = enabled
	}
}

// WithCaller enables or disables caller information
func WithCaller(enabled bool) Option {
	return func(c *Config) {
		c.Caller = enabled
	}
}

// WithStacktrace enables or disables stacktrace for errors
func WithStacktrace(enabled bool) Option {
	return func(c *Config) {
		c.Stacktrace = enabled
	}
}

// WithMaxSize sets the maximum size in megabytes before rotation
func WithMaxSize(maxSize int) Option {
	return func(c *Config) {
		c.MaxSize = maxSize
	}
}

// WithMaxBackups sets the maximum number of old log files to retain
func WithMaxBackups(maxBackups int) Option {
	return func(c *Config) {
		c.MaxBackups = maxBackups
	}
}

// WithMaxAge sets the maximum number of days to retain old log files
func WithMaxAge(maxAge int) Option {
	return func(c *Config) {
		c.MaxAge = maxAge
	}
}

// WithCompress enables or disables compression of rotated log files
func WithCompress(enabled bool) Option {
	return func(c *Config) {
		c.Compress = enabled
	}
}
