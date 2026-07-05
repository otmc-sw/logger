/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/otmc-sw/logger"
)

func main() {
	// Create logs directory
	_ = os.MkdirAll("logs", 0755)

	// Test 1: Basic console logging with all levels
	testBasicConsoleLogging()

	// Test 2: File logging
	testFileLogging()

	// Test 3: JSON formatting
	testJSONLogging()

	// Test 4: Custom logger instances
	testCustomLogger()

	// Test 5: Log level filtering
	testLogLevelFiltering()

	// Test 6: Request logging
	testRequest()

	logger.Info("✅ All tests completed!")
	_ = logger.Sync()

	testCrit()
}

func testBasicConsoleLogging() {
	logger.Info("=== Test 1: Basic Console Logging ===")

	// Initialize logger for console only
	logger.Init(logger.Config{
		Level:   logger.TraceLevel,
		Console: true,
		Caller:  true,
	})

	logger.Trace("🔍 Trace message - detailed debugging")
	logger.Debug("🐛 Debug message - debugging info")
	logger.Info("✅ Info message - general information")
	logger.Warn("⚠️ Warn message - warning condition")
	logger.Error("❌ Error message - error occurred")
}

func testFileLogging() {
	logger.Info("=== Test 2: File Logging ===")

	logPath, _ := filepath.Abs("logs/test.log")

	// Initialize logger for file output
	logger.Init(logger.Config{
		Level:    logger.DebugLevel,
		Console:  true,
		File:     true,
		Filename: logPath,
		Caller:   true,
	})

	logger.Trace("🔍 This trace should not appear (level is Debug)")
	logger.Debug("🐛 Debug message to file")
	logger.Info("✅ Info message to file")
	logger.Warn("⚠️ Warn message to file")
	logger.Error("❌ Error message to file")

	logger.Info("📁 Check logs/test.log for file output")
	_ = logger.Sync()
}

func testJSONLogging() {
	logger.Info("=== Test 3: JSON Formatting ===")

	logPath, _ := filepath.Abs("logs/test.json")

	// Initialize logger with JSON formatting
	logger.Init(logger.Config{
		Level:    logger.InfoLevel,
		Console:  true,
		File:     true,
		Filename: logPath,
		JSON:     true,
		Caller:   true,
	})

	logger.Info("✅ JSON formatted message")
	logger.Warn("⚠️ JSON warning message")
	logger.Error("❌ JSON error message")

	logger.Info("📁 Check logs/test.json for JSON output")
	_ = logger.Sync()
}

func testCustomLogger() {
	logger.Info("=== Test 4: Custom Logger Instances ===")

	// Create custom console logger
	consoleLog := logger.New(
		logger.WithConsole(true),
		logger.WithLevel(logger.DebugLevel),
		logger.WithCaller(true),
	)

	consoleLog.Trace("🔍 Custom logger - trace")
	consoleLog.Debug("🐛 Custom logger - debug")
	consoleLog.Info("✅ Custom logger - info")
	consoleLog.Warn("⚠️ Custom logger - warn")
	consoleLog.Error("❌ Custom logger - error")

	// Create custom file logger
	logPath, _ := filepath.Abs("logs/custom.log")
	fileLog := logger.New(
		logger.WithFile(logPath),
		logger.WithLevel(logger.InfoLevel),
		logger.WithCaller(true),
	)

	fileLog.Info("✅ Custom file logger - info")
	fileLog.Warn("⚠️ Custom file logger - warn")
	fileLog.Error("❌ Custom file logger - error")

	_ = fileLog.Sync()
	logger.Info("📁 Check logs/custom.log for custom logger output")
}

func testLogLevelFiltering() {
	logger.Info("=== Test 5: Log Level Filtering ===")

	// Test with Warn level - should only show Warn, Error, Crit
	logger.Init(logger.Config{
		Level:   logger.WarnLevel,
		Console: true,
		Caller:  true,
	})

	logger.Trace("🔍 This trace should NOT appear")
	logger.Debug("🐛 This debug should NOT appear")
	logger.Info("✅ This info should NOT appear")
	logger.Warn("⚠️ This warn SHOULD appear")
	logger.Error("❌ This error SHOULD appear")

	// Test with Error level - should only show Error, Crit
	logger.Init(logger.Config{
		Level:   logger.ErrorLevel,
		Console: true,
		Caller:  true,
	})

	logger.Trace("🔍 This trace should NOT appear")
	logger.Debug("🐛 This debug should NOT appear")
	logger.Info("✅ This info should NOT appear")
	logger.Warn("⚠️ This warn should NOT appear")
	logger.Error("❌ This error SHOULD appear")

	// Reset to Info level for remaining tests
	logger.Init(logger.Config{
		Level:   logger.InfoLevel,
		Console: true,
		Caller:  true,
	})
}

func testCrit() {
	logger.Info("=== Test 6: Critical Logging ===")
	logger.Crit("💥 This crit SHOULD appear and program will exit")
}

func testRequest() {
	logger.Request("GET", "/documents", 200, 1*time.Millisecond, "127.0.0.1")
	logger.Request("POST", "/api/users", 201, 150*time.Millisecond, "192.168.1.100")
	logger.Request("DELETE", "/api/users/123", 204, 50*time.Millisecond, "10.0.0.1")
	logger.Request("GET", "/not-found", 404, 10*time.Millisecond, "127.0.0.1")
	logger.Request("POST", "/error", 500, 200*time.Millisecond, "192.168.1.50")
	logger.Request("PUT", "/api/data", 200, 100*time.Millisecond, "172.16.0.100")
}
