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
		c.alarmPath = alarmPath
		file, err := os.OpenFile(alarmPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err == nil {
			c.alarmFile = file
		}
	}
	if eventPath != "" {
		c.eventPath = eventPath
		file, err := os.OpenFile(eventPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err == nil {
			c.eventFile = file
		}
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

	if c.alarmFile != nil && level >= WarnLevel {
		c.alarmMu.Lock()
		line := fmt.Sprintf("%s %s %s\n", entry.Time.Format("2006-01-02 15:04:05.000 -07:00"), entry.Level.String(), entry.Message)
		_, _ = c.alarmFile.WriteString(line)
		c.alarmMu.Unlock()
	}

	if c.eventFile != nil && level == InfoLevel {
		c.eventMu.Lock()
		line := fmt.Sprintf("%s %s\n", entry.Time.Format("2006-01-02 15:04:05.000 -07:00"), entry.Message)
		_, _ = c.eventFile.WriteString(line)
		c.eventMu.Unlock()
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

func (c *Core) Close() error {
	var err error
	if c.writer != nil {
		if e := c.writer.Close(); e != nil {
			err = e
		}
	}
	if c.alarmFile != nil {
		c.alarmMu.Lock()
		if e := c.alarmFile.Close(); e != nil {
			err = e
		}
		c.alarmFile = nil
		c.alarmMu.Unlock()
	}
	if c.eventFile != nil {
		c.eventMu.Lock()
		if e := c.eventFile.Close(); e != nil {
			err = e
		}
		c.eventFile = nil
		c.eventMu.Unlock()
	}
	return err
}

var osExit = func(code int) {
	os.Exit(code)
}
