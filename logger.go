/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
 **/
package logger

import (
	"sync"
	"time"

	"github.com/otmc-sw/logger/core"
	"github.com/otmc-sw/logger/formatter"
)

var global = New()

type LogEntry = core.Entry

type Logger struct {
	mu        sync.RWMutex
	config    Config
	core      *core.Core
	logs      []LogEntry
	listeners []chan LogEntry
}

func New(opts ...Option) *Logger {
	cfg := DefaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}
	l := &Logger{
		config:    cfg,
		core:      buildCore(cfg),
		logs:      make([]LogEntry, 0),
		listeners: make([]chan LogEntry, 0),
	}
	l.core.AddHook(&streamHook{logger: l})
	return l
}

func (l *Logger) Configure(opts ...Option) {
	l.reconfigure(func(cfg *Config) {
		for _, opt := range opts {
			opt(cfg)
		}
	}, true)
}

func (l *Logger) Config() Config {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.config
}

func (l *Logger) Update(cfg Config) {
	l.reconfigure(func(c *Config) {
		*c = cfg
	}, true)
}

func (l *Logger) SetLevel(level Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.config.Level = level
	l.core.SetLevel(level)
}

func (l *Logger) Trace(format string, args ...any) {
	l.core.Log(core.TraceLevel, 3, format, args...)
}

func (l *Logger) Debug(format string, args ...any) {
	l.core.Log(core.DebugLevel, 3, format, args...)
}

func (l *Logger) Info(format string, args ...any) {
	l.core.Log(core.InfoLevel, 3, format, args...)
}

func (l *Logger) Warn(format string, args ...any) {
	l.core.Log(core.WarnLevel, 3, format, args...)
}

func (l *Logger) Error(format string, args ...any) {
	l.core.Log(core.ErrorLevel, 3, format, args...)
}

func (l *Logger) Crit(format string, args ...any) {
	l.core.Log(core.CritLevel, 3, format, args...)
}

func (l *Logger) Request(method, path string, statusCode int, latency time.Duration, clientIP string) {
	l.core.LogRequest(core.Request{
		Time:       time.Now(),
		Method:     method,
		Path:       path,
		StatusCode: statusCode,
		Latency:    latency,
		ClientIP:   clientIP,
	})
}

func (l *Logger) Sync() error {
	return l.core.Sync()
}

func (l *Logger) Close() error {
	return l.core.Close()
}

func (l *Logger) GetRecentLogs(limit int) []LogEntry {
	l.mu.Lock()
	defer l.mu.Unlock()
	start := 0
	if limit > 0 && len(l.logs) > limit {
		start = len(l.logs) - limit
	}
	res := make([]LogEntry, len(l.logs[start:]))
	copy(res, l.logs[start:])
	return res
}

func (l *Logger) AddListener() chan LogEntry {
	l.mu.Lock()
	defer l.mu.Unlock()
	ch := make(chan LogEntry, 100)
	l.listeners = append(l.listeners, ch)
	return ch
}

func (l *Logger) RemoveListener(ch chan LogEntry) {
	l.mu.Lock()
	defer l.mu.Unlock()
	for i, c := range l.listeners {
		if c == ch {
			l.listeners = append(l.listeners[:i], l.listeners[i+1:]...)
			close(ch)
			break
		}
	}
}

func (l *Logger) reconfigure(fn func(*Config), rebuild bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	fn(&l.config)
	if rebuild {
		l.rebuild()
	}
}

func (l *Logger) rebuild() {
	old := l.core
	l.core = buildCore(l.config)
	l.core.AddHook(&streamHook{logger: l})
	_ = old.Sync()
}

func buildCore(cfg Config) *core.Core {
	return core.NewCore(
		cfg.Level,
		cfg.Caller,
		buildFormatter(cfg),
		buildWriter(cfg),
	)
}

func buildFormatter(cfg Config) core.Formatter {
	if cfg.JSON {
		return formatter.NewJSONFormatter(cfg.TimeFormat)
	}
	return formatter.NewPrettyFormatter(cfg.Console, cfg.TimeFormat)
}

func buildWriter(cfg Config) core.Writer {
	var writers []core.Writer
	if cfg.Console {
		writers = append(writers, core.NewConsoleWriter(nil))
	}
	if cfg.File && cfg.Filename != "" {
		writers = append(writers, core.NewRotateWriter(
			cfg.Filename, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge, cfg.Compress,
		))
	}
	if len(writers) == 0 {
		return nil
	}
	if len(writers) == 1 {
		return writers[0]
	}
	return core.NewMultiWriter(writers...)
}

type streamHook struct {
	logger *Logger
}

func (h *streamHook) Fire(entry core.Entry) error {
	h.logger.mu.Lock()
	defer h.logger.mu.Unlock()

	h.logger.logs = append(h.logger.logs, entry)

	maxEntries := h.logger.config.MaxLogEntries
	if maxEntries > 0 && len(h.logger.logs) > maxEntries {
		h.logger.logs = h.logger.logs[len(h.logger.logs)-maxEntries:]
	}

	for _, ch := range h.logger.listeners {
		select {
		case ch <- entry:
		default:
		}
	}
	return nil
}
