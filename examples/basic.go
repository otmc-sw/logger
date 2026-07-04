/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package main

import (
	"os"
	"path/filepath"

	"github.com/otmc-sw/logger"
)

func main() {
	// Create logs directory
	_ = os.MkdirAll("logs", 0755)

	// Get absolute path for log file
	logPath, _ := filepath.Abs("logs/app.log")

	// Initialize logger with default configuration
	logger.Init(logger.Config{
		Level:    logger.DebugLevel,
		Console:  true,
		File:     true,
		Filename: logPath,
		Caller:   true,
	})

	logger.Trace("🔍 This is a trace message")
	logger.Debug("🐛 Debugging information")
	logger.Info("✅ Application started successfully")
	logger.Warn("⚠️ This is a warning message")
	logger.Error("❌ An error occurred")
	logger.Crit("💥 Critical failure - shutting down")

	// Example with formatting
	logger.Info("🌐 Server listening on %s:%d", "localhost", 8080)
	logger.Info("📦 Loaded %d modules", 42)
	logger.Warn("⚠️ Memory usage: %.2f%%", 85.5)

	// Example with error
	err := someOperation()
	if err != nil {
		logger.Error("❌ Operation failed: %v", err)
	}

	logger.Info("🎉 All done!")
	_ = logger.Sync()
}

func someOperation() error {
	return nil
}
