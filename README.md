# 📚 OTMC Logger

A simple, high-performance logging library for Go applications, designed for CLI, Desktop, API, and Server applications.

## ✨ Features

- **Simple API** - Similar to `fmt.Printf`, easy to use
- **Beautiful Console Output** - Automatic formatting with colors and alignment
- **File Logging** - With automatic rotation support
- **Multiple Formatters** - Pretty (default), Text, and JSON
- **Log Levels** - TRACE, DEBUG, INFO, WARN, ERROR, CRIT
- **Caller Information** - Automatically captures function, file, and line number
- **Runtime Configuration** - Change settings at runtime without restart
- **Multiple Logger Instances** - Create independent loggers with different configs
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

### 🛠️ Creating a Logger

```go
// With no options — uses default config (console, InfoLevel, caller on)
log := logger.New()

// With one or more options
log := logger.New(
    logger.WithLevel(logger.DebugLevel),
    logger.WithConsole(),
    logger.WithFile("logs/app.log"),
    logger.WithJSON(),
    logger.WithCaller(false),
    logger.WithMaxSize(20),
    logger.WithMaxBackups(3),
    logger.WithMaxAge(30),
    logger.WithCompress(true),
    logger.WithTimeFormat(time.RFC3339),
)
```

### 🔄 Runtime Configuration

The logger supports runtime configuration changes. Most changes can be done
through `Configure()` or dedicated methods:

#### Configure (batch update)

Apply one or more options at once. This is the recommended way to change
settings that affect the formatter or writer (JSON toggle, file path, etc.).

```go
log.Configure(
    logger.WithJSON(),
    logger.WithFile("logs/app.log"),
    logger.WithCaller(false),
)

// Global
logger.Configure(
    logger.WithLevel(logger.DebugLevel),
    logger.WithJSON(),
)
```

#### SetLevel (single update — no rebuild)

Changing the log level is the most frequent operation and does not require
recreating the logger internals.

```go
log.SetLevel(logger.DebugLevel)

// Global
logger.SetLevel(logger.WarnLevel)
```

#### Config / Update (round-trip)

Get a copy of the current config, modify it, and apply it back.

```go
cfg := log.Config()
cfg.Level = logger.DebugLevel
cfg.JSON = true
cfg.Filename = "logs/app.log"
log.Update(cfg)

// Global
cfg := logger.GetConfig()
cfg.Console = true
logger.Update(cfg)
```

### Available Options

| Option | Default | Description |
|---|---|---|
| `WithLevel(level)` | `InfoLevel` | Minimum log level |
| `WithConsole(enabled)` | `true` | Enable/disable console output |
| `WithFile(filename)` | `""` | Enable file output with rotation |
| `WithJSON(enabled)` | `false` | JSON format instead of pretty |
| `WithCaller(enabled)` | `true` | Include caller info (file:line) |
| `WithStacktrace(enabled)` | `false` | Include stack trace on errors |
| `WithMaxSize(mb)` | `10` | Max file size before rotation (MB) |
| `WithMaxBackups(n)` | `3` | Max rotated files to keep |
| `WithMaxAge(days)` | `90` | Max age of rotated files (days) |
| `WithCompress(enabled)` | `true` | Compress rotated files |
| `WithTimeFormat(format)` | `"2006-01-02 15:04:05.000 -07:00"` | Time format string |

## 📊 Log Levels

```go
logger.Trace("Detailed trace information")
logger.Debug("Debug information")
logger.Info("General informational messages")
logger.Warn("Warning messages")
logger.Error("Error messages")
logger.Crit("Critical errors — exits after logging")
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

The library provides a global logger instance that mirrors the instance API:

```go
// Logging
logger.Trace("...")
logger.Debug("...")
logger.Info("...")
logger.Warn("...")
logger.Error("...")
logger.Crit("...")

// Request logging
logger.Request("GET", "/api/users", 200, 150*time.Millisecond, "10.0.0.1")

// Configuration
logger.SetLevel(logger.DebugLevel)
logger.Configure(logger.WithJSON(), logger.WithFile("app.log"))
logger.Update(cfg)
cfg := logger.GetConfig()
```

## 🔧 Custom Logger Instances

Create independent logger instances with different configurations:

```go
// Console logger — debug level, pretty format
consoleLog := logger.New(
    logger.WithConsole(),
    logger.WithLevel(logger.DebugLevel),
)
consoleLog.Info("This is a console log")
consoleLog.SetLevel(logger.TraceLevel)

// File logger — info level, rotated files
fileLog := logger.New(
    logger.WithFile("logs/app.log"),
    logger.WithLevel(logger.InfoLevel),
    logger.WithMaxSize(20),
)
fileLog.Info("This is a file log")

// JSON logger — for structured monitoring
jsonLog := logger.New(
    logger.WithFile("logs/metrics.json"),
    logger.WithJSON(),
    logger.WithCaller(false),
)
jsonLog.Info("This is a JSON log for metrics")

// Change settings at runtime
jsonLog.Configure(logger.WithCaller(true))
```

## ♻️ Log Rotation

Automatic log rotation using the built-in rotator package:

```go
log := logger.New(
    logger.WithFile("logs/app.log"),
    logger.WithMaxSize(20),    // 20 MB per file
    logger.WithMaxBackups(3),  // keep 3 rotated files
    logger.WithMaxAge(30),     // keep for 30 days
    logger.WithCompress(true), // compress rotated files
)
```

The rotator package is lightweight, dependency-free, and supports:
- Configurable naming strategies (index, date, timestamp, or custom)
- Gzip compression
- Automatic cleanup by count and age
- Thread-safe concurrent writes

For advanced usage, see the [rotator package documentation](https://github.com/otmc-sw/logger/tree/main/rotator).

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
Logger API (Trace, Debug, Info, Warn, Error, Crit, Request)
    ↓
┌─────────────────────────────────────┐
│ Logger struct                       │
│   ┌─────────┐  ┌──────────────────┐ │
│   │ Config  │  │ Core             │ │
│   │ Level   │  │   → Formatter    │ │
│   │ JSON    │  │   → Writer       │ │
│   │ Console │  │   → Hooks        │ │
│   │ File    │  └──────────────────┘ │
│   │ Caller  │                       │
│   └─────────┘                       │
│                                     │
│   Configure()  SetLevel()           │
│   Config()     Update()             │
└─────────────────────────────────────┘
    ↓
├── Formatter (Pretty, Text, JSON)
├── Writer (Console, File, Multi)
└── Hooks (Webhook, Discord, Slack, etc.)
```

## 📄 License

* Apache License 2.0
* Copyright (c) 2026 OTMC Softwares.

## ✨ Contributors

* 🌿 Nguyen Van Trung
* 🌿 Nguyen Thi Hoai
* 🌿 OTMC Contributors