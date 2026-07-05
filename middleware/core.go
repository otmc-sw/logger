/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
**/

package middleware

import (
	"net/http"
	"time"

	"github.com/otmc-sw/logger"
	"github.com/valyala/fasthttp"
)

// responseWriter wraps http.ResponseWriter to capture the status code.
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// StatusCode returns the captured HTTP status code.
func (rw *responseWriter) StatusCode() int {
	return rw.statusCode
}

// Config holds optional configuration for the logging middleware.
type Config struct {
	// SkipPaths is a list of paths to skip logging for.
	SkipPaths []string
}

// =============================================================================
// net/http (standard library)
// =============================================================================

// Logger returns a standard net/http middleware handler that logs HTTP requests.
// Works with any framework that supports http.Handler:
//   - Chi:        r.Use(middleware.Logger())
//   - Gin:        router.Use(gin.WrapH(middleware.Logger()(nextHandler)))
//   - Echo:       e.Use(echo.WrapMiddleware(middleware.Logger()))
//   - std mux:    handler := middleware.Logger()(mux)
func Logger() func(http.Handler) http.Handler {
	return LoggerWithConfig(Config{})
}

// LoggerWithConfig returns a net/http middleware handler with configuration.
func LoggerWithConfig(config Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			path := r.URL.Path
			raw := r.URL.RawQuery
			requestID := r.Header.Get("X-Request-ID")

			for _, skip := range config.SkipPaths {
				if path == skip {
					next.ServeHTTP(w, r)
					return
				}
			}

			wrapped := newResponseWriter(w)
			next.ServeHTTP(wrapped, r)

			latency := time.Since(start)
			clientIP := r.RemoteAddr
			method := r.Method
			statusCode := wrapped.StatusCode()

			if raw != "" {
				path = path + "?" + raw
			}

			writeLog(requestID, method, path, statusCode, latency, clientIP)
		})
	}
}

// =============================================================================
// fasthttp (valyala/fasthttp)
// =============================================================================

// FastHTTPLogger returns a fasthttp middleware handler that logs HTTP requests.
// Works with any framework built on fasthttp:
//   - Fiber:      use fasthttp.adaptor or this handler directly
//   - FastHTTP:   h = middleware.FastHTTPLogger()(h)
//
// Usage with raw fasthttp server:
//
//	handler := middleware.FastHTTPLogger()(myHandler)
//	fasthttp.ListenAndServe(":8080", handler)
func FastHTTPLogger() func(fasthttp.RequestHandler) fasthttp.RequestHandler {
	return FastHTTPLoggerWithConfig(Config{})
}

// FastHTTPLoggerWithConfig returns a fasthttp middleware handler with configuration.
func FastHTTPLoggerWithConfig(config Config) func(fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			start := time.Now()
			path := string(ctx.Path())
			raw := string(ctx.QueryArgs().QueryString())
			requestID := string(ctx.Request.Header.Peek("X-Request-ID"))

			for _, skip := range config.SkipPaths {
				if path == skip {
					next(ctx)
					return
				}
			}

			// Process request
			next(ctx)

			// After request
			latency := time.Since(start)
			clientIP := ctx.RemoteIP().String()
			method := string(ctx.Method())
			statusCode := ctx.Response.StatusCode()

			if raw != "" {
				path = path + "?" + raw
			}

			writeLog(requestID, method, path, statusCode, latency, clientIP)
		}
	}
}

// =============================================================================
// Shared helpers
// =============================================================================

// writeLog writes a structured HTTP log entry using the logger package.
func writeLog(requestID, method, path string, statusCode int, latency time.Duration, clientIP string) {
	logMessage := method + " " + path
	switch {
	case statusCode >= 500:
		logger.Error("[%s] %s | %d | %v | %s",
			requestID, logMessage, statusCode, latency, clientIP)
	case statusCode >= 400:
		logger.Warn("[%s] %s | %d | %v | %s",
			requestID, logMessage, statusCode, latency, clientIP)
	default:
		logger.Info("[%s] %s | %d | %v | %s",
			requestID, logMessage, statusCode, latency, clientIP)
	}
}
