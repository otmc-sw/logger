# 📚 OTMC Logger

A simple, high-performance logging library for Go applications, designed for CLI, Desktop, API, and Server applications.

## ✨ Features

- **Simple API** - Similar to `fmt.Printf`, easy to use
- **Beautiful Console Output** - Automatic formatting with colors and alignment
- **File Logging** - With automatic rotation support
- **Multiple Formatters** - Pretty (default), Text, and JSON
- **Log Levels** - TRACE, DEBUG, INFO, WARN, ERROR, CRIT
- **Caller Information** - Automatically captures function, file, and line number
- **Thread-Safe** - Safe for concurrent use
- **Production Ready** - Default configuration optimized for production

## 📦 Installation

```bash
go get github.com/otmc-sw/logger
```

## 🚀 Quick Start

```go
package main

import (
    "github.com/otmc-sw/logger"
)

func main() {
    err := "OTMC Testing Message"

	logger.Trace("🚀 Starting application...")
	logger.Debug("📝 Configuration loaded from %s", "config.yaml")
	logger.Info("🌐 Server listening on %s:%d", "localhost", 8080)
	logger.Warn("⚠️ Memory usage is high: %.1f%%", 85.5)
	logger.Error("❌ Failed to connect to database: %s", "postgres")
    // Warning: This will close the application.
	logger.Crit("❌ Unable to initialize application: %v", err)
}
```

## 🖥️ Console Output

The logger automatically formats output with timestamps, caller information, and colors:

```
2026-07-06 08:57:34.208 +07:00     Initializer()      main.go:104   | INFO  | ✅ Database connection established.
2026-07-06 08:57:34.208 +07:00     Initializer()      main.go:106   | INFO  | 🧩 Loading settings...
2026-07-06 08:57:34.209 +07:00     Initializer()      main.go:112   | INFO  | 🤖 LLM Provider: groq
2026-07-06 08:57:34.209 +07:00     Initializer()      main.go:119   | INFO  | ✅ LLM client initialized.
2026-07-06 08:57:34.209 +07:00     Initializer()      main.go:149   | INFO  | 📁 Run directory: D:\SCM\GitHub\OTMC\Softwares\document-hub\backend
2026-07-06 08:57:34.209 +07:00     Initializer()      main.go:150   | INFO  | 📦 Frontend dist: D:\SCM\GitHub\OTMC\Softwares\document-hub\frontend\dist
2026-07-06 08:57:34.209 +07:00          Runner()      main.go:154   | INFO  | 🌿 Running application ...
2026-07-06 08:57:34.209 +07:00          Runner()      main.go:157   | INFO  | 🌐 Registering APIs ...
2026-07-06 08:57:34.210 +07:00           Start() scheduler.go:57    | INFO  | ✅ Backup scheduler engine started
2026-07-06 08:57:34.210 +07:00          reload() scheduler.go:142   | INFO  | 🔀 Backup scheduled with cron expression: 0 2 * * * (max backups: 10)
```

## ⚙️ Configuration

### 🛠️ Using Functional Options

```go
log := logger.New(
    logger.WithLevel(logger.DebugLevel),
    logger.WithConsole(),
    logger.WithFile("logs/app.log"),
    logger.WithJSON(),
    logger.WithCaller(true),
    logger.WithMaxSize(20),
    logger.WithMaxBackups(3),
    logger.WithMaxAge(30),
    logger.WithCompress(true),
)
```

## 📊 Log Levels

```go
logger.Trace("Detailed trace information")
logger.Debug("Debug information")
logger.Info("General informational messages")
logger.Warn("Warning messages")
logger.Error("Error messages")
logger.Crit("Critical errors (exits after logging)")
```

## 🎨 Formatters

### 🎯 Pretty Formatter (Default)

```go
log := logger.New(
    logger.WithConsole(),
)
```

Output:
```
2026-07-06 08:57:34.208 +07:00     Initializer()      main.go:106   | INFO  | 🧩 Loading settings...
2026-07-06 08:57:34.209 +07:00          Runner()      main.go:154   | INFO  | 🌿 Running application ...
```

### 📝 Text Formatter

```go
log := logger.New(
    logger.WithConsole(),
)
```

### 📋 JSON Formatter

```go
log := logger.New(
    logger.WithJSON(),
    logger.WithConsole(),
)
```

Output:
```json
{
  "time": "2026-07-04T08:49:40.404+07:00",
  "level": "INFO",
  "message": "Application started",
  "function": "main",
  "file": "main.go",
  "line": 20
}
```

## 🌍 Global Logger

The library provides a global logger instance for convenience:

```go
// Use global functions
logger.Info("Message")
logger.Error("Error: %v", err)
logger.SetLevel(logger.WarnLevel)
```

## 🔧 Custom Logger

Create multiple logger instances with different configurations:

```go
// Global logger
logger.SetLevel(logger.DebugLevel)

logger.Info("This is a global log")

// Console logger
consoleLog := logger.New(
    logger.WithConsole(),
    logger.WithLevel(logger.DebugLevel),
)

consoleLog.Info("This is a console log")

// File logger
fileLog := logger.New(
    logger.WithFile("logs/app.log"),
    logger.WithLevel(logger.InfoLevel),
)

fileLog.Info("This is a file log")

// JSON logger for monitoring
jsonLog := logger.New(
    logger.WithFile("logs/metrics.log"),
    logger.WithJSON(),
)

jsonLog.Info("This is a JSON log")
```

## ♻️ Log Rotation

Automatic log rotation using lumberjack:

```go
logger.WithMaxSize(20)
logger.WithMaxBackups(3)
logger.WithMaxAge(30)
logger.WithCompress(true)
```

## 🎨 Color Output

Colors are automatically applied to console output:

- **TRACE** - Gray
- **DEBUG** - Cyan
- **INFO**  - Blue
- **WARN**  - Yellow
- **ERROR** - Red
- **CRIT**  - Bright Red

Colors are automatically stripped when writing to files.

## 🏗️ Architecture

```
Application
    ↓
Logger API (Trace, Debug, Info, Warn, Error, Crit)
    ↓
Core Engine
    ↓
├── Formatter (Pretty, Text, JSON)
├── Writer (Console, File, Multi)
└── Hooks (Webhook, Discord, Slack, etc.)
```

## 📄 License

Apache License 2.0
Copyright (c) 2026 OTMC Softwares.

## ✨ Contributors

- 🌿 Nguyen Van Trung
- 🌿 Nguyen Thi Hoai
- 🌿 OTMC Contributors

