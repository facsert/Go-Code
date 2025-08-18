package str

import (
	"fmt"
	"strings"
)

// 判断字符串是否包含某个子串
func Contains(s, sub string) bool {
	return strings.Contains(s, sub)
}

// 字符串转小写
func Lower(s string) string {
	return strings.ToLower(s)
}

// 字符串转大写
func Upper(s string) string {
	return strings.ToUpper(s)
}

// 字符串反转
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 字符串分割
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// 字符串合并
func Join(a []string, sep string) string {
	return strings.Join(a, sep)
}

// 子串定位
func Index(s, sub string) int {
	return strings.Index(s, sub)
}

// 字符串替换
func Replace(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

// 字符串重复
func Repeat(s string, count int) string {
	return strings.Repeat(s, count)
}

// 字符串开头是否包含某个子串
func StartsWith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// 字符串结尾是否包含某个子串
func EndsWith(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// 去除字符串前后空格
func Trim(s string) string {
	return strings.TrimSpace(s)
}

// 自定义结构体打印内容
type Tool struct {
	Name string
    Size int
}

func NewTool(name string, size int) Tool {
	return Tool{
		Name: name,
		Size: size,
	}
}

func (t Tool) String() string {
	return fmt.Sprintf("A tool %s with %d size", t.Name, t.Size)
}

// fmt.Println(NewTool("car", 10))
// A tool car with 10 size