/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
 **/
package rotator

type Option func(*Config)

func WithFilename(filename string) Option {
	return func(c *Config) {
		c.Filename = filename
	}
}

func WithMaxSize(maxSize float64) Option {
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

func WithCompress(compress bool) Option {
	return func(c *Config) {
		c.Compress = compress
	}
}

func WithTimeFormat(format string) Option {
	return func(c *Config) {
		c.TimeFormat = format
	}
}

func WithNaming(naming NamingFunc) Option {
	return func(c *Config) {
		c.Naming = naming
	}
}
