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
	_ = os.MkdirAll("logs", 0755)

	testBasicConsoleLogging()

	testFileLogging()

	testJSONLogging()

	testCustomLogger()

	testLogLevelFiltering()

	testRequest()

	logger.Info("âœ… All tests completed!")
	_ = logger.Sync()

	testCrit()
}

func testBasicConsoleLogging() {
	logger.Info("=== Test 1: Basic Console Logging ===")

	logger.Init(logger.Config{
		Level:   logger.TraceLevel,
		Console: true,
		Caller:  true,
	})

	logger.Trace("ðŸ” Trace message - detailed debugging")
	logger.Debug("ðŸ› Debug message - debugging info")
	logger.Info("âœ… Info message - general information")
	logger.Warn("âš ï¸ Warn message - warning condition")
	logger.Error("âŒ Error message - error occurred")
}

func testFileLogging() {
	logger.Info("=== Test 2: File Logging ===")

	logPath, _ := filepath.Abs("logs/test.log")

	logger.Init(logger.Config{
		Level:    logger.DebugLevel,
		Console:  true,
		File:     true,
		Filename: logPath,
		Caller:   true,
	})

	logger.Trace("ðŸ” This trace should not appear (level is Debug)")
	logger.Debug("ðŸ› Debug message to file")
	logger.Info("âœ… Info message to file")
	logger.Warn("âš ï¸ Warn message to file")
	logger.Error("âŒ Error message to file")

	logger.Info("ðŸ“ Check logs/test.log for file output")
	_ = logger.Sync()
}

func testJSONLogging() {
	logger.Info("=== Test 3: JSON Formatting ===")

	logPath, _ := filepath.Abs("logs/test.json")

	logger.Init(logger.Config{
		Level:    logger.InfoLevel,
		Console:  true,
		File:     true,
		Filename: logPath,
		JSON:     true,
		Caller:   true,
	})

	logger.Info("âœ… JSON formatted message")
	logger.Warn("âš ï¸ JSON warning message")
	logger.Error("âŒ JSON error message")

	logger.Info("ðŸ“ Check logs/test.json for JSON output")
	_ = logger.Sync()
}

func testCustomLogger() {
	logger.Info("=== Test 4: Custom Logger Instances ===")

	consoleLog := logger.New(
		logger.WithConsole(true),
		logger.WithLevel(logger.DebugLevel),
		logger.WithCaller(true),
	)

	consoleLog.Trace("ðŸ” Custom logger - trace")
	consoleLog.Debug("ðŸ› Custom logger - debug")
	consoleLog.Info("âœ… Custom logger - info")
	consoleLog.Warn("âš ï¸ Custom logger - warn")
	consoleLog.Error("âŒ Custom logger - error")

	logPath, _ := filepath.Abs("logs/custom.log")
	fileLog := logger.New(
		logger.WithFile(logPath),
		logger.WithLevel(logger.InfoLevel),
		logger.WithCaller(true),
	)

	fileLog.Info("âœ… Custom file logger - info")
	fileLog.Warn("âš ï¸ Custom file logger - warn")
	fileLog.Error("âŒ Custom file logger - error")

	_ = fileLog.Sync()
	logger.Info("ðŸ“ Check logs/custom.log for custom logger output")
}

func testLogLevelFiltering() {
	logger.Info("=== Test 5: Log Level Filtering ===")

	logger.Init(logger.Config{
		Level:   logger.WarnLevel,
		Console: true,
		Caller:  true,
	})

	logger.Trace("ðŸ” This trace should NOT appear")
	logger.Debug("ðŸ› This debug should NOT appear")
	logger.Info("âœ… This info should NOT appear")
	logger.Warn("âš ï¸ This warn SHOULD appear")
	logger.Error("âŒ This error SHOULD appear")

	logger.Init(logger.Config{
		Level:   logger.ErrorLevel,
		Console: true,
		Caller:  true,
	})

	logger.Trace("ðŸ” This trace should NOT appear")
	logger.Debug("ðŸ› This debug should NOT appear")
	logger.Info("âœ… This info should NOT appear")
	logger.Warn("âš ï¸ This warn should NOT appear")
	logger.Error("âŒ This error SHOULD appear")

	logger.Init(logger.Config{
		Level:   logger.InfoLevel,
		Console: true,
		Caller:  true,
	})
}

func testCrit() {
	logger.Info("=== Test 6: Critical Logging ===")
	logger.Crit("ðŸ’¥ This crit SHOULD appear and program will exit")
}

func testRequest() {
	logger.Request("GET", "/documents", 200, 1*time.Millisecond, "127.0.0.1")
	logger.Request("POST", "/api/users", 201, 150*time.Millisecond, "192.168.1.100")
	logger.Request("DELETE", "/api/users/123", 204, 50*time.Millisecond, "10.0.0.1")
	logger.Request("GET", "/not-found", 404, 10*time.Millisecond, "127.0.0.1")
	logger.Request("POST", "/error", 500, 200*time.Millisecond, "192.168.1.50")
	logger.Request("PUT", "/api/data", 200, 100*time.Millisecond, "172.16.0.100")
	logger.Request("PATCH", "/api/data", 200, 100*time.Millisecond, "172.16.0.100")

}
