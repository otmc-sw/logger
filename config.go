/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package logger

type Config struct {
	Level      Level
	Console    bool
	File       bool
	Filename   string
	JSON       bool
	Caller     bool
	Stacktrace bool
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
	TimeFormat string
}

func DefaultConfig() Config {
	return Config{
		Level:      InfoLevel,
		Console:    true,
		File:       false,
		JSON:       false,
		Caller:     true,
		Stacktrace: false,
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     90,
		Compress:   true,
		TimeFormat: "2006-01-02 15:04:05.000 -07:00",
	}
}
