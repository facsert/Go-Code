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
	// 可执行文件路径
	execPath, err := os.Executable()
	if err != nil {
		panic(fmt.Errorf("get file root path failed: %w", err))
	}
    
	// 文件路径不包含临时文件路径
	if !strings.Contains(execPath, os.TempDir()) {
		ROOT_PATH = filepath.Dir(execPath)
		return
	}
    
	// 源码路径
	_, runtimePath, _, ok := runtime.Caller(0)
	if !ok {
		panic(fmt.Errorf("get root path failed"))
	}	
	ROOT_PATH = filepath.Dir(runtimePath)

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

func Exists(dir string) error {
	if  _, err := os.Stat(dir); err == nil || os.IsNotExist(err) { 
		return nil 
	} else {
		return err
	}
}

// 获取指定路径下的所有文件func ListDir(root string, abs bool) ([]string, error) {
func ListDir(root string, abs bool) ([]string, error) {
	files := make([]string, 0, 10)
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
	if err != nil {
		return []string{}, fmt.Errorf("walk error: %w", err)
	}
	return files, nil
}