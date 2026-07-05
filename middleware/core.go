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

func (rw *responseWriter) StatusCode() int {
	return rw.statusCode
}

type Config struct {
	SkipPaths []string
}

func Logger() func(http.Handler) http.Handler {
	return LoggerWithConfig(Config{})
}

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

func FastHTTPLogger() func(fasthttp.RequestHandler) fasthttp.RequestHandler {
	return FastHTTPLoggerWithConfig(Config{})
}

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

			next(ctx)

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

func writeLog(requestID, method, path string, statusCode int, latency time.Duration, clientIP string) {
	logger.Request(method, path, statusCode, latency, clientIP)
}
