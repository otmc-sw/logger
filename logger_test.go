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
	defer log.Close()

	log.Trace("trace message")
	log.Debug("debug message")
	log.Info("info message")
	log.Warn("warn message")
	log.Error("error message")
	_ = log.Sync()
}

func TestLoggerRequest(t *testing.T) {
	tmpDir := t.TempDir()
	logPath := filepath.Join(tmpDir, "test.log")

	log := New(
		WithFile(logPath),
		WithConsole(false),
	)
	defer log.Close()

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
	defer log.Close()

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
	defer log.Close()

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

func TestGlobalGetRecentLogs(t *testing.T) {
	originalCfg := GetConfig()
	defer Update(originalCfg)

	Configure(WithConsole(false))

	initialLogs := GetRecentLogs(0)
	initialCount := len(initialLogs)

	Info("global log 1")
	Info("global log 2")
	Info("global log 3")
	Info("global log 4")
	Info("global log 5")

	recentLogs := GetRecentLogs(0)
	expectedCount := initialCount + 5
	if len(recentLogs) != expectedCount {
		t.Errorf("GetRecentLogs(0) returned %d logs, want %d", len(recentLogs), expectedCount)
	}

	recentLogs = GetRecentLogs(3)
	if len(recentLogs) != 3 {
		t.Errorf("GetRecentLogs(3) returned %d logs, want 3", len(recentLogs))
	}

	recentLogs = GetRecentLogs(10)
	if len(recentLogs) != expectedCount {
		t.Errorf("GetRecentLogs(10) returned %d logs, want %d", len(recentLogs), expectedCount)
	}

	recentLogs = GetRecentLogs(-1)
	if len(recentLogs) != expectedCount {
		t.Errorf("GetRecentLogs(-1) returned %d logs, want %d", len(recentLogs), expectedCount)
	}
}

func TestGlobalAddListener(t *testing.T) {
	originalCfg := GetConfig()
	defer Update(originalCfg)

	Configure(WithConsole(false))

	listener := AddListener()
	if listener == nil {
		t.Fatal("AddListener() returned nil")
	}

	Info("global listener test")

	time.Sleep(10 * time.Millisecond)

	select {
	case entry := <-listener:
		t.Logf("Received log entry from global listener: Level=%d", entry.Level)
		if entry.Level != InfoLevel {
			t.Errorf("Expected log level %d, got %d", InfoLevel, entry.Level)
		}
	default:
		t.Error("No message received from global listener channel")
	}

	listener2 := AddListener()
	listener3 := AddListener()

	if listener == listener2 || listener == listener3 || listener2 == listener3 {
		t.Error("AddListener() should return different channels for different calls")
	}

	Info("broadcast to global listeners")

	time.Sleep(20 * time.Millisecond)

	receivedCount := 0
	for _, ch := range []chan LogEntry{listener, listener2, listener3} {
		select {
		case <-ch:
			receivedCount++
		default:
		}
	}

	t.Logf("Global message received by %d out of 3 listeners", receivedCount)

	RemoveListener(listener)
	RemoveListener(listener2)
	RemoveListener(listener3)
}

func TestGlobalRemoveListener(t *testing.T) {
	originalCfg := GetConfig()
	defer Update(originalCfg)

	Configure(WithConsole(false))

	listener := AddListener()
	if listener == nil {
		t.Fatal("AddListener() returned nil")
	}

	Info("before global removal")
	time.Sleep(10 * time.Millisecond)

	select {
	case <-listener:
		t.Log("Global listener received message before removal")
	default:
		t.Error("Global listener should have received message before removal")
	}

	RemoveListener(listener)

	_, ok := <-listener
	if ok {
		t.Error("Global channel should be closed after RemoveListener")
	}

	nonExistentListener := make(chan LogEntry, 100)
	RemoveListener(nonExistentListener)
	t.Log("Global RemoveListener handled non-existent listener without panic")

	listener2 := AddListener()
	RemoveListener(listener2)
	RemoveListener(listener2)
	t.Log("Global RemoveListener handled duplicate removal without panic")
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

func TestGetRecentLogs(t *testing.T) {
	log := New(WithConsole(false))
	defer log.Close()

	log.Info("log 1")
	log.Info("log 2")
	log.Info("log 3")
	log.Info("log 4")
	log.Info("log 5")

	recentLogs := log.GetRecentLogs(0)
	if len(recentLogs) != 5 {
		t.Errorf("GetRecentLogs(0) returned %d logs, want 5", len(recentLogs))
	}

	recentLogs = log.GetRecentLogs(-1)
	if len(recentLogs) != 5 {
		t.Errorf("GetRecentLogs(-1) returned %d logs, want 5", len(recentLogs))
	}

	recentLogs = log.GetRecentLogs(10)
	if len(recentLogs) != 5 {
		t.Errorf("GetRecentLogs(10) returned %d logs, want 5", len(recentLogs))
	}

	recentLogs = log.GetRecentLogs(3)
	if len(recentLogs) != 3 {
		t.Errorf("GetRecentLogs(3) returned %d logs, want 3", len(recentLogs))
	}

	recentLogs = log.GetRecentLogs(2)
	if len(recentLogs) != 2 {
		t.Errorf("GetRecentLogs(2) returned %d logs, want 2", len(recentLogs))
	}

	if len(recentLogs) == 2 {
		t.Logf("GetRecentLogs(2) returned %d most recent logs", len(recentLogs))
	}

	emptyLog := New(WithConsole(false))
	defer emptyLog.Close()
	emptyLogs := emptyLog.GetRecentLogs(5)
	if len(emptyLogs) != 0 {
		t.Errorf("GetRecentLogs(5) on empty logger returned %d logs, want 0", len(emptyLogs))
	}
}

func TestAddListener(t *testing.T) {
	log := New(WithConsole(false))
	defer log.Close()

	listener := log.AddListener()
	if listener == nil {
		t.Fatal("AddListener() returned nil")
	}


	listener2 := log.AddListener()
	if listener2 == nil {
		t.Fatal("AddListener() returned nil for second listener")
	}

	if listener == listener2 {
		t.Error("AddListener() should return different channels for different calls")
	}

	log.Info("test message for listener")

	time.Sleep(10 * time.Millisecond)

	select {
	case entry := <-listener:
		t.Logf("Received log entry from listener: Level=%d", entry.Level)
		if entry.Level != InfoLevel {
			t.Errorf("Expected log level %d, got %d", InfoLevel, entry.Level)
		}
	default:
		t.Error("No message received from listener channel")
	}

	listeners := make([]chan LogEntry, 0)
	for i := 0; i < 5; i++ {
		ch := log.AddListener()
		if ch == nil {
			t.Fatalf("AddListener() returned nil at iteration %d", i)
		}
		listeners = append(listeners, ch)
	}

	log.Info("broadcast test")

	time.Sleep(20 * time.Millisecond)

	receivedCount := 0
	for _, ch := range listeners {
		select {
		case <-ch:
			receivedCount++
		default:
		}
	}

	t.Logf("Message received by %d out of %d listeners", receivedCount, len(listeners))

	for _, ch := range listeners {
		log.RemoveListener(ch)
	}
}

func TestRemoveListener(t *testing.T) {
	log := New(WithConsole(false))
	defer log.Close()

	listener := log.AddListener()
	if listener == nil {
		t.Fatal("AddListener() returned nil")
	}

	log.Info("before removal")
	time.Sleep(10 * time.Millisecond)

	select {
	case <-listener:
		t.Log("Listener received message before removal")
	default:
		t.Error("Listener should have received message before removal")
	}

	log.RemoveListener(listener)

	_, ok := <-listener
	if ok {
		t.Error("Channel should be closed after RemoveListener")
	}

	log.Info("after removal")
	time.Sleep(10 * time.Millisecond)

	_, ok = <-listener
	if ok {
		t.Error("Channel should still be closed after sending new logs")
	}

	nonExistentListener := make(chan LogEntry, 100)
	log.RemoveListener(nonExistentListener)
	t.Log("RemoveListener handled non-existent listener without panic")

	listener2 := log.AddListener()
	log.RemoveListener(listener2)
	log.RemoveListener(listener2)
	t.Log("RemoveListener handled duplicate removal without panic")

	listener3 := log.AddListener()
	listener4 := log.AddListener()
	listener5 := log.AddListener()

	log.RemoveListener(listener3)
	log.RemoveListener(listener5)

	_, ok = <-listener3
	if ok {
		t.Error("listener3 should be closed")
	}

	_, ok = <-listener5
	if ok {
		t.Error("listener5 should be closed")
	}

	log.Info("listener4 test")
	time.Sleep(10 * time.Millisecond)

	select {
	case entry := <-listener4:
		if entry.Level == InfoLevel {
			t.Logf("listener4 is still active and received message (Level=%d)", entry.Level)
		}
	default:
		t.Error("listener4 should still be active and receive messages")
	}

	log.RemoveListener(listener4)
}
