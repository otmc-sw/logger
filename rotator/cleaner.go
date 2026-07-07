/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
 **/
package rotator

import (
	"os"
	"path/filepath"
	"sort"
	"time"
)

func cleanup(cfg Config) error {
	if cfg.MaxBackups <= 0 && cfg.MaxAge <= 0 {
		return nil
	}

	dir := filepath.Dir(cfg.Filename)
	if dir == "" {
		dir = "."
	}

	baseName, ext := parseFilename(cfg.Filename)

	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	var backups []backupInfo

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		info, err := file.Info()
		if err != nil {
			continue
		}

		if !isBackupFile(name, baseName, ext) {
			continue
		}

		backups = append(backups, backupInfo{
			name:     name,
			path:     filepath.Join(dir, name),
			modTime:  info.ModTime(),
			size:     info.Size(),
		})
	}

	sort.Slice(backups, func(i, j int) bool {
		return backups[i].modTime.After(backups[j].modTime)
	})

	if cfg.MaxBackups > 0 && len(backups) > cfg.MaxBackups {
		for i := cfg.MaxBackups; i < len(backups); i++ {
			_ = os.Remove(backups[i].path)
		}
		backups = backups[:cfg.MaxBackups]
	}

	if cfg.MaxAge > 0 {
		cutoff := time.Now().AddDate(0, 0, -cfg.MaxAge)
		for _, backup := range backups {
			if backup.modTime.Before(cutoff) {
				_ = os.Remove(backup.path)
			}
		}
	}

	return nil
}

type backupInfo struct {
	name    string
	path    string
	modTime time.Time
	size    int64
}

func isBackupFile(filename, baseName, ext string) bool {
	if filename == baseName+ext {
		return false // This is the active file
	}
	
	if len(filename) <= len(baseName) {
		return false
	}
	
	if filename[:len(baseName)] != baseName {
		return false
	}
	
	if ext == "" {
		return true
	}
	
	return true
}
