package middleware

import (
	"io"
	"os"
	"time"
    
	"gopkg.in/natefinch/lumberjack.v2"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/logger"

	"devops/utils/comm"
)

var (
	logFile  = comm.AbsPath("log", "report.log")
	logLevel = log.LevelInfo
)

func LoggerInit(app *fiber.App) {

	// 设置 log 输出 level 和 输出文件
	log.SetLevel(logLevel)

	writer := io.MultiWriter(os.Stdout,
		&lumberjack.Logger{
		Filename:   logFile, // 日志文件的位置
		MaxSize:    1,       // 文件最大尺寸（以MB为单位）
		MaxBackups: 3,       // 保留的最大旧文件数量
		MaxAge:     28,      // 保留旧文件的最大天数
		Compress:   false,   // 是否压缩/归档旧文件
		LocalTime:  true,    // 使用本地时间创建时间戳
	})

	log.SetOutput(writer)

	// 设置 fiber 默认打印格式
	app.Use(logger.New(logger.Config{
		Next:          nil,
		Done:          nil,
		Format:        "[${time}] ${ip} ${status} - ${latency} ${method} ${path} ${error}\n",
		TimeFormat:    "2006/01/02 15:04:05",
		TimeZone:      "Asia/Shanghai",
		TimeInterval:  500 * time.Millisecond,
		Output:        writer,
		DisableColors: false,
	}))

	// Default
	// return logger.New(logger.Config{
	// 	Next:          nil,
	// 	Done:          nil,
	// 	Format:        "[${time}] ${ip} ${status} - ${latency} ${method} ${path} ${error}\n",
	// 	TimeFormat:    "15:04:05",
	// 	TimeZone:      "Local",
	// 	TimeInterval:  500 * time.Millisecond,
	// 	Output:        os.Stdout,
	// 	DisableColors: false,
	// })
}