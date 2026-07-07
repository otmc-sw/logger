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

// Rotator implements io.Writer and io.Closer for log rotation.
type Rotator struct {
	mu     sync.Mutex
	config Config
	file   *os.File
	size   int64
	closed bool
}

// New creates a new Rotator with the given options.
func New(opts ...Option) *Rotator {
	cfg := defaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	r := &Rotator{
		config: cfg,
	}

	// Open the initial file
	if err := r.openFile(); err != nil {
		// If we can't open the file, return a rotator that will error on write
		// This allows the caller to handle the error
	}

	return r
}

// Write implements io.Writer.
func (r *Rotator) Write(p []byte) (n int, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.closed {
		return 0, ErrClosed
	}

	// Check if we need to rotate
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

// Close implements io.Closer.
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

// Rotate forces a rotation.
func (r *Rotator) Rotate() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.closed {
		return ErrClosed
	}

	return r.rotateLocked()
}

// Sync flushes the file to disk.
func (r *Rotator) Sync() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.closed || r.file == nil {
		return ErrClosed
	}

	return r.file.Sync()
}

// Size returns the current size of the log file.
func (r *Rotator) Size() int64 {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.size
}

// openFile opens the log file for writing.
func (r *Rotator) openFile() error {
	if r.config.Filename == "" {
		return ErrInvalidFilename
	}

	file, err := os.OpenFile(r.config.Filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	// Get current file size
	info, err := file.Stat()
	if err != nil {
		file.Close()
		return err
	}

	r.file = file
	r.size = info.Size()
	return nil
}

// needRotation checks if rotation is needed based on the incoming data size.
func (r *Rotator) needRotation(incomingSize int) bool {
	if r.config.MaxSize <= 0 {
		return false
	}

	maxBytes := megabytesToBytes(r.config.MaxSize)
	return r.size+int64(incomingSize) > maxBytes
}

// rotateLocked performs rotation with the mutex already held.
func (r *Rotator) rotateLocked() error {
	// Close current file
	if r.file != nil {
		if err := r.file.Close(); err != nil {
			return err
		}
	}

	// Generate backup filename
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

	// Rename current file to backup
	if err := os.Rename(r.config.Filename, backupPath); err != nil {
		// If rename fails (file might not exist), try to open new file anyway
		if !os.IsNotExist(err) {
			return err
		}
	}

	// Compress if enabled
	if r.config.Compress {
		compressPath := backupPath + ".gz"
		if err := compressFile(backupPath, compressPath); err == nil {
			_ = os.Remove(backupPath)
		}
	}

	// Cleanup old backups
	_ = cleanup(r.config)

	// Open new file
	return r.openFile()
}

// findNextIndexLocked finds the next available index for backup files (mutex must be held).
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
