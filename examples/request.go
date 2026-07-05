/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package examples

import (
	"time"

	"github.com/otmc-sw/logger"
)

func requestExample() {
	config := logger.Config{
		Level:   logger.InfoLevel,
		Console: true,
		Caller:  false,
	}
	logger.Init(config)

	logger.Request("GET", "/documents", 200, 0, "127.0.0.1")
	logger.Request("POST", "/api/users", 201, 150*time.Millisecond, "192.168.1.100")
	logger.Request("DELETE", "/api/users/123", 204, 50*time.Millisecond, "10.0.0.1")
	logger.Request("GET", "/not-found", 404, 10*time.Millisecond, "127.0.0.1")
	logger.Request("POST", "/error", 500, 200*time.Millisecond, "192.168.1.50")
	logger.Request("PUT", "/api/data", 200, 100*time.Millisecond, "172.16.0.100")
}

func init() {
	requestExample()
}
