/*
 * @Author: facsert
 * @Date: 2023-08-02 20:03:24
 * @LastEditTime: 2023-08-16 22:59:19
 * @LastEditors: facsert
 * @Description: logger package record log
 */

package logger

import (
	"fmt"
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)


var (
	info    *log.Logger
	erro    *log.Logger
	warning *log.Logger
	fp      *os.File
	mu      sync.Mutex
)

var logFile = "process.log"

func init() {
	logFile = absPath(logFile)
	os.WriteFile(logFile, []byte(""), 0666)
	SetLogOutput(logFile)
}

func SetLogOutput(logPath string) {
	Close()
	
	logFile = absPath(logPath)
	fp, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Sprintf("open %s failed, err: %v\n", err, logFile))
	}
    
	mu.Lock()
	defer mu.Unlock()

	out := io.MultiWriter(fp, os.Stdout)
	info = log.New(out, "[ INFO  ]", log.Ldate|log.Ltime)
	erro = log.New(out, "[ ERROR ]", log.Ldate|log.Ltime)
	warning = log.New(out, "[WARNING]", log.Ldate|log.Ltime)
}

func absPath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	rootPath, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("get root path failed: %v\n", err))
	}
	return filepath.Join(rootPath, path)
}

func GetOutput() {
	info.Printf("logFile: %s\n", logFile)
}

func Info(format string, a ...any) {
	info.Printf(format, a...)
}

func Error(format string, a ...any) {
	erro.Printf(format, a...)
}

func Warring(format string, a ...any) {
	warning.Printf(format, a...)
}

func Display(isPass bool, format string, a ...any) {
	s := fmt.Sprintf(format, a...)
	red, green, reset := "\033[1;31m", "\033[1;32m", "\033[0m"
	if isPass {
		info.Printf("%s%-80s [PASS]%s\n", green, s, reset)
	} else {
		erro.Printf("%s%-80s [FAIL]%s\n", red, s, reset)
	}
}

func Close() {
	mu.Lock()
	defer mu.Unlock()

	if fp != nil {
		bufWriter := bufio.NewWriter(fp)
		bufWriter.Flush()
		fp.Close()
	}
}
