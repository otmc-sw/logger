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

type Writer = writer.Writer

func NewConsoleWriter(out io.Writer) *writer.ConsoleWriter {
	return writer.NewConsoleWriter(out)
}

func NewMultiWriter(writers ...Writer) *writer.MultiWriter {
	return writer.NewMultiWriter(writers...)
}

func NewFileWriter(filename string) (*writer.FileWriter, error) {
	return writer.NewFileWriter(filename)
}
