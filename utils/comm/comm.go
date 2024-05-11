/*
 * @Author: facsert
 * @Date: 2023-08-06 22:18:09
 * @LastEditTime: 2023-08-06 22:23:27
 * @LastEditors: facsert
 * @Description:
 */

package comm

import (
	"fmt"
	"os"
	"log"
	"log/slog"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	ROOT_PATH string
)

func init() {
    _, file, _, ok := runtime.Caller(0)
	if !ok { panic("Failed to get caller info") }
    ROOT_PATH = filepath.Dir(filepath.Dir(filepath.Dir(file)))
}

// 基于根目录的绝对路径
func AbsPath(elems ...string) string {
	for i := (len(elems)-1); i > 0; i-- {
		if filepath.IsAbs(elems[i]) {
			elems = elems[i:]
			break
		}
	}
	path := filepath.Join(elems...)
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(ROOT_PATH, path)
}

// 标题打印
func Title(title string, level int) string {
	separator := [...]string{"#", "=", "*", "-"}[level % 4]
	// space := [...]string{"\n\n", "\n", "", ""}[level % 4]
	line := strings.Repeat(separator, 50)
	log.Printf("%s %s %s", line, title, line)
	return title
}

// 阶段性结果打印
func Display(msg string, success bool) string {
	length, chars := 80, ""
	if success {
		slog.Info(fmt.Sprintf("%-" + fmt.Sprintf("%d", length) + "s  [PASS]\n", msg))
		return msg
	}
	if len(msg) > length {
		chars = strings.Repeat(">", length - len(msg))
	}
	fmt.Println(msg + " " + chars + " [FAIL]")
	return msg
}


// 获取指定路径下的所有文件
func ListDir(root string) []string {
	var files []string
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil { return err }
		if !info.IsDir() { files = append(files, path) }
		return nil
	})
	return files
}
