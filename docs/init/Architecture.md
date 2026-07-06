# OTMC Logger Architecture Summary

## Project Structure

```
logger/
│
├── go.mod
├── README.md
├── LICENSE
│
├── logger.go          # Public API
├── config.go          # Configuration
├── options.go         # Functional Options
├── level.go           # Log Levels
├── field.go           # Structured Fields
├── formatter.go       # Formatter Interface
├── hook.go            # Hook Interface
├── middleware.go      # Middleware Entry
├── global.go          # Global Logger
│
├── internal/
│   ├── core.go        # Logger implementation
│   ├── runtime.go     # Caller information
│   ├── encoder.go     # Console/JSON encoder
│   ├── writer.go      # Multi Writer
│   ├── rotate.go      # File Rotation
│   └── color.go       # ANSI Color
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

# Architecture

```
                 Application
                     │
         logger.Info / Warn / Error
                     │
             Global Logger API
                     │
               Logger Interface
                     │
              Logger Core Engine
     ┌───────────────┼────────────────┐
     │               │                │
     ▼               ▼                ▼
 Formatter       Multi Writer      Hooks
     │               │                │
     │               │                ├──────── Discord
     │               │                ├──────── Slack
     │               │                ├──────── Telegram
     │               │                └──────── Webhook
     │               │
     ▼               ▼
 Pretty        Console Writer
 Text          File Writer
 JSON          Rotate Writer
```
---

# Core Components

| Component | Responsibility |
| --- | --- |
| Logger | Public logging API |
| Core | Process log entries |
| Formatter | Format output (Pretty, Text, JSON) |
| Writer | Write logs to Console/File |
| Rotation | Rotate log files |
| Runtime | Get Function/File/Line |
| Hook | Send logs to external services |
| Middleware | Web framework integration |

---

# Public API

## Initialization

```
logger.Init(...)
```

## Create Logger

```
logger.New(...)
```

## Logging

```
logger.Trace(...)
logger.Debug(...)
logger.Info(...)
logger.Warn(...)
logger.Error(...)
logger.Crit(...)​
```

## Configuration

```
logger.WithLevel(...)
logger.WithConsole(...)
logger.WithFile(...)
logger.WithJSON(...)
logger.WithFormatter(...)
```

---

# Log Flow

```
Application
      │
logger.Info(...)
      │
Collect Runtime Information
      │
Build Log Entry
      │
Apply Formatter
      │
Write Outputs
      │
Execute Hooks
```

---

# Log Entry

```
type Entry struct {
    Time
    Level
    Function
    File
    Line
    Message
}
```

---

# Default Output

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
- ANSI Color (Console only)

---

# Supported Outputs

- Console
- File
- Multi Writer
- Rotating File

---

# Supported Formatters

- Pretty (Default)
- Text
- JSON

---

# Supported Hooks

- Webhook
- Discord
- Slack
- Telegram
- Email

---

# Supported Middleware

- Gin
- Fiber
- Echo
- Chi

---

# Design Principles

- Simple and consistent API
- Structured Logging
- High performance
- Thread-safe
- Functional Options
- Interface-based architecture
- Extensible through Formatter, Writer and Hook
- No direct dependency exposed to users
- Production-ready by default