/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
 **/
package writer

import "io"

type ConsoleWriter struct {
	out io.Writer
}

func NewConsoleWriter(out io.Writer) *ConsoleWriter {
	return &ConsoleWriter{out: out}
}

func (w *ConsoleWriter) Write(p []byte) (n int, err error) {
	if w.out != nil {
		return w.out.Write(p)
	}
	return len(p), nil
}

func (w *ConsoleWriter) Sync() error {
	return nil
}

func (w *ConsoleWriter) Close() error {
	return nil
}
