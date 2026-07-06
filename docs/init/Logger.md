# OTMC Logger Library

Một thư viện logging dùng chung cho tất cả các dự án Go, được thiết kế theo hướng **đơn giản, hiệu năng cao, dễ sử dụng và ổn định**.

> **Design Philosophy**- API tối giản.
> - Sử dụng cú pháp giống `fmt.Printf`.
> - Không yêu cầu Structured Logging.
> - Tối ưu cho Console và File Logging.
> - Mặc định có format đẹp, không cần cấu hình.

---

# Mục tiêu

- API thống nhất cho mọi project.
- Dễ học, dễ sử dụng.
- Console Logging.
- File Logging.
- Log Rotation.
- JSON/Text Formatter.
- Middleware cho Web Framework.
- Hook (Webhook, Discord, Slack...).
- Có thể thay đổi logging engine mà không ảnh hưởng API.

---

# Cấu trúc project

```
otmc-logger/
│
├── go.mod
├── README.md
├── LICENSE
├── NOTICE
│
├── logger.go
├── config.go
├── options.go
├── level.go
├── formatter.go
├── hook.go
├── global.go
│
├── internal/
│   ├── core.go
│   ├── runtime.go
│   ├── encoder.go
│   ├── writer.go
│   ├── rotate.go
│   └── color.go
│
├── formatter/
│   ├── pretty.go
│   ├── text.go
│   └── json.go
│
├── middleware/
│   ├── gin.go
│   ├── fiber.go
│   ├── echo.go
│   └── chi.go
│
├── hooks/
│   ├── webhook.go
│   ├── discord.go
│   ├── slack.go
│   ├── telegram.go
│   └── email.go
│
├── examples/
└── tests/
```

---

# Public API

## Initialization

```
logger.Init(logger.Config{
    Level: logger.InfoLevel,
})
```

## Logging

```
logger.Trace("Loading configuration")
logger.Debug("Connecting to %s", host)
logger.Info("Server started on %s:%d", host, port)
logger.Warn("Memory usage: %d MB", memory)
logger.Error("Database error: %v", err)
logger.Crit("Application crashed: %v", err)
```

---

# API Design

Tất cả logging function đều sử dụng cùng một format.

```
logger.Info(format string, args ...any)
```

Ví dụ

```
logger.Info("Application Started")
logger.Info("Listening on %s:%d", host, port)
logger.Info("%d records loaded", total)
logger.Warn("Memory usage %.2f%%", percent)
logger.Error("Open file %s failed: %v", filename, err)
```

Bên trong logger:

```
message := fmt.Sprintf(format, args...)
```

Người dùng không cần gọi `fmt.Sprintf()`.

---

# Logger Interface

```
type Logger interface {
    Trace(format string, args ...any)
    Debug(format string, args ...any)
    Info(format string, args ...any)
    Warn(format string, args ...any)
    Error(format string, args ...any)
    Crit(format string, args ...any)
    Sync() error
}
```

---

# Functional Options

```
logger.New(
    logger.WithConsole(),
    logger.WithFile("logs/app.log"),
    logger.WithJSON(),
    logger.WithLevel(logger.DebugLevel),
)
```

---

# Config

```
type Config struct {
    Level Level
    Console bool
    File bool
    Filename string
    JSON bool
    Caller bool
    Stacktrace bool
    MaxSize int
    MaxBackups int
    MaxAge int
    Compress bool
}
```

---

# Default Console Format

```
2026-07-04 08:24:27.126 +07:00     Runner()             main.go:157   | INFO  | 🌐 Registering APIs...
```

Automatically includes:

- Timestamp
- Timezone
- Function
- File
- Line
- Log Level
- ANSI Color (Console)
- Pretty Padding

Người dùng chỉ cần:

```
logger.Info("🌐 Registering APIs...")
```

---

# Output

Hỗ trợ ghi đồng thời nhiều nơi.

```
Console
+File
+Webhook
```

Ví dụ

```
logger.New(
    logger.WithConsole(),
    logger.WithFile("logs/app.log"),
)
```

---

# Formatter

## Pretty (Default)

```
2026-07-04 08:24:27.126 +07:00     Runner()             main.go:157   | INFO  | 🌐 Registering APIs...
```

## Text

```
INFO Application Started
```

## JSON

```
{
    "time":"2026-07-04T08:24:27Z",
    "level":"INFO",
    "message":"Application Started"
}
```

---

# Log Rotation

Khuyến nghị sử dụng

```
lumberjack
```

Hỗ trợ

- Max Size
- Max Backups
- Max Age
- Compress

---

# Middleware

Hỗ trợ

- Gin
- Fiber
- Echo
- Chi

Ví dụ

```
r.Use(logger.GinMiddleware())
```

---

# Hooks

Có thể gửi log tới

- Discord
- Slack
- Telegram
- Webhook
- Email

Ví dụ

```
Application
↓
Logger
↓
Webhook
↓
Discord
```

---

# Log Levels

```
TRACE
DEBUG
INFO
NOTE
EVENT
WARN
ERROR
CRIT
```

---

# Ví dụ sử dụng

```
package main

import "github.com/otmc-sw/logger"

func main() {

    logger.Init(logger.Config{
        Level: logger.DebugLevel,
        Console: true,
        File: true,
        Filename: "logs/app.log",

    })

    logger.Info("🚀 Application Started")
    logger.Info("🌐 Listening on %s:%d", host, port)
    logger.Info("📦 Loaded %d modules", modules)
    logger.Warn("⚠️ Memory usage %d MB", memory)
    logger.Error("❌ Database error: %v", err)
    logger.Crit("💥 Critical failure: %v", err)
}
```

---

# Dependencies

| Package | Purpose |
| --- | --- |
| `go.uber.org/zap` | Logging Engine |
| `gopkg.in/natefinch/lumberjack.v2` | Log Rotation |

---

# Design Principles

- API đơn giản.
- Giống `fmt.Printf`.
- Không yêu cầu Structured Logging.
- Không expose `zap`.
- Thread-safe.
- High Performance.
- Zero-allocation tối đa (phụ thuộc backend).
- Functional Options.
- Production Ready.
- Dễ mở rộng Formatter.
- Dễ mở rộng Hooks.
- Dễ tích hợp vào mọi ứng dụng Go.

---

# Roadmap

## v1.0

- Console Logger
- File Logger
- Pretty Formatter
- Text Formatter
- JSON Formatter
- Log Rotation
- Global Logger
- Custom Logger
- ANSI Color
- Caller Information

## v1.1

- Context Logger
- Request ID
- Trace ID
- Gin Middleware
- Fiber Middleware
- Echo Middleware
- Chi Middleware

## v1.2

- Webhook Hook
- Discord Hook
- Slack Hook
- Telegram Hook
- Email Hook
- Async Logging

## v2.0

- OpenTelemetry
- Metrics
- Sampling
- Plugin System
- Custom Formatter
- Dynamic Configuration
- Dashboard Integration