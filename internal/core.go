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

// Level represents the severity level of a log entry
type Level int

const (
	TraceLevel Level = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	CritLevel
)

// String returns the string representation of the level
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

// Entry represents a log entry
type Entry struct {
	Time     time.Time
	Level    Level
	Function string
	File     string
	Line     int
	Message  string
}

// Formatter is the interface for formatting log entries
type Formatter interface {
	Format(entry Entry) string
}

// Writer is the interface for writing log output
type Writer interface {
	io.Writer
	Sync() error
}

// Hook is the interface for log hooks
type Hook interface {
	Fire(entry Entry) error
}

// Core is the core logger implementation
type Core struct {
	mu        sync.Mutex
	level     Level
	enabled   bool
	caller    bool
	formatter Formatter
	writer    Writer
	hooks     []Hook
}

// NewCore creates a new core logger
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

// SetLevel sets the log level
func (c *Core) SetLevel(level Level) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.level = level
}

// Enable enables or disables the logger
func (c *Core) Enable(enabled bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.enabled = enabled
}

// AddHook adds a hook to the logger
func (c *Core) AddHook(hook Hook) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.hooks = append(c.hooks, hook)
}

// Log logs a message at the specified level
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

	// Format the entry
	formatted := c.formatter.Format(entry)

	// Write to output
	if c.writer != nil {
		_, _ = c.writer.Write([]byte(formatted))
	}

	// Execute hooks
	for _, hook := range c.hooks {
		_ = hook.Fire(entry)
	}

	// Handle crit
	if level == CritLevel {
		osExit(1)
	}
}

// Sync flushes any buffered log entries
func (c *Core) Sync() error {
	if c.writer != nil {
		return c.writer.Sync()
	}
	return nil
}

// osExit is a wrapper for os.Exit to allow testing
var osExit = func(code int) {
	os.Exit(code)
}
