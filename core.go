/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package logger

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

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
func (c *Core) Log(level Level, format string, args ...any) {
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
		caller := getCaller(3)
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

	// Handle fatal and panic
	if level == FatalLevel {
		osExit(1)
	}
	if level == PanicLevel {
		panic(message)
	}
}

// Sync flushes any buffered log entries
func (c *Core) Sync() error {
	if c.writer != nil {
		return c.writer.Sync()
	}
	return nil
}

// Writer is the interface for writing log output
type Writer interface {
	io.Writer
	Sync() error
}

// osExit is a wrapper for os.Exit to allow testing
var osExit = func(code int) {
	os.Exit(code)
}
