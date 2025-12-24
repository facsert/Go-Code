package comm

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"time"
)

var DefaultLog *Logger

type Logger struct {
	FilePath  string
	MaxSizeMB int // MB

	maxSizeByte int64 // byte
	maxBackup   int
	currSize    int64
	file        *os.File
	mu          sync.Mutex
}

func exists(dir string) bool {
	_, err := os.Stat(dir)
	return err == nil
}

func NewLogger(filePath string, maxSize, maxBackup int) error {
	if !exists(filepath.Dir(filePath)) {
		if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
			return fmt.Errorf("create directory: %w", err)
		}
	}

	l := &Logger{
		FilePath:  filePath,
		MaxSizeMB: maxSize,

		maxSizeByte: int64(maxSize * 1024 * 1025),
		maxBackup:   maxBackup,
		currSize:    0,
		file:        nil,
	}

	if err := l.openFile(filePath); err != nil {
		return err
	}

	DefaultLog = l
	slog.SetDefault(
		slog.New(slog.NewTextHandler(
			io.MultiWriter(l, os.Stdout),
			&slog.HandlerOptions{Level: slog.LevelInfo},
		)),
	)
	return nil
}

func (l *Logger) openFile(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}

	stat, err := file.Stat()
	if err != nil {
		file.Close()
		return fmt.Errorf("get stat: %w", err)
	}

	l.file = file
	l.currSize = stat.Size()
	return nil
}

func (l *Logger) Write(p []byte) (n int, err error) {
	writeLen := int64(len(p))
	if writeLen > l.maxSizeByte {
		return 0, fmt.Errorf("write length bigger than max size")
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	if l.currSize+writeLen > l.maxSizeByte {
		if err := l.rotate(); err != nil {
			return 0, err
		}
	}

	n, err = l.file.Write(p)
	if err == nil {
		l.currSize += int64(n)
	}
	return n, err
}

func (l *Logger) rotate() error {
	if l.file != nil {
		l.file.Close()
		l.file = nil
	}

	dir := filepath.Dir(l.FilePath)
	name := filepath.Base(l.FilePath)
	ext := filepath.Ext(name)
	prefix := name[:len(name)-len(ext)]
	newName := filepath.Join(dir, fmt.Sprintf("%s-%s%s", prefix, time.Now().Format("20060102_150405"), ext))

	if err := os.Rename(l.FilePath, newName); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("rename file: %w", err)
	}

	l.clean(dir, prefix)
	l.currSize = 0
	return l.openFile(l.FilePath)
}

func (l *Logger) clean(dir, prefix string) {
	if l.maxBackup <= 0 {
		return
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return
	}

	type LogFile struct {
		Name    string
		ModTime time.Time
	}

	logFiles := make([]LogFile, 0, len(files))
	for _, file := range files {
		name := file.Name()
		if file.IsDir() {
			continue
		}

		if !strings.HasPrefix(name, prefix) {
			continue
		}

		info, err := file.Info()
		if err != nil {
			continue
		}

		logFiles = append(logFiles, LogFile{name, info.ModTime()})
	}

	removeNum := len(logFiles) - l.maxBackup
	if removeNum <= 0 {
		return
	}

	slices.SortFunc(logFiles, func(a, b LogFile) int {
		if a.ModTime.After(b.ModTime) {
			return 1
		}
		return -1
	})

	for i := range removeNum {
		os.Remove(filepath.Join(dir, logFiles[i].Name))
	}
}

func (l *Logger) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.file != nil {
		err := l.file.Close()
		l.file = nil
		return err
	}
	return nil
}
