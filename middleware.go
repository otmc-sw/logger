/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package logger

// Middleware provides logging middleware for web frameworks
// This package contains middleware implementations for:
// - Gin (middleware/gin.go)
// - Fiber (middleware/fiber.go)
// - Echo (middleware/echo.go)
// - Chi (middleware/chi.go)
//
// Each middleware automatically logs HTTP requests with:
// - Request method and path
// - Response status code
// - Request duration
// - Request ID (if available)
// - Client IP
//
// Example usage with Gin:
//
//	import "github.com/otmc-sw/logger/middleware/gin"
//
//	router.Use(gin.Logger())
//
// Example usage with Fiber:
//
//	import "github.com/otmc-sw/logger/middleware/fiber"
//
//	app.Use(fiber.Logger())
