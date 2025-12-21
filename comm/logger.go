package comm

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
	"log/slog"
)

type Logger struct {
	FilePath string
	MaxSize  int

	maxSize  int64
	currSize int64
	file     *os.File
	mu       sync.Mutex
}

func exists(dir string) bool {
	if _, err := os.Stat(dir); err == nil || os.IsExist(err) {
		return true
	}
	return false
}
func NewLogger(filePath string, maxSize int) error {
	if !exists(filepath.Dir(filePath)) {
		if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
			return fmt.Errorf("create directory: %w", err)
		}
	}

	l := &Logger{
		FilePath: filePath,
		MaxSize:  maxSize,

		maxSize:  int64(maxSize * 10),
		currSize: 0,
		file:     nil,
	}

	if err := l.openFile(filePath); err != nil {
		return err
	}
    
	slog.SetDefault(
		slog.New(slog.NewTextHandler(
			io.MultiWriter(l, os.Stdout), 
			&slog.HandlerOptions{ Level: slog.LevelInfo},
		)),
    )

	return nil
}

func (l *Logger) openFile(FilePath string) error {
	file, err := os.OpenFile(FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
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
	l.mu.Lock()
	defer l.mu.Unlock()

	writeLen := int64(len(p))
	if writeLen > l.maxSize {
		return 0, fmt.Errorf("write length bigger than max size")
	}

	if l.currSize+int64(n) > l.maxSize {
		if err := l.rotate(); err != nil {
			return 0, err
		}
	}

	l.currSize += int64(n)
	n, err = l.file.Write(p)
	return n, err
}

func (l *Logger) rotate() error {
	if l.file != nil {
		if err := l.file.Close(); err != nil {
			return fmt.Errorf("close file before rotate: %w", err)
		}
	}

	stat, err := os.Stat(l.FilePath)
	if os.IsNotExist(err) || stat.Size() == 0 {
		return l.openFile(l.FilePath)
	}

	dir := filepath.Dir(l.FilePath)
	name := filepath.Base(l.FilePath)
	ext := filepath.Ext(name)
	prefix := name[:len(name)-len(ext)]
	newName := filepath.Join(dir, fmt.Sprintf("%s-%s%s", prefix, time.Now().Format("20060102_150405"), ext))

	if err := os.Rename(l.FilePath, newName); err != nil {
		return fmt.Errorf("rename file: %w", err)
	}

	return l.openFile(l.FilePath)
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
