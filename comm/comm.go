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
	"strings"
)

var (
	ROOT_PATH, GetPathError = os.Getwd()
)

func Init() {
    if GetPathError != nil { panic("Failed to get current path") }
	LoggerInit()
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

// 创建路径
func MakeDirs(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("create %s failed: %v", path, err)
		}
	} else if err != nil {
		return fmt.Errorf("check %s exist failed: %v", path, err)
	}
	return nil
}

// 标题打印
func Title(title string, level int) string {
	separator := [...]string{"#", "=", "*", "-"}[level % 4]
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
func ListDir(root string, abs bool) []string {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil { return err }
		if !info.IsDir() { 
			if abs {
				files = append(files, path)
			} else {
				files = append(files, filepath.Base(path))
			}
		}
		return nil
	})
	if err != nil { panic(fmt.Sprintf("Walk error: %v\n", err)) }
	return files
}
