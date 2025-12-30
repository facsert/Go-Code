package dateTime

import (
	// "fmt"
	"log/slog"
	"time"

	"learn/comm"
)

func Test() {
	Now()
	Date()
	ParseDuration()
	Add()
	Before()
	After()
	Equal()
	Sub()
	Since()
	Parse()
	Format()
}

// 返回当前时间点
func Now() {
	comm.Title(3, "time.Now() -> time.Time")
	slog.Info("currentTime := time.Now()", slog.Time("currentTime", time.Now()))
}

// 创建时间点
func Date() {
	comm.Title(3, "time.Date(year, month, day, hour, minute, second, nanosecond, loc) -> time.Time")
	slog.Info("t := time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)", slog.Time("time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)", time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)))
}

// 解析时间片段
func ParseDuration() {
	comm.Title(3, `time.ParseDuration("8h1m1s") -> time.Duration, error`)
	d, _ := time.ParseDuration("8h1m1s")
	slog.Info(`duration := time.ParseDuration("8h1m1s")`, slog.Float64("parse 8h1m1s to second", d.Seconds()))
}

// 时间点加减时间片段
func Add() {
	comm.Title(3, "time.Time.Add(duration) -> time.Time")
	slog.Info("time.Now().Add(2 * time.Hour)", slog.Time("now + 2h", time.Now().Add(2*time.Hour)))
}

// 时间点比较
func After() {
	comm.Title(3, "time.Time.After(t) -> bool")
	future := time.Now().Add(2 * time.Hour)
	slog.Info("time.Now().Add(2 * time.Hour).After(time.Now())", slog.Bool("future > now", future.After(time.Now())))
}

// 时间点比较
func Before() {
	comm.Title(3, "time.Time.Before(t) -> bool")
	past := time.Now().Add(-2 * time.Hour)
	slog.Info("time.Now().Add(-2 * time.Hour).Before(time.Now())", slog.Bool("past < now", past.Before(time.Now())))
}

// 时间点比较
func Equal() {
	comm.Title(3, "time.Time.Equal(t) -> bool")
	now := time.Now()
	slog.Info("now.Equal(now)", slog.Bool("now == now", now.Equal(now)))
}

// 时间点差值
func Sub() {
	comm.Title(3, "time.Time.Sub(t) -> time.Duration")
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)
	slog.Info("end.Sub(start)", slog.Float64("(end - start)s", end.Sub(start).Seconds()))
}

// 计算时间点与当前时间的差值
func Since() {
	comm.Title(3, "time.Since(t) -> time.Duration")
	t := time.Now()
	time.Sleep(1 * time.Second)
	slog.Info("time.Since(t)", slog.Float64("(now - t)s", time.Since(t).Seconds()))
}

// 将字符串转化为时间点
func Parse() {
	comm.Title(3, "time.Parse(format, timeStr) -> time.Time, error")
	t, _ := time.Parse(time.DateTime, "2023-01-01 12:00:00")
	slog.Info(
		`time.Parse(time.DateTime, "2023-01-01 12:00:00")`,
		slog.Time("parse str to time,Time", t),
	)
	// time.DateTime    = "2006-01-02 15:04:05"
	// time.RFC3339     = "2006-01-02T15:04:05Z07:00"
	// time.RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
}

// 将时间点转化为字符串
func Format() {
	comm.Title(3, "time.Time.Format(format) -> string")
	slog.Info(
		`time.Time.Format(time.DateTime)`,
		slog.String("currentTime with format 2006-01-02 15:04:05", time.Now().Format(time.DateTime)),
	)
}
