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

var ROOT_PATH string

func Init() {
	defer func() {
        if ROOT_PATH == "" {
			panic(fmt.Errorf("get root path failed"))
		}
		LoggerInit()
	}()
    // build 可执行文件
    if execPath, err := os.Executable(); err == nil {
		if !strings.Contains(execPath, os.TempDir()) {
			ROOT_PATH = filepath.Dir(execPath)
			return
		}
	}
	// go run 源码路径
	if _, runtimePath, _, ok := runtime.Caller(0); ok {
        ROOT_PATH = filepath.Dir(filepath.Dir(filepath.Dir(runtimePath)))
		return
	}
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
	fileInfo, err := os.Stat(path)
    if os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}
	if err != nil {
		return err
	}
	if fileInfo.IsDir() {
		return fmt.Errorf("%s is not dir", path)
	}
	return nil
}

// 标题打印
func Title(title string, level int) string {
	separator := [...]string{"#", "=", "*", "-"}[level % 4]
	line := strings.Repeat(separator, 100)
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

// 判断文件或路径是否存在
func Exists(dir string) bool {
	if  _, err := os.Stat(dir); err == nil || os.IsExist(err) { 
		return true
	}
	return false
}

// 获取路径下所有文件
// dst 目标路径
// abs 是否返回绝对路径
func ListDir(dst string, abs bool) ([]string, error) {
	files := make([]string, 0, 10)
	err := filepath.Walk(dst, func(path string, info os.FileInfo, err error) error {
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
	if err != nil {
		return []string{}, fmt.Errorf("walk error: %w", err)
	}
	return files, nil
}