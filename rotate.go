/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package logger

import (
	"gopkg.in/natefinch/lumberjack.v2"
)

// RotateWriter wraps lumberjack for log rotation
type RotateWriter struct {
	lumberjack *lumberjack.Logger
}

// NewRotateWriter creates a new rotate writer
func NewRotateWriter(filename string, maxSize, maxBackups, maxAge int, compress bool) *RotateWriter {
	lj := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}
	return &RotateWriter{lumberjack: lj}
}

// Write writes to the file
func (w *RotateWriter) Write(p []byte) (n int, err error) {
	return w.lumberjack.Write(p)
}

// Sync flushes the file
func (w *RotateWriter) Sync() error {
	return w.lumberjack.Rotate()
}

// Close closes the file
func (w *RotateWriter) Close() error {
	return w.lumberjack.Close()
}
