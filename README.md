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
    logger.Init(logger.Config{
        Level:    logger.InfoLevel,
        Console:  true,
        File:     true,
        Filename: "logs/app.log",
        Caller:   true,
    })
    err := "OTMC Testing Message"

	logger.Trace("🚀 Starting application...")
	logger.Debug("📝 Configuration loaded from %s", "config.yaml")
	logger.Info("🌐 Server listening on %s:%d", "localhost", 8080)
	logger.Warn("⚠️ Memory usage is high: %.1f%%", 85.5)
	logger.Error("❌ Failed to connect to database: %s", "postgres")
	logger.Crit("❌ Unable to initialize application: %v", err)
}
```

## 🖥️ Console Output

The logger automatically formats output with timestamps, caller information, and colors:

```
2026-07-04 15:49:40.404 +07:00                 Info()       global.go:20    | INFO  | ✅ Application started
2026-07-04 15:49:40.404 +07:00                 Info()       global.go:20    | INFO  | 🌐 Server listening on localhost:8080
2026-07-04 15:49:40.405 +07:00                 Warn()       global.go:25    | WARN  | ⚠️ Memory usage is high 85.5
```

## ⚙️ Configuration

### 🔧 Using Config

```go
logger.Init(logger.Config{
    Level:      logger.DebugLevel,
    Console:    true,
    File:       true,
    Filename:   "logs/app.log",
    JSON:       false,
    Caller:     true,
    MaxSize:    100,    // MB
    MaxBackups: 3,
    MaxAge:     30,     // days
    Compress:   true,
})
```

### 🛠️ Using Functional Options

```go
log := logger.New(
    logger.WithLevel(logger.DebugLevel),
    logger.WithConsole(),
    logger.WithFile("logs/app.log"),
    logger.WithJSON(),
    logger.WithCaller(true),
    logger.WithMaxSize(100),
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
2026-07-04 15:49:40.404 +07:00                 Info()       main.go:20    | INFO  | Application started
```

### 📝 Text Formatter

```go
log := logger.New(
    logger.WithConsole(),
)
// Use TextFormatter by setting it on the core
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

// Configure global logger
logger.Init(logger.Config{
    Level: logger.DebugLevel,
})

// Set log level
logger.SetLevel(logger.WarnLevel)
```

## 🔧 Custom Logger

Create multiple logger instances with different configurations:

```go
// Console logger
consoleLog := logger.New(
    logger.WithConsole(),
    logger.WithLevel(logger.DebugLevel),
)

// File logger
fileLog := logger.New(
    logger.WithFile("logs/app.log"),
    logger.WithLevel(logger.InfoLevel),
)

// JSON logger for monitoring
jsonLog := logger.New(
    logger.WithFile("logs/metrics.log"),
    logger.WithJSON(),
)
```

## ♻️ Log Rotation

Automatic log rotation using lumberjack:

```go
logger.Init(logger.Config{
    File:       true,
    Filename:   "logs/app.log",
    MaxSize:    100,    // Max size in MB
    MaxBackups: 3,      // Max number of old log files
    MaxAge:     30,     // Max age in days
    Compress:   true,   // Compress rotated files
})
```

## 🎨 Color Output

Colors are automatically applied to console output:

- ⚪ **TRACE** - Gray
- 🔵 **DEBUG** - Blue
- 🟢 **INFO** - Green
- 🟡 **WARN** - Yellow
- 🔴 **ERROR** - Red
- 🔴 **CRIT** - Bright Red

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

## ©️ Copyright

Copyright (c) 2026 OTMC Softwares

## ✨ Contributors

- 🌿 Nguyen Van Trung
- 🌿 Nguyen Thi Hoai
- 🌿 OTMC Contributors