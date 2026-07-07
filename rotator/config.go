/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
 **/
package rotator

type Config struct {
	Filename   string
	MaxSize    int // megabytes
	MaxBackups int // number of backups
	MaxAge     int // days
	Compress   bool

	TimeFormat string
	Naming     NamingFunc
}

func defaultConfig() Config {
	return Config{
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     30,
		Compress:   false,
		TimeFormat: "20060102",
		Naming:     NameWithIndex,
	}
}
