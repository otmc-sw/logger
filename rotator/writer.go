/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
 **/
package rotator

import (
	"os"
	"path/filepath"
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

	baseName, ext := parseFilename(r.config.Filename)
	index := r.findNextIndexLocked()

	info := RotateInfo{
		BaseName:  baseName,
		Extension: ext,
		Index:     index,
		Time:      time.Now(),
	}

	backupName := r.config.Naming(info)
	backupPath := filepath.Join(filepath.Dir(r.config.Filename), backupName)

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

func (r *Rotator) findNextIndexLocked() int {
	dir := filepath.Dir(r.config.Filename)
	if dir == "" {
		dir = "."
	}

	baseName, ext := parseFilename(r.config.Filename)

	files, err := os.ReadDir(dir)
	if err != nil {
		return 1
	}

	maxIndex := 0
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		if idx := extractIndex(name, baseName, ext); idx > maxIndex {
			maxIndex = idx
		}
	}

	return maxIndex + 1
}
