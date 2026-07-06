/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/otmc-sw/logger"
)

var TEST_ID = 1

func main() {
	_ = os.MkdirAll("logs", 0755)
	testBasicConsoleLogging()
	testNoCaller()
	testFileLogging()
	testJSONLogging()
	testCustomLogger()
	testLogLevelFiltering()
	testRequest()
	testRotate()
	_ = logger.Sync()

	testCrit()
}

func buildHeader(title string) {
	if TEST_ID > 1 {
		fmt.Println("")
	}

	const green = "\033[32m"
	const reset = "\033[0m"

	content := fmt.Sprintf("│ ▶ TEST %d: %s ", TEST_ID, title)

	lineLength := len([]rune(content))

	topBorder := "┌"
	bottomBorder := "└"
	for i := 0; i < lineLength-1; i++ {
		topBorder += "─"
		bottomBorder += "─"
	}
	topBorder += "┐"
	bottomBorder += "┘"

	content += "│"

	fmt.Println(green + topBorder + reset)
	fmt.Println(green + content + reset)
	fmt.Println(green + bottomBorder + reset)

	TEST_ID++
}

func testBasicConsoleLogging() {
	buildHeader("Basic Console Logging")

	logger.Trace("🔍 Trace message - detailed debugging")
	logger.Debug("🐛 Debug message - debugging info")
	logger.Info("✅ Info message - general information")
	logger.Warn("⚠️ Warn message - warning condition")
	logger.Error("❌ Error message - error occurred")
}

func testNoCaller() {
	buildHeader("No Caller Information")

	logger.Configure(logger.WithCaller(false))
	logger.Debug("🐛 Debug message - debugging info")
	logger.Info("✅ Info message - general information")
	logger.Warn("⚠️ Warn message - warning condition")
	logger.Error("❌ Error message - error occurred")

	logger.Configure(logger.WithCaller(true))
}

func testFileLogging() {
	buildHeader("File Logging")

	logger.Configure(logger.WithFile("logs/test.log"))
	logger.Trace("🔍 This trace should not appear (level is Debug)")
	logger.Debug("🐛 Debug message to file")
	logger.Info("✅ Info message to file")
	logger.Warn("⚠️ Warn message to file")
	logger.Error("❌ Error message to file")

	logger.Info("📁 Check logs/test.log for file output")
	_ = logger.Sync()

	logger.Configure(logger.WithConsole(true))
}

func testJSONLogging() {
	buildHeader("JSON Formatting")
	logger.Configure(logger.WithJSON(true), logger.WithTimeFormat("15:04:05.000"))

	logger.Info("✅ JSON formatted message")
	logger.Warn("⚠️ JSON warning message")
	logger.Error("❌ JSON error message")

	logger.Info("📁 Check logs/test.json for JSON output")
	_ = logger.Sync()

	logger.Configure(logger.WithJSON(false), logger.WithConsole(true))
}

func testCustomLogger() {
	buildHeader("Custom Logger Instances")

	consoleLog := logger.New(
		logger.WithLevel(logger.DebugLevel),
	)

	consoleLog.Trace("🔍 Custom logger - trace")
	consoleLog.Debug("🐛 Custom logger - debug")
	consoleLog.Info("✅ Custom logger - info")
	consoleLog.Warn("⚠️ Custom logger - warn")
	consoleLog.Error("❌ Custom logger - error")

	logPath, _ := filepath.Abs("logs/custom.log")
	fileLog := logger.New(
		logger.WithFile(logPath),
	)

	fileLog.Info("✅ Custom file logger - info")
	fileLog.Warn("⚠️ Custom file logger - warn")
	fileLog.Error("❌ Custom file logger - error")
	_ = fileLog.Sync()
}

func testLogLevelFiltering() {
	buildHeader("Log Level Filtering")

	logger.SetLevel(logger.WarnLevel)

	logger.Trace("🔍 This trace should NOT appear")
	logger.Debug("🐛 This debug should NOT appear")
	logger.Info("✅ This info should NOT appear")
	logger.Warn("⚠️ This warn SHOULD appear")
	logger.Error("❌ This error SHOULD appear")

	logger.SetLevel(logger.InfoLevel)

	logger.Trace("🔍 This trace should NOT appear")
	logger.Debug("🐛 This debug should NOT appear")
	logger.Info("✅ This info should NOT appear")
	logger.Warn("⚠️ This warn should NOT appear")
	logger.Error("❌ This error SHOULD appear")

	logger.SetLevel(logger.InfoLevel)
}

func testRequest() {
	buildHeader("Request Logging")
	logger.Request("GET", "/documents", 200, 1*time.Millisecond, "127.0.0.1")
	logger.Request("POST", "/api/users", 201, 150*time.Millisecond, "192.168.1.100")
	logger.Request("DELETE", "/api/users/123", 204, 50*time.Millisecond, "10.0.0.1")
	logger.Request("GET", "/not-found", 404, 10*time.Millisecond, "127.0.0.1")
	logger.Request("POST", "/error", 500, 200*time.Millisecond, "192.168.1.50")
	logger.Request("PUT", "/api/data", 200, 100*time.Millisecond, "172.16.0.100")
	logger.Request("PATCH", "/api/data", 200, 100*time.Millisecond, "172.16.0.100")
}

func testCrit() {
	buildHeader("Critical Logging")
	logger.Crit("💥 This crit SHOULD appear and program will exit")
}

func testRotate() {
	buildHeader("Log Rotation")
	data := make([]byte, 900*1024)
	_ = os.WriteFile("logs/rotated.log", data, 0644)
	rotatedLogger := logger.New(
		logger.WithFile("logs/rotated.log"),
		logger.WithConsole(false),
		logger.WithMaxSize(1), // 1MB
		logger.WithMaxBackups(3),
		logger.WithMaxAge(30),
		logger.WithCompress(true),
	)

	for range 1000 {
		rotatedLogger.Info("%s", strings.Repeat("A", 200))
	}

	_ = rotatedLogger.Sync()
}
