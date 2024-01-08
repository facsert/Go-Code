package re

import (
	"log/slog"
	"regexp"
)

// 正则表达式查找字符串(第一个匹配)
func Search(regex, s string) []byte {
    re := regexp.MustCompile(regex)
	sub := re.Find([]byte(s))
	slog.Info(string(sub))
	return sub
}

// 正则表达式查找字符串(所有匹配)
func SearchAll(regex, s string) [][]byte {
    re := regexp.MustCompile(regex)
	sub := re.FindAll([]byte(s), -1)
    for _, v := range sub {
		slog.Info(string(v))
	}
	return sub
}

// 正则表达式是否在字符串中
func Contain(regex, s string) bool {
    re := regexp.MustCompile(regex)
	return re.MatchString(s)
}

// 正则表达式替换
func Replace(regex, s, r string) string {
    re := regexp.MustCompile(regex)
	return re.ReplaceAllString(s, r)
}