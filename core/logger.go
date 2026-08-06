/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package core

import (
	"fmt"
	"os"
	"time"
)

func NewCore(level Level, caller bool, formatter Formatter, writer Writer, alarmPath string, eventPath string) *Core {
	c := &Core{
		level:     level,
		enabled:   true,
		caller:    caller,
		formatter: formatter,
		writer:    writer,
		hooks:     make([]Hook, 0),
	}
	if alarmPath != "" {
		c.alarmRotator = NewRotateWriter(alarmPath, 10, 3, 30, false)
	}
	if eventPath != "" {
		c.eventRotator = NewRotateWriter(eventPath, 10, 3, 30, false)
	}
	return c
}

func (l Level) String() string {
	switch l {
	case TraceLevel:
		return "TRACE"
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case EventLevel:
		return "EVENT"
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
	c.LogWithMetadata(level, skip, nil, format, args...)
}

func (c *Core) LogWithMetadata(level Level, skip int, metadata interface{}, format string, args ...any) {
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
	entry.Metadata = metadata

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

	if c.alarmRotator != nil && level >= WarnLevel {
		_, _ = c.alarmRotator.Write([]byte(c.formatter.Format(entry)))
	}

	if c.eventRotator != nil && level == EventLevel {
		_, _ = c.eventRotator.Write([]byte(c.formatter.Format(entry)))
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
	var err error
	if c.writer != nil {
		if e := c.writer.Sync(); e != nil {
			err = e
		}
	}
	if c.alarmRotator != nil {
		if e := c.alarmRotator.Sync(); e != nil {
			err = e
		}
	}
	if c.eventRotator != nil {
		if e := c.eventRotator.Sync(); e != nil {
			err = e
		}
	}
	return err
}

func (c *Core) Close() error {
	var err error
	if c.writer != nil {
		if e := c.writer.Close(); e != nil {
			err = e
		}
	}
	if c.alarmRotator != nil {
		if e := c.alarmRotator.Close(); e != nil {
			err = e
		}
	}
	if c.eventRotator != nil {
		if e := c.eventRotator.Close(); e != nil {
			err = e
		}
	}
	return err
}

var osExit = func(code int) {
	os.Exit(code)
}
