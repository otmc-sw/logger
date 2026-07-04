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

	"github.com/otmc-sw/logger"
)

// Formatter is the interface for formatting log entries
type Formatter interface {
	Format(entry logger.Entry) string
}

// Writer is the interface for writing log output
type Writer interface {
	io.Writer
	Sync() error
}

// Hook is the interface for log hooks
type Hook interface {
	Fire(entry logger.Entry) error
}

// Core is the core logger implementation
type Core struct {
	mu        sync.Mutex
	level     logger.Level
	enabled   bool
	caller    bool
	formatter Formatter
	writer    Writer
	hooks     []Hook
}

// NewCore creates a new core logger
func NewCore(level logger.Level, caller bool, formatter Formatter, writer Writer) *Core {
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
func (c *Core) SetLevel(level logger.Level) {
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
func (c *Core) Log(level logger.Level, format string, args ...any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.enabled || level < c.level {
		return
	}

	message := fmt.Sprintf(format, args...)

	var entry logger.Entry
	entry.Time = time.Now()
	entry.Level = level
	entry.Message = message

	if c.caller {
		caller := GetCaller(3)
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
	if level == logger.FatalLevel {
		osExit(1)
	}
	if level == logger.PanicLevel {
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

// osExit is a wrapper for os.Exit to allow testing
var osExit = func(code int) {
	os.Exit(code)
}
