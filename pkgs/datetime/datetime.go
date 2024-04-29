package datetime

import (
	"fmt"
	"time"
)

// 返回当前时间点
func Now() time.Time {
	return time.Now()
}


func Duration[T int8 | int32| int64](hours, minutes, seconds T) time.Duration {
	d, err := time.ParseDuration(fmt.Sprintf("%vh%vm%vs", hours, minutes, seconds))
	if err != nil { panic(fmt.Sprintf("Parse %vh%vm%vs failed", hours, minutes, seconds)) }
	return d
}


// time.Time: 时间点  time.Duration: 时间片段
func TimeParse(format, timeStr string) (time.Time, error) {
	return time.Parse(format, timeStr)
}

// 按指定格式将时间点转化为字符串
func TimeFormat(t time.Time, format string) string {
	return t.Format(format)
}