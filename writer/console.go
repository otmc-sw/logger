/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
 **/
package writer

import "io"

// ConsoleWriter writes log output to the standard console.
type ConsoleWriter struct {
	out io.Writer
}

// NewConsoleWriter creates a new ConsoleWriter.
// If out is nil, output is discarded (caller should set it).
func NewConsoleWriter(out io.Writer) *ConsoleWriter {
	return &ConsoleWriter{out: out}
}

// Write implements the io.Writer interface.
func (w *ConsoleWriter) Write(p []byte) (n int, err error) {
	if w.out != nil {
		return w.out.Write(p)
	}
	return len(p), nil
}

// Sync implements the Writer interface (no-op for console).
func (w *ConsoleWriter) Sync() error {
	return nil
}