/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package internal

import "github.com/otmc-sw/logger/rotator"

type RotateWriter struct {
	rotator *rotator.Rotator
}

func NewRotateWriter(filename string, maxSize, maxBackups, maxAge int, compress bool) *RotateWriter {
	r := rotator.New(
		rotator.WithFilename(filename),
		rotator.WithMaxSize(maxSize),
		rotator.WithMaxBackups(maxBackups),
		rotator.WithMaxAge(maxAge),
		rotator.WithCompress(compress),
	)
	return &RotateWriter{rotator: r}
}

func (w *RotateWriter) Write(p []byte) (n int, err error) {
	stripped := StripColorCodes(string(p))
	return w.rotator.Write([]byte(stripped))
}

func (w *RotateWriter) Sync() error {
	return w.rotator.Rotate()
}

func (w *RotateWriter) Close() error {
	return w.rotator.Close()
}
