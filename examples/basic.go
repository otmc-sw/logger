/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package examples

import "github.com/otmc-sw/logger"

func basic() {
	logger.Trace("🔍 This is a trace message")
	logger.Debug("👽 Debugging information")
	logger.Info("✅ Application started successfully")
	logger.Warn("⚠️ This is a warning message")
	logger.Error("❌ An error occurred")
	logger.Crit("🔥 Critical failure - shutting down")
}

func init() {
	basic()
}
