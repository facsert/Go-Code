/*
 * @Author: facsert
 * @Date: 2023-08-02 20:03:24
 * @LastEditTime: 2023-08-16 22:59:19
 * @LastEditors: facsert
 * @Description: logger package record log
 */

package logger

import (
	"io"
	"os"
	"log/slog"

	"learn/utils/comm"
)

var (
	logFile = comm.AbsPath("report.log")
	// logLevel = slog.LevelDebug
	// codeSource = false
)

func Init() {
    file, err := os.OpenFile(logFile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil { panic(err) }
    
	// logger := slog.New(slog.NewJSONHandler(io.MultiWriter(os.Stdout, file), &slog.HandlerOptions{
	// 	Level: logLevel,
	// 	AddSource: codeSource,
	// }))

	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(os.Stdout, file), nil))

    slog.SetDefault(logger)
}

