/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/
package logger

// Middleware provides logging middleware supporting both net/http and fasthttp.
// The core implementation lives in middleware/core.go and supports both
// HTTP libraries in a single file.
//
// net/http support — works with Chi, Gin, Echo, and standard mux:
//   - Chi:          r.Use(middleware.Logger())
//   - Gin:          router.Use(gin.WrapH(middleware.Logger()(nextHandler)))
//   - Echo:         e.Use(echo.WrapMiddleware(middleware.Logger()))
//   - net/http mux: handler = middleware.Logger()(mux)
//
// fasthttp support — works with Fiber, FastHTTP, and other fasthttp-based frameworks:
//   - Fiber:        c := fasthttputil.NewRequestCtx(...) or use adaptor
//   - FastHTTP:     handler := middleware.FastHTTPLogger()(myHandler)
//   - fasthttp mux: h = middleware.FastHTTPLogger()(fasthttpMux.Handler)
//
// Each middleware automatically logs HTTP requests with:
//   - Request method and path
//   - Response status code
//   - Request duration
//   - Request ID (X-Request-ID header)
//   - Client IP
//
// Both variants share the same Config, SkipPaths, and log format.
//
// Example usage with standard net/http:
//
//	mux := http.NewServeMux()
//	mux.HandleFunc("/", handler)
//	loggedMux := middleware.Logger()(mux)
//	http.ListenAndServe(":8080", loggedMux)
//
// Example usage with fasthttp:
//
//	handler := middleware.FastHTTPLogger()(myHandler)
//	fasthttp.ListenAndServe(":8080", handler)
//
// Example usage with Chi:
//
//	import "github.com/otmc-sw/logger/middleware"
//
//	r := chi.NewRouter()
//	r.Use(middleware.Logger())
