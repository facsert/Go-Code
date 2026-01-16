package comm

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

var ROOT_PATH string

func Init() {
	defer func() {
		if ROOT_PATH == "" {
			log.Fatalln("get root path failed")
		}
		if err := NewLogger(AbsPath("log", "report.log"), 50, 3); err != nil {
			log.Fatalf("initialize logger failed: %v", err)
		}
	}()

	// build 可执行文件
	execPath, err := os.Executable()
	if err == nil {
		if !strings.Contains(execPath, "go-build") || !strings.Contains(execPath, "b001"){
			ROOT_PATH = filepath.Dir(execPath)
			return
		}
	}

	// go run 源码路径
	_, runtimePath, _, ok := runtime.Caller(0)
	if ok && strings.HasSuffix(runtimePath, ".go") {
		ROOT_PATH = filepath.Dir(filepath.Dir(runtimePath))
		return
	}
}

// 基于根目录的绝对路径
func AbsPath(elems ...string) string {
	for i := (len(elems) - 1); i > 0; i-- {
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
func Title(level int, title string, a ...any) string {
	title = fmt.Sprintf(title, a)
	separator := [...]string{"#", "=", "*", "-"}[level%4]
	line := strings.Repeat(separator, max((100-len(title))/2, 0))
	log.Printf("%s %s %s", line, title, line)
	return title
}

// 阶段性结果打印
func Display(success bool, msg string, a ...any) string {
	msg = fmt.Sprintf(msg, a)
	length, chars := 80, ""
	if success {
		slog.Info(fmt.Sprintf("%-"+strconv.Itoa(length)+"s  [SUCCESS]", msg))
		return msg
	}
	if len(msg) > length {
		chars = strings.Repeat(">", length-len(msg))
	}
	slog.Info(msg + " " + chars + " [FAILED]")
	return msg
}

// 判断文件或路径是否存在
func Exists(dir string) bool {
	_, err := os.Stat(dir)
	return err == nil
}

// 获取路径下所有文件
// dst 目标路径
// abs 是否返回绝对路径
func ListDir(dst string, abs bool) ([]string, error) {
	files := make([]string, 0, 10)
	err := filepath.Walk(dst, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
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
