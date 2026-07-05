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

type ConsoleWriter struct {
	out io.Writer
}

func NewConsoleWriter(out io.Writer) *ConsoleWriter {
	if out == nil {
		out = os.Stdout
	}
	return &ConsoleWriter{out: out}
}

func (w *ConsoleWriter) Write(p []byte) (n int, err error) {
	return w.out.Write(p)
}

func (w *ConsoleWriter) Sync() error {
	return nil
}

type MultiWriter struct {
	writers []Writer
}

func NewMultiWriter(writers ...Writer) *MultiWriter {
	return &MultiWriter{writers: writers}
}

func (w *MultiWriter) Write(p []byte) (n int, err error) {
	for _, writer := range w.writers {
		_, err = writer.Write(p)
		if err != nil {
			return
		}
	}
	return len(p), nil
}

func (w *MultiWriter) Sync() error {
	for _, writer := range w.writers {
		if err := writer.Sync(); err != nil {
			return err
		}
	}
	return nil
}
