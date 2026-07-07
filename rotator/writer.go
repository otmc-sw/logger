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
	"sync"
	"time"
)

type Rotator struct {
	mu     sync.Mutex
	config Config
	file   *os.File
	size   int64
	closed bool
}

func New(opts ...Option) *Rotator {
	cfg := defaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	r := &Rotator{
		config: cfg,
	}

	if err := r.openFile(); err != nil {
	}

	return r
}

func (r *Rotator) Write(p []byte) (n int, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.closed {
		return 0, ErrClosed
	}

	if r.needRotation(len(p)) {
		if err := r.rotateLocked(); err != nil {
			return 0, err
		}
	}

	n, err = r.file.Write(p)
	if err != nil {
		return n, err
	}

	r.size += int64(n)
	return n, nil
}

func (r *Rotator) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.closed {
		return nil
	}

	r.closed = true
	if r.file != nil {
		return r.file.Close()
	}
	return nil
}

func (r *Rotator) Rotate() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.closed {
		return ErrClosed
	}

	return r.rotateLocked()
}

func (r *Rotator) Sync() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.closed || r.file == nil {
		return ErrClosed
	}

	return r.file.Sync()
}

func (r *Rotator) Size() int64 {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.size
}

func (r *Rotator) openFile() error {
	if r.config.Filename == "" {
		return ErrInvalidFilename
	}

	file, err := os.OpenFile(r.config.Filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	info, err := file.Stat()
	if err != nil {
		file.Close()
		return err
	}

	r.file = file
	r.size = info.Size()
	return nil
}

func (r *Rotator) needRotation(incomingSize int) bool {
	if r.config.MaxSize <= 0 {
		return false
	}

	maxBytes := megabytesToBytes(r.config.MaxSize)
	return r.size+int64(incomingSize) > maxBytes
}

func (r *Rotator) rotateLocked() error {
	if r.file != nil {
		if err := r.file.Close(); err != nil {
			return err
		}
	}

	dir := filepath.Dir(r.config.Filename)
	baseName, ext := parseFilename(r.config.Filename)

	r.shiftBackupsLocked(dir, baseName, ext)

	info := RotateInfo{
		BaseName:  baseName,
		Extension: ext,
		Index:     1,
		Time:      time.Now(),
	}
	backupName := r.config.Naming(info)
	backupPath := filepath.Join(dir, backupName)

	if err := os.Rename(r.config.Filename, backupPath); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	if r.config.Compress {
		compressPath := backupPath + ".gz"
		if err := compressFile(backupPath, compressPath); err == nil {
			_ = os.Remove(backupPath)
		}
	}

	_ = cleanup(r.config)

	return r.openFile()
}

func (r *Rotator) shiftBackupsLocked(dir, baseName, ext string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return
	}

	type backupEntry struct {
		oldPath string
		newPath string
		index   int
	}

	var backups []backupEntry

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		name := file.Name()
		if !isBackupFile(name, baseName, ext) {
			continue
		}

		idx := extractIndex(name, baseName, ext)
		if idx == 0 {
			continue
		}

		info := RotateInfo{
			BaseName:  baseName,
			Extension: ext,
			Index:     idx + 1,
			Time:      time.Now(),
		}
		newName := r.config.Naming(info)
		backups = append(backups, backupEntry{
			oldPath: filepath.Join(dir, name),
			newPath: filepath.Join(dir, newName),
			index:   idx,
		})
	}

	sort.Slice(backups, func(i, j int) bool {
		return backups[i].index > backups[j].index
	})

	for _, b := range backups {
		_ = os.Rename(b.oldPath, b.newPath)
	}
}
