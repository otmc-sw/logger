/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package logger

type Option func(*Config)

func WithLevel(level Level) Option {
	return func(c *Config) {
		c.Level = level
	}
}

func WithConsole(enabled bool) Option {
	return func(c *Config) {
		c.Console = enabled
	}
}

func WithFile(filename string) Option {
	return func(c *Config) {
		c.File = true
		c.Filename = filename
	}
}

func WithJSON(enabled bool) Option {
	return func(c *Config) {
		c.JSON = enabled
	}
}

func WithCaller(enabled bool) Option {
	return func(c *Config) {
		c.Caller = enabled
	}
}

func WithStacktrace(enabled bool) Option {
	return func(c *Config) {
		c.Stacktrace = enabled
	}
}

func WithMaxSize(maxSize int) Option {
	return func(c *Config) {
		c.MaxSize = maxSize
	}
}

func WithMaxBackups(maxBackups int) Option {
	return func(c *Config) {
		c.MaxBackups = maxBackups
	}
}

func WithMaxAge(maxAge int) Option {
	return func(c *Config) {
		c.MaxAge = maxAge
	}
}

func WithCompress(enabled bool) Option {
	return func(c *Config) {
		c.Compress = enabled
	}
}
