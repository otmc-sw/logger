/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
 **/
package core

import (
	"io"

	"github.com/otmc-sw/logger/writer"
)

// Writer defines the interface for log output destinations.
// Reuses writer.Writer to avoid duplication.
type Writer = writer.Writer

// NewConsoleWriter creates a new ConsoleWriter that writes to the given output.
// If out is nil, defaults to os.Stdout.
func NewConsoleWriter(out io.Writer) *writer.ConsoleWriter {
	return writer.NewConsoleWriter(out)
}

// NewMultiWriter creates a new MultiWriter that duplicates writes to all provided writers.
func NewMultiWriter(writers ...Writer) *writer.MultiWriter {
	return writer.NewMultiWriter(writers...)
}