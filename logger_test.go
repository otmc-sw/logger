/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
 **/
package logger

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	if cfg.Level != InfoLevel {
		t.Errorf("DefaultConfig().Level = %v, want %v", cfg.Level, InfoLevel)
	}
	if !cfg.Console {
		t.Error("DefaultConfig().Console should be true")
	}
	if cfg.File {
		t.Error("DefaultConfig().File should be false")
	}
	if cfg.JSON {
		t.Error("DefaultConfig().JSON should be false")
	}
	if !cfg.Caller {
		t.Error("DefaultConfig().Caller should be true")
	}
	if cfg.Stacktrace {
		t.Error("DefaultConfig().Stacktrace should be false")
	}
	if cfg.MaxSize != 10.0 {
		t.Errorf("DefaultConfig().MaxSize = %v, want 10", cfg.MaxSize)
	}
	if cfg.MaxBackups != 3 {
		t.Errorf("DefaultConfig().MaxBackups = %d, want 3", cfg.MaxBackups)
	}
	if cfg.MaxAge != 90 {
		t.Errorf("DefaultConfig().MaxAge = %d, want 90", cfg.MaxAge)
	}
	if cfg.Compress {
		t.Error("DefaultConfig().Compress should be false")
	}
	if cfg.TimeFormat != "2006-01-02 15:04:05.000 -07:00" {
		t.Errorf("DefaultConfig().TimeFormat = %q, want %q", cfg.TimeFormat, "2006-01-02 15:04:05.000 -07:00")
	}
}

func TestParseLevel(t *testing.T) {
	tests := []struct {
		input    string
		expected Level
	}{
		{"TRACE", TraceLevel},
		{"DEBUG", DebugLevel},
		{"INFO", InfoLevel},
		{"WARN", WarnLevel},
		{"ERROR", ErrorLevel},
		{"CRIT", CritLevel},
		{"trace", InfoLevel}, // case sensitive, should default to Info
		{"invalid", InfoLevel},
		{"", InfoLevel},
	}

	for _, tt := range tests {
		result := ParseLevel(tt.input)
		if result != tt.expected {
			t.Errorf("ParseLevel(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestNew(t *testing.T) {
	log := New()
	if log == nil {
		t.Fatal("New() returned nil")
	}

	cfg := log.Config()
	if cfg.Level != InfoLevel {
		t.Errorf("New().Config().Level = %v, want %v", cfg.Level, InfoLevel)
	}
}

func TestNewWithOptions(t *testing.T) {
	log := New(
		WithLevel(DebugLevel),
		WithConsole(false),
		WithCaller(false),
	)

	if log == nil {
		t.Fatal("New() returned nil")
	}

	cfg := log.Config()
	if cfg.Level != DebugLevel {
		t.Errorf("New().Config().Level = %v, want %v", cfg.Level, DebugLevel)
	}
	if cfg.Console {
		t.Error("New().Config().Console should be false")
	}
	if cfg.Caller {
		t.Error("New().Config().Caller should be false")
	}
}

func TestLoggerConfigure(t *testing.T) {
	log := New()
	log.Configure(
		WithLevel(WarnLevel),
		WithCaller(false),
	)

	cfg := log.Config()
	if cfg.Level != WarnLevel {
		t.Errorf("Configure().Level = %v, want %v", cfg.Level, WarnLevel)
	}
	if cfg.Caller {
		t.Error("Configure().Caller should be false")
	}
}

func TestLoggerConfig(t *testing.T) {
	log := New(WithLevel(ErrorLevel))
	cfg := log.Config()

	if cfg.Level != ErrorLevel {
		t.Errorf("Config().Level = %v, want %v", cfg.Level, ErrorLevel)
	}
}

func TestLoggerUpdate(t *testing.T) {
	log := New()
	newCfg := Config{
		Level:      DebugLevel,
		Console:    false,
		File:       false,
		JSON:       false,
		Caller:     false,
		Stacktrace: false,
		MaxSize:    5.0,
		MaxBackups: 2,
		MaxAge:     30,
		Compress:   false,
		TimeFormat: "2006-01-02",
	}

	log.Update(newCfg)
	cfg := log.Config()

	if cfg.Level != DebugLevel {
		t.Errorf("Update().Level = %v, want %v", cfg.Level, DebugLevel)
	}
	if cfg.Console {
		t.Error("Update().Console should be false")
	}
	if cfg.MaxSize != 5.0 {
		t.Errorf("Update().MaxSize = %v, want 5", cfg.MaxSize)
	}
}

func TestLoggerSetLevel(t *testing.T) {
	log := New()
	log.SetLevel(ErrorLevel)

	cfg := log.Config()
	if cfg.Level != ErrorLevel {
		t.Errorf("SetLevel().Level = %v, want %v", cfg.Level, ErrorLevel)
	}
}

func TestLoggerSync(t *testing.T) {
	log := New()
	err := log.Sync()
	if err != nil {
		t.Errorf("Sync() returned error: %v", err)
	}
}

func TestLoggerLogMethods(t *testing.T) {
	tmpDir := t.TempDir()
	logPath := filepath.Join(tmpDir, "test.log")

	log := New(
		WithFile(logPath),
		WithConsole(false),
	)

	log.Trace("trace message")
	log.Debug("debug message")
	log.Info("info message")
	log.Warn("warn message")
	log.Error("error message")
	log.Crit("crit message")

	_ = log.Sync()
}

func TestLoggerRequest(t *testing.T) {
	tmpDir := t.TempDir()
	logPath := filepath.Join(tmpDir, "test.log")

	log := New(
		WithFile(logPath),
		WithConsole(false),
	)

	log.Request("GET", "/api/test", 200, 100*time.Millisecond, "127.0.0.1")
	log.Request("POST", "/api/users", 201, 150*time.Millisecond, "192.168.1.1")
	log.Request("DELETE", "/api/users/1", 204, 50*time.Millisecond, "10.0.0.1")

	_ = log.Sync()
}

func TestLoggerWithFile(t *testing.T) {
	tmpDir := t.TempDir()
	logPath := filepath.Join(tmpDir, "test.log")

	log := New(
		WithFile(logPath),
		WithConsole(false),
	)

	log.Info("test message")
	err := log.Sync()
	if err != nil {
		t.Errorf("Sync() returned error: %v", err)
	}

	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		t.Error("Log file was not created")
	}
}

func TestLoggerWithJSON(t *testing.T) {
	tmpDir := t.TempDir()
	logPath := filepath.Join(tmpDir, "test.json")

	log := New(
		WithFile(logPath),
		WithJSON(true),
		WithConsole(false),
	)

	log.Info("test message")
	err := log.Sync()
	if err != nil {
		t.Errorf("Sync() returned error: %v", err)
	}

	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		t.Error("Log file was not created")
	}
}

func TestGlobalFunctions(t *testing.T) {
	originalCfg := GetConfig()
	defer Update(originalCfg)

	Configure(WithConsole(false))

	Trace("trace message")
	Debug("debug message")
	Info("info message")
	Warn("warn message")
	Error("error message")


	Request("GET", "/api/test", 200, 100*time.Millisecond, "127.0.0.1")

	err := Sync()
	if err != nil {
		t.Errorf("Sync() returned error: %v", err)
	}
}

func TestGlobalSetLevel(t *testing.T) {
	originalCfg := GetConfig()
	defer Update(originalCfg)

	SetLevel(ErrorLevel)
	cfg := GetConfig()

	if cfg.Level != ErrorLevel {
		t.Errorf("SetLevel().Level = %v, want %v", cfg.Level, ErrorLevel)
	}
}

func TestGlobalConfigure(t *testing.T) {
	originalCfg := GetConfig()
	defer Update(originalCfg)

	Configure(
		WithLevel(DebugLevel),
		WithCaller(false),
	)

	cfg := GetConfig()
	if cfg.Level != DebugLevel {
		t.Errorf("Configure().Level = %v, want %v", cfg.Level, DebugLevel)
	}
	if cfg.Caller {
		t.Error("Configure().Caller should be false")
	}
}

func TestGlobalGetConfig(t *testing.T) {
	cfg := GetConfig()

	if cfg.Level != InfoLevel {
		t.Errorf("GetConfig().Level = %v, want %v", cfg.Level, InfoLevel)
	}
}

func TestGlobalUpdate(t *testing.T) {
	originalCfg := GetConfig()
	defer Update(originalCfg)

	newCfg := Config{
		Level:      WarnLevel,
		Console:    false,
		File:       false,
		JSON:       false,
		Caller:     false,
		Stacktrace: false,
		MaxSize:    5.0,
		MaxBackups: 2,
		MaxAge:     30,
		Compress:   false,
		TimeFormat: "2006-01-02",
	}

	Update(newCfg)
	cfg := GetConfig()

	if cfg.Level != WarnLevel {
		t.Errorf("Update().Level = %v, want %v", cfg.Level, WarnLevel)
	}
	if cfg.MaxSize != 5.0 {
		t.Errorf("Update().MaxSize = %v, want 5", cfg.MaxSize)
	}
}

func TestWithLevel(t *testing.T) {
	cfg := DefaultConfig()
	opt := WithLevel(DebugLevel)
	opt(&cfg)

	if cfg.Level != DebugLevel {
		t.Errorf("WithLevel() = %v, want %v", cfg.Level, DebugLevel)
	}
}

func TestWithConsole(t *testing.T) {
	cfg := DefaultConfig()
	opt := WithConsole(false)
	opt(&cfg)

	if cfg.Console {
		t.Error("WithConsole(false) should set Console to false")
	}
}

func TestWithFile(t *testing.T) {
	cfg := DefaultConfig()
	opt := WithFile("test.log")
	opt(&cfg)

	if !cfg.File {
		t.Error("WithFile() should set File to true")
	}
	if cfg.Filename != "test.log" {
		t.Errorf("WithFile() = %q, want %q", cfg.Filename, "test.log")
	}
}

func TestWithJSON(t *testing.T) {
	cfg := DefaultConfig()
	opt := WithJSON(true)
	opt(&cfg)

	if !cfg.JSON {
		t.Error("WithJSON(true) should set JSON to true")
	}
}

func TestWithCaller(t *testing.T) {
	cfg := DefaultConfig()
	opt := WithCaller(false)
	opt(&cfg)

	if cfg.Caller {
		t.Error("WithCaller(false) should set Caller to false")
	}
}

func TestWithStacktrace(t *testing.T) {
	cfg := DefaultConfig()
	opt := WithStacktrace(true)
	opt(&cfg)

	if !cfg.Stacktrace {
		t.Error("WithStacktrace(true) should set Stacktrace to true")
	}
}

func TestWithMaxSize(t *testing.T) {
	cfg := DefaultConfig()
	opt := WithMaxSize(100.0)
	opt(&cfg)

	if cfg.MaxSize != 100.0 {
		t.Errorf("WithMaxSize() = %v, want 100", cfg.MaxSize)
	}
}

func TestWithMaxBackups(t *testing.T) {
	cfg := DefaultConfig()
	opt := WithMaxBackups(10)
	opt(&cfg)

	if cfg.MaxBackups != 10 {
		t.Errorf("WithMaxBackups() = %d, want 10", cfg.MaxBackups)
	}
}

func TestWithMaxAge(t *testing.T) {
	cfg := DefaultConfig()
	opt := WithMaxAge(365)
	opt(&cfg)

	if cfg.MaxAge != 365 {
		t.Errorf("WithMaxAge() = %d, want 365", cfg.MaxAge)
	}
}

func TestWithCompress(t *testing.T) {
	cfg := DefaultConfig()
	opt := WithCompress(false)
	opt(&cfg)

	if cfg.Compress {
		t.Error("WithCompress(false) should set Compress to false")
	}
}

func TestWithTimeFormat(t *testing.T) {
	cfg := DefaultConfig()
	opt := WithTimeFormat("2006-01-02")
	opt(&cfg)

	if cfg.TimeFormat != "2006-01-02" {
		t.Errorf("WithTimeFormat() = %q, want %q", cfg.TimeFormat, "2006-01-02")
	}
}

func TestMultipleOptions(t *testing.T) {
	cfg := DefaultConfig()

	WithLevel(DebugLevel)(&cfg)
	WithConsole(false)(&cfg)
	WithCaller(false)(&cfg)
	WithJSON(true)(&cfg)
	WithMaxSize(50.0)(&cfg)

	if cfg.Level != DebugLevel {
		t.Errorf("Multiple options: Level = %v, want %v", cfg.Level, DebugLevel)
	}
	if cfg.Console {
		t.Error("Multiple options: Console should be false")
	}
	if cfg.Caller {
		t.Error("Multiple options: Caller should be false")
	}
	if !cfg.JSON {
		t.Error("Multiple options: JSON should be true")
	}
	if cfg.MaxSize != 50.0 {
		t.Errorf("Multiple options: MaxSize = %v, want 50", cfg.MaxSize)
	}
}
