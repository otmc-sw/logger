/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package internal

import (
	"gopkg.in/natefinch/lumberjack.v2"
)

type RotateWriter struct {
	lumberjack *lumberjack.Logger
}

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

func (w *RotateWriter) Write(p []byte) (n int, err error) {
	stripped := StripColorCodes(string(p))
	return w.lumberjack.Write([]byte(stripped))
}

func (w *RotateWriter) Sync() error {
	return w.lumberjack.Rotate()
}

func (w *RotateWriter) Close() error {
	return w.lumberjack.Close()
}
