/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package internal

import (
	"fmt"
	"io"
	"os"
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

func (l Level) String() string {
	switch l {
	case TraceLevel:
		return "TRACE"
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case CritLevel:
		return "CRIT"
	default:
		return "UNKNOWN"
	}
}

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

type Writer interface {
	io.Writer
	Sync() error
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

func NewCore(level Level, caller bool, formatter Formatter, writer Writer) *Core {
	return &Core{
		level:     level,
		enabled:   true,
		caller:    caller,
		formatter: formatter,
		writer:    writer,
		hooks:     make([]Hook, 0),
	}
}

func (c *Core) SetLevel(level Level) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.level = level
}

func (c *Core) Enable(enabled bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.enabled = enabled
}

func (c *Core) AddHook(hook Hook) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.hooks = append(c.hooks, hook)
}

func (c *Core) Log(level Level, skip int, format string, args ...any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.enabled || level < c.level {
		return
	}

	message := fmt.Sprintf(format, args...)

	var entry Entry
	entry.Time = time.Now()
	entry.Level = level
	entry.Message = message

	if c.caller {
		caller := GetCaller(skip)
		entry.Function = caller.Function
		entry.File = caller.File
		entry.Line = caller.Line
	}

	formatted := c.formatter.Format(entry)

	if c.writer != nil {
		_, _ = c.writer.Write([]byte(formatted))
	}

	for _, hook := range c.hooks {
		_ = hook.Fire(entry)
	}

	if level == CritLevel {
		osExit(1)
	}
}

func (c *Core) LogRequest(req Request) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.enabled {
		return
	}

	formatted := c.formatter.FormatRequest(req)

	if c.writer != nil {
		_, _ = c.writer.Write([]byte(formatted))
	}
}

func (c *Core) Sync() error {
	if c.writer != nil {
		return c.writer.Sync()
	}
	return nil
}

var osExit = func(code int) {
	os.Exit(code)
}
