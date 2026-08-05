

## 📦 Installation

```bash
go get github.com/otmc-sw/logger@latest
```


## 🚀 Quick Start

```go
package main

import (
    "github.com/otmc-sw/logger"
)

func SetupLogger() {
	logFile := filepath.Join(DIR_RUN, "data", "logs", "app.log")

	logger.Configure(
		logger.WithFile(logFile),
	)

	if FLAG_DEBUG {
		logger.SetLevel(logger.DebugLevel)
	}

	logger.Info("✅ Logger initialized successfully.")
}

func main() {
    SetupLogger()

	logger.Trace("🚀 Starting application...")
	logger.Debug("📝 Configuration loaded from %s", "config.yaml")
	logger.Info("🌐 Server listening on %s:%d", "localhost", 8080)
	logger.Warn("⚠️ Memory usage is high: %.1f%%", 85.5)
	logger.Error("❌ Failed to connect to database: %s", "postgres")
    // Warning: This will close the application.
	logger.Crit("❌ Unable to initialize application")
}
```



