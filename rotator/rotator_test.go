/**
 * @License Apache License 2.0
 * @Copyright (c) 2026 OTMC Softwares. OTMC Golang Logger.
 * @Contributors Nguyen Van Trung, Nguyen Thi Hoai, OTMC Contributors.
 **/
package rotator

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "test.log")

	r := New(
		WithFilename(filename),
		WithMaxSize(1),
		WithMaxBackups(3),
		WithMaxAge(7),
		WithCompress(false),
	)

	if r == nil {
		t.Fatal("New() returned nil")
	}

	defer r.Close()

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Error("Log file was not created")
	}
}

func TestWrite(t *testing.T) {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "test.log")

	r := New(WithFilename(filename))
	defer r.Close()

	data := []byte("test log entry\n")
	n, err := r.Write(data)
	if err != nil {
		t.Fatalf("Write() failed: %v", err)
	}

	if n != len(data) {
		t.Errorf("Write() wrote %d bytes, expected %d", n, len(data))
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	if !bytes.Equal(content, data) {
		t.Errorf("File content mismatch: got %q, want %q", content, data)
	}
}

func TestRotation(t *testing.T) {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "test.log")

	r := New(
		WithFilename(filename),
		WithMaxSize(1),
		WithMaxBackups(2),
		WithCompress(false),
	)
	defer r.Close()

	data := make([]byte, 1024*1024+1)
	for i := range data {
		data[i] = 'a'
	}

	_, err := r.Write(data)
	if err != nil {
		t.Fatalf("Write() failed: %v", err)
	}

	time.Sleep(100 * time.Millisecond)

	files, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to read directory: %v", err)
	}

	if len(files) < 2 {
		t.Errorf("Expected at least 2 files after rotation, got %d", len(files))
	}
}

func TestCompression(t *testing.T) {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "test.log")

	r := New(
		WithFilename(filename),
		WithMaxSize(1),
		WithMaxBackups(1),
		WithCompress(true),
	)
	defer r.Close()

	data := make([]byte, 1024*1024+1)
	for i := range data {
		data[i] = 'a'
	}

	_, err := r.Write(data)
	if err != nil {
		t.Fatalf("Write() failed: %v", err)
	}

	time.Sleep(100 * time.Millisecond)

	files, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to read directory: %v", err)
	}

	foundGz := false
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".gz" {
			foundGz = true
			break
		}
	}

	if !foundGz {
		t.Error("Expected compressed backup file (.gz) not found")
	}
}

func TestMaxBackups(t *testing.T) {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "test.log")

	r := New(
		WithFilename(filename),
		WithMaxSize(1),
		WithMaxBackups(2),
		WithCompress(false),
	)
	defer r.Close()

	for i := 0; i < 4; i++ {
		data := make([]byte, 1024*1024+1)
		for j := range data {
			data[j] = byte('a' + i)
		}

		_, err := r.Write(data)
		if err != nil {
			t.Fatalf("Write() failed on iteration %d: %v", i, err)
		}

		time.Sleep(50 * time.Millisecond)
	}

	time.Sleep(100 * time.Millisecond)

	files, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to read directory: %v", err)
	}

	if len(files) > 3 {
		t.Errorf("Expected at most 3 files with MaxBackups=2, got %d", len(files))
	}
}

func TestClose(t *testing.T) {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "test.log")

	r := New(WithFilename(filename))

	err := r.Close()
	if err != nil {
		t.Fatalf("Close() failed: %v", err)
	}

	_, err = r.Write([]byte("test"))
	if err != ErrClosed {
		t.Errorf("Expected ErrClosed after Close(), got %v", err)
	}
}

func TestSync(t *testing.T) {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "test.log")

	r := New(WithFilename(filename))
	defer r.Close()

	_, err := r.Write([]byte("test"))
	if err != nil {
		t.Fatalf("Write() failed: %v", err)
	}

	err = r.Sync()
	if err != nil {
		t.Fatalf("Sync() failed: %v", err)
	}
}

func TestSize(t *testing.T) {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "test.log")

	r := New(WithFilename(filename))
	defer r.Close()

	data := []byte("test log entry\n")
	_, err := r.Write(data)
	if err != nil {
		t.Fatalf("Write() failed: %v", err)
	}

	size := r.Size()
	if size != int64(len(data)) {
		t.Errorf("Size() returned %d, expected %d", size, len(data))
	}
}

func TestRotate(t *testing.T) {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "test.log")

	r := New(
		WithFilename(filename),
		WithMaxBackups(1),
		WithCompress(false),
	)
	defer r.Close()

	_, err := r.Write([]byte("test data\n"))
	if err != nil {
		t.Fatalf("Write() failed: %v", err)
	}

	err = r.Rotate()
	if err != nil {
		t.Fatalf("Rotate() failed: %v", err)
	}

	time.Sleep(50 * time.Millisecond)

	files, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to read directory: %v", err)
	}

	if len(files) < 2 {
		t.Errorf("Expected at least 2 files after manual rotation, got %d", len(files))
	}
}

func TestNamingWithIndex(t *testing.T) {
	info := RotateInfo{
		BaseName:  "app",
		Extension: ".log",
		Index:     1,
		Time:      time.Now(),
	}

	name := NameWithIndex(info)
	expected := "app_1.log"

	if name != expected {
		t.Errorf("NameWithIndex() = %q, want %q", name, expected)
	}
}

func TestNamingWithDateAndIndex(t *testing.T) {
	now := time.Date(2025, 10, 21, 0, 0, 0, 0, time.UTC)
	info := RotateInfo{
		BaseName:  "app",
		Extension: ".log",
		Index:     1,
		Time:      now,
	}

	name := NameWithDateAndIndex(info)
	expected := "app_20251021_1.log"

	if name != expected {
		t.Errorf("NameWithDateAndIndex() = %q, want %q", name, expected)
	}
}

func TestNamingWithTimestamp(t *testing.T) {
	now := time.Date(2025, 10, 21, 15, 30, 15, 0, time.UTC)
	info := RotateInfo{
		BaseName:  "app",
		Extension: ".log",
		Index:     0,
		Time:      now,
	}

	name := NameWithTimestamp(info)
	expected := "app_20251021_153015.log"

	if name != expected {
		t.Errorf("NameWithTimestamp() = %q, want %q", name, expected)
	}
}

func TestNamingWithTimestampAndIndex(t *testing.T) {
	now := time.Date(2025, 10, 21, 15, 30, 15, 0, time.UTC)
	info := RotateInfo{
		BaseName:  "app",
		Extension: ".log",
		Index:     1,
		Time:      now,
	}

	name := NameWithTimestampAndIndex(info)
	expected := "app_20251021_153015_1.log"

	if name != expected {
		t.Errorf("NameWithTimestampAndIndex() = %q, want %q", name, expected)
	}
}

func TestCustomNaming(t *testing.T) {
	customFunc := func(info RotateInfo) string {
		return "custom_" + info.BaseName + info.Extension
	}

	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "test.log")

	r := New(
		WithFilename(filename),
		WithMaxSize(1),
		WithMaxBackups(1),
		WithNaming(customFunc),
		WithCompress(false),
	)
	defer r.Close()

	data := make([]byte, 1024*1024+1)
	for i := range data {
		data[i] = 'a'
	}

	_, err := r.Write(data)
	if err != nil {
		t.Fatalf("Write() failed: %v", err)
	}

	time.Sleep(100 * time.Millisecond)

	files, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to read directory: %v", err)
	}

	foundCustom := false
	for _, file := range files {
		if file.Name() == "custom_test.log" {
			foundCustom = true
			break
		}
	}

	if !foundCustom {
		t.Error("Custom naming was not applied")
	}
}

func TestParseFilename(t *testing.T) {
	tests := []struct {
		input string
		base  string
		ext   string
	}{
		{"app.log", "app", ".log"},
		{"app.txt", "app", ".txt"},
		{"app", "app", ""},                 // no extension
		{".log", "log", ".log"},            // hidden file with extension
		{"path/to/app.log", "app", ".log"}, // path is stripped
	}

	for _, tt := range tests {
		base, ext := parseFilename(tt.input)
		if base != tt.base {
			t.Errorf("parseFilename(%q) base = %q, want %q", tt.input, base, tt.base)
		}
		if ext != tt.ext {
			t.Errorf("parseFilename(%q) ext = %q, want %q", tt.input, ext, tt.ext)
		}
	}
}

func TestMegabytesToBytes(t *testing.T) {
	tests := []struct {
		mb    int
		bytes int64
	}{
		{1, 1024 * 1024},
		{10, 10 * 1024 * 1024},
		{100, 100 * 1024 * 1024},
	}

	for _, tt := range tests {
		result := megabytesToBytes(tt.mb)
		if result != tt.bytes {
			t.Errorf("megabytesToBytes(%d) = %d, want %d", tt.mb, result, tt.bytes)
		}
	}
}

func TestExtractIndex(t *testing.T) {
	tests := []struct {
		filename string
		baseName string
		ext      string
		expected int
	}{
		{"app_1.log", "app", ".log", 1},
		{"app_2.log", "app", ".log", 2},
		{"app_10.log", "app", ".log", 10},
		{"app_1.log.gz", "app", ".log", 1},
		{"app.log", "app", ".log", 0},
		{"other_1.log", "app", ".log", 0},
	}

	for _, tt := range tests {
		result := extractIndex(tt.filename, tt.baseName, tt.ext)
		if result != tt.expected {
			t.Errorf("extractIndex(%q, %q, %q) = %d, want %d",
				tt.filename, tt.baseName, tt.ext, result, tt.expected)
		}
	}
}

func BenchmarkWrite(b *testing.B) {
	tmpDir := b.TempDir()
	filename := filepath.Join(tmpDir, "bench.log")

	r := New(
		WithFilename(filename),
		WithMaxSize(1000),
	)
	defer r.Close()

	data := []byte("benchmark log entry\n")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Write(data)
	}
}

func BenchmarkWriteParallel(b *testing.B) {
	tmpDir := b.TempDir()
	filename := filepath.Join(tmpDir, "bench.log")

	r := New(
		WithFilename(filename),
		WithMaxSize(1000),
	)
	defer r.Close()

	data := []byte("benchmark log entry\n")

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r.Write(data)
		}
	})
}
