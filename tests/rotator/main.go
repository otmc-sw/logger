/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
 **/
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/otmc-sw/logger/rotator"
)

func main() {
	fmt.Println("=== Log Rotator Simulation ===")
	fmt.Println("This program simulates a running process that writes logs")
	fmt.Println("and demonstrates automatic log rotation.")
	fmt.Println()

	// Create a dedicated log directory
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}

	logFile := filepath.Join(logDir, "app.log")

	// Configure the rotator with a small MaxSize (0.24 MB = ~240 KB) to trigger rotation quickly
	// MaxBackups=3 keeps at most 3 rotated files, MaxAge=7 days, Compress=false
	r := rotator.New(
		rotator.WithFilename(logFile),
		rotator.WithMaxSize(0.24), // 0.24 MB (~240 KB) – rotate when file exceeds this
		rotator.WithMaxBackups(3), // keep up to 3 backup files
		rotator.WithMaxAge(7),     // keep backups for 7 days
		rotator.WithCompress(false),
	)
	defer r.Close()

	// Channel to listen for OS interrupt signals (Ctrl+C)
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
			// Simulate a log entry with timestamp, level, and message
			elapsed := time.Since(startTime).Truncate(time.Second)
			lineCount++

			// Generate a log line of roughly 200 bytes
			logLine := fmt.Sprintf(
				"[%s] [%s] [%d] request_id=%s user=%s action=%s status=%d duration=%dms payload=%s\n",
				time.Now().Format(time.RFC3339),
				randomLevel(),
				lineCount,
				randomString(8),
				randomString(6),
				randomAction(),
				randomStatus(),
				randomInt(10, 500),
				randomString(20),
			)

			_, err := r.Write([]byte(logLine))
			if err != nil {
				log.Printf("Write error: %v", err)
			}

			// Print progress every 500 lines
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

// printLogFiles lists all files in the given directory with their sizes.
func printLogFiles(dir string) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Printf("Failed to read directory %s: %v", dir, err)
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

// --- Helpers to simulate realistic log data ---

var levels = []string{"INFO", "WARN", "ERROR", "DEBUG"}

func randomLevel() string {
	return levels[randomInt(0, len(levels)-1)]
}

var actions = []string{"GET", "POST", "PUT", "DELETE", "PATCH"}

func randomAction() string {
	return actions[randomInt(0, len(actions)-1)]
}

func randomStatus() int {
	// Weighted toward 2xx/3xx, occasional 4xx/5xx
	weights := []int{200, 201, 204, 301, 302, 400, 401, 403, 404, 500, 502, 503}
	return weights[randomInt(0, len(weights)-1)]
}

// Simple pseudo-random number generator (no external dependency)
var seed = time.Now().UnixNano()

func randomInt(min, max int) int {
	seed = seed*6364136223846793005 + 1442695040888963407
	if seed < 0 {
		seed = -seed
	}
	return min + int(seed%int64(max-min+1))
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[randomInt(0, len(charset)-1)]
	}
	return string(b)
}