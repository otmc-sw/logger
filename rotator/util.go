/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
 **/
package rotator

import (
	"fmt"
	"path/filepath"
	"strings"
)

// parseFilename splits a filename into base name and extension.
func parseFilename(filename string) (base, ext string) {
	// Get just the filename without directory path
	filename = filepath.Base(filename)
	ext = filepath.Ext(filename)
	base = strings.TrimSuffix(filename, ext)
	if base == "" {
		base = "log"
	}
	return
}

// megabytesToBytes converts megabytes to bytes.
func megabytesToBytes(mb int) int64 {
	return int64(mb) * 1024 * 1024
}

// extractIndex attempts to extract the index from a backup filename.
func extractIndex(filename, baseName, ext string) int {
	// Remove base name
	if len(filename) <= len(baseName) {
		return 0
	}

	suffix := filename[len(baseName):]

	// Simple case: _123.ext
	var idx int
	_, err := fmt.Sscanf(suffix, "_%d"+ext, &idx)
	if err == nil {
		return idx
	}

	// Try with .gz extension
	_, err = fmt.Sscanf(suffix, "_%d"+ext+".gz", &idx)
	if err == nil {
		return idx
	}

	return 0
}
