/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
 **/
package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/otmc-sw/logger"
)

func main() {
	fmt.Println("=== Log Rotator Simulation ===")
	fmt.Println("This program simulates a running process that writes logs")
	fmt.Println("and demonstrates automatic log rotation.")
	fmt.Println()

	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create log directory: %v\n", err)
		os.Exit(1)
	}

	logFile := filepath.Join(logDir, "app.log")

	log := logger.New(
		logger.WithFile(logFile),
		logger.WithMaxSize(0.1),
		logger.WithMaxBackups(3),
		logger.WithMaxAge(7),
		logger.WithCompress(false),
		logger.WithConsole(false),
	)
	defer log.Sync()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Writing logs. Press Ctrl+C to stop.")
	fmt.Println()

	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	lineCount := 0
	startTime := time.Now()

loop:
	for {
		select {
		case <-ticker.C:
			elapsed := time.Since(startTime).Truncate(time.Second)
			lineCount++

			switch randomInt(0, 3) {
			case 0:
				log.Info("Request processed: method=%s path=/api/%s status=%d duration=%dms client=%s",
					randomAction(), randomString(8), randomStatus(), randomInt(10, 500), randomIP())
			case 1:
				log.Debug("Cache lookup: key=%s hit=%v ttl=%ds", randomString(12), randomBool(), randomInt(30, 3600))
			case 2:
				log.Warn("Slow query detected: query_id=%s duration=%dms threshold=%dms",
					randomString(8), randomInt(500, 3000), 500)
			case 3:
				log.Error("Failed to process request: request_id=%s error=%s retry=%d",
					randomString(8), randomError(), randomInt(0, 3))
			}

			if lineCount%500 == 0 {
				fmt.Printf("  Wrote %d lines (elapsed: %s)\n", lineCount, elapsed)
				printLogFiles(logDir)
			}

		case <-sigCh:
			fmt.Println()
			fmt.Println("Received interrupt signal. Shutting down...")
			break loop
		}
	}

	elapsed := time.Since(startTime).Truncate(time.Second)
	fmt.Println()
	fmt.Println("=== Summary ===")
	fmt.Printf("  Total lines written: %d\n", lineCount)
	fmt.Printf("  Elapsed time:        %s\n", elapsed)
	fmt.Println()
	fmt.Println("Log files in directory:")
	printLogFiles(logDir)
	fmt.Println()
	fmt.Println("Done.")
}

func printLogFiles(dir string) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read directory %s: %v\n", dir, err)
		return
	}

	if len(entries) == 0 {
		fmt.Println("  (no files)")
		return
	}

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		sizeMB := float64(info.Size()) / (1024 * 1024)
		fmt.Printf("    %-30s %8.3f MB\n", entry.Name(), sizeMB)
	}
}


var actions = []string{"GET", "POST", "PUT", "DELETE", "PATCH"}

func randomAction() string {
	return actions[randomInt(0, len(actions)-1)]
}

func randomStatus() int {
	weights := []int{200, 201, 204, 301, 302, 400, 401, 403, 404, 500, 502, 503}
	return weights[randomInt(0, len(weights)-1)]
}

var errors = []string{
	"connection refused",
	"timeout exceeded",
	"invalid input",
	"permission denied",
	"rate limit exceeded",
	"internal server error",
}

func randomError() string {
	return errors[randomInt(0, len(errors)-1)]
}

var seed = time.Now().UnixNano()

func randomInt(min, max int) int {
	seed = seed*6364136223846793005 + 1442695040888963407
	if seed < 0 {
		seed = -seed
	}
	return min + int(seed%int64(max-min+1))
}

func randomBool() bool {
	return randomInt(0, 1) == 1
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[randomInt(0, len(charset)-1)]
	}
	return string(b)
}

func randomIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", randomInt(1, 255), randomInt(0, 255), randomInt(0, 255), randomInt(1, 255))
}
