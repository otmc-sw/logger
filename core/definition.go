/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package core

import (
	"sync"
	"time"
)

type Level int

const (
	TraceLevel Level = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	CritLevel
)

type Entry struct {
	Time     time.Time
	Level    Level
	Function string
	File     string
	Line     int
	Message  string
}

type Request struct {
	Time       time.Time
	Method     string
	Path       string
	StatusCode int
	Latency    time.Duration
	ClientIP   string
}

type Formatter interface {
	Format(entry Entry) string
	FormatRequest(req Request) string
}

type Hook interface {
	Fire(entry Entry) error
}

type Core struct {
	mu        sync.Mutex
	level     Level
	enabled   bool
	caller    bool
	formatter Formatter
	writer    Writer
	hooks     []Hook
}
