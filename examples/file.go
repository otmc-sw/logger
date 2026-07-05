/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package examples

import (
	"os"

	"github.com/otmc-sw/logger"
)

func file() {

	_ = os.MkdirAll("logs", 0755)

	logger.Init(logger.Config{
		Level:    logger.DebugLevel,
		Console:  true,
		File:     true,
		Filename: "logs/app.log",
		Caller:   true,
	})

	logger.Trace("ðŸ” This is a trace message")
	logger.Debug("ðŸ› Debugging information")
	logger.Info("âœ… Application started successfully")
	logger.Warn("âš ï¸ This is a warning message")
	logger.Error("âŒ An error occurred")
	logger.Info("ðŸŽ‰ All done!")
	_ = logger.Sync()
}

func init() {
	file()
}
