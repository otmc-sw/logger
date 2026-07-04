/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package loggers

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Level int

const (
	OFF Level = iota
	CRIT
	ERROR
	WARN
	INFO
	DEBUG
	TRACE
)

const defaultMaxSize = 5 * 1024 * 1024 // 5MB

const (
	colorReset = "\033[0m"

	colorCrit  = "\033[101m" // Light red background
	colorError = "\033[41m"  // Red background
	colorWarn  = "\033[43m"  // Yellow background
	colorInfo  = "\033[44m"  // Blue background
	colorDebug = "\033[46m"  // Cyan background
	colorTrace = "\033[47m"  // White background

	colorErrorText = "\033[31m" // Red text
	colorWarnText  = "\033[33m" // Yellow text
)

type Logger struct {
	mu           sync.Mutex
	level        Level
	enabled      bool
	consoleOut   io.Writer
	file         *os.File
	filePath     string
	maxSize      int64
	alarmFile    *os.File
	alarmPath    string
	alarmMaxSize int64
}

func New(level Level, writers ...io.Writer) *Logger {
	if len(writers) == 0 {
		writers = []io.Writer{os.Stdout}
	}

	return &Logger{
		level:      level,
		enabled:    true,
		consoleOut: io.MultiWriter(writers...),
		maxSize:    defaultMaxSize,
	}
}

func (l *Logger) WithFile(path string, maxSize int64) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if maxSize <= 0 {
		maxSize = defaultMaxSize
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		return err
	}

	l.file = f
	l.filePath = path
	l.maxSize = maxSize

	return nil
}

func (l *Logger) WithAlarmFile(path string, maxSize int64) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if maxSize <= 0 {
		maxSize = defaultMaxSize
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		return err
	}

	l.alarmFile = f
	l.alarmPath = path
	l.alarmMaxSize = maxSize

	return nil
}

func (l *Logger) SetLevel(level Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

func (l *Logger) Enable(v bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.enabled = v
}

func (l *Logger) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	var err error
	if l.file != nil {
		err = l.file.Close()
	}
	if l.alarmFile != nil {
		if e := l.alarmFile.Close(); e != nil {
			return e
		}
	}
	return err
}

func (l *Logger) log(level Level, format string, args ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if !l.enabled || level > l.level || level == OFF {
		return
	}

	if l.file != nil {
		_ = l.rotateIfNeeded()
	}
	if l.alarmFile != nil && level <= WARN {
		_ = l.rotateAlarmFileIfNeeded()
	}

	fn, file, line := caller(3)
	timestamp := time.Now().Format("2006-01-02 15:04:05.000 -07:00")

	msg := fmt.Sprintf(format, args...)

	consoleMsg := msg
	switch level {
	case CRIT, ERROR:
		consoleMsg = colorErrorText + msg + colorReset
	case WARN:
		consoleMsg = colorWarnText + msg + colorReset
	}

	entry := fmt.Sprintf(
		"%s %15.15s() %12.12s:%-5d |%s| %s\n",
		timestamp,
		fn,
		file,
		line,
		level.String(),
		consoleMsg,
	)

	_, _ = l.consoleOut.Write([]byte(entry))

	if l.file != nil {
		fileEntry := stripColorCodes(entry)
		_, _ = l.file.Write([]byte(fileEntry))
	}

	if l.alarmFile != nil && level <= WARN {
		alarmEntry := stripColorCodes(entry)
		_, _ = l.alarmFile.Write([]byte(alarmEntry))
	}
}

func stripColorCodes(s string) string {
	s = strings.ReplaceAll(s, colorReset, "")
	s = strings.ReplaceAll(s, colorCrit, "")
	s = strings.ReplaceAll(s, colorError, "")
	s = strings.ReplaceAll(s, colorWarn, "")
	s = strings.ReplaceAll(s, colorInfo, "")
	s = strings.ReplaceAll(s, colorDebug, "")
	s = strings.ReplaceAll(s, colorTrace, "")
	s = strings.ReplaceAll(s, colorErrorText, "")
	s = strings.ReplaceAll(s, colorWarnText, "")
	return s
}

func (l *Logger) rotateIfNeeded() error {
	info, err := l.file.Stat()
	if err != nil {
		return err
	}

	if info.Size() < l.maxSize {
		return nil
	}

	_ = l.file.Close()

	backup := l.filePath + ".1"
	_ = os.Remove(backup)
	if err := os.Rename(l.filePath, backup); err != nil {
		return err
	}

	f, err := os.OpenFile(l.filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		return err
	}

	l.file = f

	return nil
}

func (l *Logger) rotateAlarmFileIfNeeded() error {
	if l.alarmFile == nil {
		return nil
	}

	info, err := l.alarmFile.Stat()
	if err != nil {
		return err
	}

	if info.Size() < l.alarmMaxSize {
		return nil
	}

	_ = l.alarmFile.Close()

	backup := l.alarmPath + ".1"
	_ = os.Remove(backup)
	if err := os.Rename(l.alarmPath, backup); err != nil {
		return err
	}

	f, err := os.OpenFile(l.alarmPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		return err
	}

	l.alarmFile = f

	return nil
}

func caller(skip int) (string, string, int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "Unknown", "Unknown", 0
	}

	fn := "Unknown"
	if f := runtime.FuncForPC(pc); f != nil {
		parts := strings.Split(f.Name(), ".")
		fn = parts[len(parts)-1]
	}

	return fn, filepath.Base(file), line
}

func (l Level) String() string {
	switch l {
	case TRACE:
		return colorTrace + " TRACE " + colorReset
	case DEBUG:
		return colorDebug + " DEBUG " + colorReset
	case INFO:
		return colorInfo + " INFO  " + colorReset
	case WARN:
		return colorWarn + " WARN  " + colorReset
	case ERROR:
		return colorError + " ERROR " + colorReset
	case CRIT:
		return colorCrit + " CRIT  " + colorReset
	default:
		return "UNKNOWN"
	}
}

func (l *Logger) Trace(f string, a ...any) { l.log(TRACE, f, a...) }
func (l *Logger) Debug(f string, a ...any) { l.log(DEBUG, f, a...) }
func (l *Logger) Info(f string, a ...any)  { l.log(INFO, f, a...) }
func (l *Logger) Warn(f string, a ...any)  { l.log(WARN, f, a...) }
func (l *Logger) Error(f string, a ...any) { l.log(ERROR, f, a...) }
func (l *Logger) Crit(f string, a ...any)  { l.log(CRIT, f, a...) }

func NewLLMLogger(level Level, logDir string) (*Logger, error) {
	if logDir == "" {
		logDir = "data/logs"
	}

	llmLogPath := filepath.Join(logDir, "llm.log")

	logger := New(level)

	if err := logger.WithFile(llmLogPath, defaultMaxSize); err != nil {
		return nil, fmt.Errorf("failed to create LLM log file: %w", err)
	}

	return logger, nil
}
