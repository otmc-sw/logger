# Rotator

A lightweight, dependency-free log rotation library for Go.

## Features

- **Zero external dependencies** - Uses only Go standard library
- **Thread-safe** - Concurrent writes are safe
- **Configurable naming strategies** - Built-in and custom naming for rotated files
- **Gzip compression** - Optional compression for rotated logs
- **Automatic cleanup** - Remove old backups by count or age
- **Cross-platform** - Works on Windows, Linux, and macOS

## Installation

```bash
go get github.com/otmc-sw/logger/rotator
```

## Quick Start

```go
package main

import (
    "github.com/otmc-sw/logger/rotator"
)

func main() {
    r := rotator.New(
        rotator.WithFilename("logs/app.log"),
        rotator.WithMaxSize(10),          // 10 MB
        rotator.WithMaxBackups(5),
        rotator.WithMaxAge(30),           // 30 days
        rotator.WithCompress(true),
    )
    defer r.Close()

    r.Write([]byte("Hello, World!\n"))
}
```

## Configuration Options

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `WithFilename` | string | - | Path to the log file (required) |
| `WithMaxSize` | int | 100 | Maximum file size in MB before rotation |
| `WithMaxBackups` | int | 3 | Maximum number of backup files to keep |
| `WithMaxAge` | int | 30 | Maximum age in days for backup files |
| `WithCompress` | bool | false | Enable gzip compression for rotated files |
| `WithTimeFormat` | string | "20060102" | Time format for naming (see naming strategies) |
| `WithNaming` | NamingFunc | NameWithIndex | Custom naming function for rotated files |

## Naming Strategies

### Built-in Strategies

#### Index (Default)
```go
rotator.WithNaming(rotator.NameWithIndex)
```
Produces: `app_1.log`, `app_2.log`, `app_3.log`

#### Date + Index
```go
rotator.WithNaming(rotator.NameWithDateAndIndex)
```
Produces: `app_20251021_1.log`, `app_20251021_2.log`

#### Timestamp
```go
rotator.WithNaming(rotator.NameWithTimestamp)
```
Produces: `app_20251021_153015.log`

#### Timestamp + Index
```go
rotator.WithNaming(rotator.NameWithTimestampAndIndex)
```
Produces: `app_20251021_153015_1.log`

### Custom Naming

```go
rotator.WithNaming(func(info rotator.RotateInfo) string {
    return fmt.Sprintf(
        "%s_%s_%03d%s",
        info.BaseName,
        info.Time.Format("20060102"),
        info.Index,
        info.Extension,
    )
})
```
Produces: `app_20251021_001.log`

## RotateInfo Structure

```go
type RotateInfo struct {
    BaseName  string    // Base filename without extension
    Extension string    // File extension (e.g., ".log")
    Index     int       // Current rotation index
    Time      time.Time // Rotation time
}
```

## API Reference

### Methods

- `Write(p []byte) (n int, err error)` - Write data to the log file
- `Close() error` - Close the log file
- `Rotate() error` - Force a manual rotation
- `Sync() error` - Flush the file to disk
- `Size() int64` - Get current file size in bytes

### Example: Manual Rotation

```go
r := rotator.New(rotator.WithFilename("app.log"))
defer r.Close()

// Write some data
r.Write([]byte("Initial data\n"))

// Force rotation
err := r.Rotate()
if err != nil {
    log.Fatal(err)
}

// Continue writing to new file
r.Write([]byte("New file data\n"))
```

## Thread Safety

The rotator is thread-safe and can be used concurrently from multiple goroutines.

```go
r := rotator.New(rotator.WithFilename("app.log"))
defer r.Close()

var wg sync.WaitGroup
for i := 0; i < 10; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        r.Write([]byte(fmt.Sprintf("Goroutine %d\n", id)))
    }(i)
}
wg.Wait()
```

## Rotation Flow

```
Write()
  ↓
Check if size exceeds MaxSize
  ↓
Yes? → Close current file
  ↓
Rename to backup (using naming strategy)
  ↓
Compress (if enabled)
  ↓
Cleanup old backups (by MaxBackups and MaxAge)
  ↓
Open new file
  ↓
Continue writing
```

## Cleanup Rules

- **MaxBackups**: Keeps only the N most recent backup files
- **MaxAge**: Removes backups older than N days
- Both rules are applied together

## License

Apache License 2.0

## Contributing

This package is part of the OTMC Logger project. See [github.com/otmc-sw/logger](https://github.com/otmc-sw/logger) for more information.
