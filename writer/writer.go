/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
 **/
package writer

import "io"

// Writer defines the interface for log output destinations.
type Writer interface {
	io.Writer
	Sync() error
}