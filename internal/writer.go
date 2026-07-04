/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package internal

import (
	"io"
	"os"
)

// ConsoleWriter writes to console
type ConsoleWriter struct {
	out io.Writer
}

// NewConsoleWriter creates a new console writer
func NewConsoleWriter(out io.Writer) *ConsoleWriter {
	if out == nil {
		out = os.Stdout
	}
	return &ConsoleWriter{out: out}
}

// Write writes to the console
func (w *ConsoleWriter) Write(p []byte) (n int, err error) {
	return w.out.Write(p)
}

// Sync is a no-op for console writer
func (w *ConsoleWriter) Sync() error {
	return nil
}

// MultiWriter writes to multiple writers
type MultiWriter struct {
	writers []Writer
}

// NewMultiWriter creates a new multi writer
func NewMultiWriter(writers ...Writer) *MultiWriter {
	return &MultiWriter{writers: writers}
}

// Write writes to all writers
func (w *MultiWriter) Write(p []byte) (n int, err error) {
	for _, writer := range w.writers {
		_, err = writer.Write(p)
		if err != nil {
			return
		}
	}
	return len(p), nil
}

// Sync syncs all writers
func (w *MultiWriter) Sync() error {
	for _, writer := range w.writers {
		if err := writer.Sync(); err != nil {
			return err
		}
	}
	return nil
}
