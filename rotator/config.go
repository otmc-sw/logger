/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
 **/
package rotator

// Config holds the rotator configuration.
type Config struct {
	Filename   string
	MaxSize    int // megabytes
	MaxBackups int // number of backups
	MaxAge     int // days
	Compress   bool

	TimeFormat string
	Naming     NamingFunc
}

// defaultConfig returns the default configuration.
func defaultConfig() Config {
	return Config{
		MaxSize:    100, // 100 MB
		MaxBackups: 3,
		MaxAge:     30, // 30 days
		Compress:   false,
		TimeFormat: "20060102",
		Naming:     NameWithIndex,
	}
}
