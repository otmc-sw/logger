/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package examples

import "github.com/otmc-sw/logger"

func basic() {
	logger.Trace("ðŸ” This is a trace message")
	logger.Debug("ðŸ› Debugging information")
	logger.Info("âœ… Application started successfully")
	logger.Warn("âš ï¸ This is a warning message")
	logger.Error("âŒ An error occurred")
	logger.Crit("ðŸ’¥ Critical failure - shutting down")
}

func init() {
	basic()
}
