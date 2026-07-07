/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
 **/
package rotator

import "errors"

var (
	ErrClosed          = errors.New("rotator: writer is closed")
	ErrInvalidFilename = errors.New("rotator: invalid filename")
	ErrInvalidOption   = errors.New("rotator: invalid option")
)
