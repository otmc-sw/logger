/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
 **/
package rotator

// Option configures the rotator.
type Option func(*Config)

// WithFilename sets the log filename.
func WithFilename(filename string) Option {
	return func(c *Config) {
		c.Filename = filename
	}
}

// WithMaxSize sets the maximum size in megabytes before rotation.
func WithMaxSize(maxSize int) Option {
	return func(c *Config) {
		c.MaxSize = maxSize
	}
}

// WithMaxBackups sets the maximum number of backup files to keep.
func WithMaxBackups(maxBackups int) Option {
	return func(c *Config) {
		c.MaxBackups = maxBackups
	}
}

// WithMaxAge sets the maximum age in days for backup files.
func WithMaxAge(maxAge int) Option {
	return func(c *Config) {
		c.MaxAge = maxAge
	}
}

// WithCompress enables or disables gzip compression for rotated files.
func WithCompress(compress bool) Option {
	return func(c *Config) {
		c.Compress = compress
	}
}

// WithTimeFormat sets the time format for naming rotated files.
func WithTimeFormat(format string) Option {
	return func(c *Config) {
		c.TimeFormat = format
	}
}

// WithNaming sets a custom naming function for rotated files.
func WithNaming(naming NamingFunc) Option {
	return func(c *Config) {
		c.Naming = naming
	}
}
