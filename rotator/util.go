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

func parseFilename(filename string) (base, ext string) {
	filename = filepath.Base(filename)
	ext = filepath.Ext(filename)
	base = strings.TrimSuffix(filename, ext)
	if base == "" {
		base = "log"
	}
	return
}

func megabytesToBytes(mb int) int64 {
	return int64(mb) * 1024 * 1024
}

func extractIndex(filename, baseName, ext string) int {
	if len(filename) <= len(baseName) {
		return 0
	}

	suffix := filename[len(baseName):]

	var idx int
	_, err := fmt.Sscanf(suffix, "_%d"+ext, &idx)
	if err == nil {
		return idx
	}

	_, err = fmt.Sscanf(suffix, "_%d"+ext+".gz", &idx)
	if err == nil {
		return idx
	}

	return 0
}
