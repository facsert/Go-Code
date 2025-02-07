package comm

import (
	"io"
	"log/slog"
	"os"

	"path/filepath"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logFile  = AbsPath("log", "report.log")
	logLevel = slog.LevelDebug
)

func LoggerInit() {
	MakeDirs(filepath.Dir(logFile))

	log := &lumberjack.Logger{
		Filename:   logFile, // 日志文件的位置
		MaxSize:    1,       // 文件最大尺寸（以MB为单位）
		MaxBackups: 3,       // 保留的最大旧文件数量
		MaxAge:     28,      // 保留旧文件的最大天数
		Compress:   false,   // 是否压缩/归档旧文件
		LocalTime:  true,    // 使用本地时间创建时间戳
	}

	handleOptions := slog.HandlerOptions{
		Level: logLevel, // 设置打印等级(slog.LevelInfo)
	}
	slog.SetDefault(slog.New(slog.NewTextHandler(io.MultiWriter(log, os.Stdout), &handleOptions)))
}
