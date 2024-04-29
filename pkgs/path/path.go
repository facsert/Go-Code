package path

import (
    "fmt"
	"os"
	"log/slog"
    "path/filepath"
)

//绝对路径
func Abs(path string) (string, error) {
    return filepath.Abs(path)
}

// 是否绝对路径
func IsAbs(path string) bool {
    return filepath.IsAbs(path)
}

// 获取除去最后一个元素的路径
func Dir(path string) string {
    return filepath.Dir(path)
}

// 获取路径最后一个元素
func Base(path string) string {
    return filepath.Base(path)
}

// 路径拼接
func Join(elem...string) string {
    return filepath.Join(elem...)
}

// 遍历路径下的文件
func Walk(root string, walkFn filepath.WalkFunc) error {
    return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
		if info.IsDir() {
			slog.Info(fmt.Sprintf("path: %v\n", path))
			return nil
		} else {
			slog.Info(fmt.Sprintf("file: %v\n", path))
			return nil
		}
	})
}

// 切换路径
func Chdir(path string) error {
    return os.Chdir(path)
}

// 获取当前路径
func Getwd() (string, error) {
    return os.Getwd()
}

// 创建路径
func Mkdir(path string) error {
    return os.Mkdir(path, 0755)
}

// 创建多层路径
func MkdirAll(path string) error {
    return os.MkdirAll(path, 0755)
}

// 删除文件或空目录
func Remove(path string) error {
    return os.Remove(path)
}

// 递归删除路径下所有文件
func RemoveAll(path string) error {
    return os.RemoveAll(path)
}

// 重命名
func Rename(old, new string) error {
    return os.Rename(old, new)
}

// 查看路径下的所有文件
func ReadDir(path string) error {
    var files []os.DirEntry
	var err error
    files, err = os.ReadDir(path)
    if err!= nil {
        return err
    }
	for _, file := range files {
		slog.Info(fmt.Sprintf("file: %v\n", file.Name()))
	}
	return nil
}

// 创建文件
func Create(name string) (*os.File, error) {
    file, err := os.Create(name)
	slog.Info(fmt.Sprintf("Create file %s\n", file.Name()))
	return file, err
}

// 判断文件或路径存在
func Exists(path string) bool {
    _, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}


// 文件读写 
// O_RDWR(读写) O_CREATE(不存在创建) O_APPEND(追加) O_RDONLY(只读) O_WRONLY(只写)
// Read 读取 len(content) 长度内容赋值给 content
func ReadWrite(name string) (*os.File, error) {
    file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	file.Write([]byte("write something to file"))
	content := make([]byte, 1024)
	file.Read(content)
	return file, err
}

// 获取所有环境变量
func Envs() []string {
    return os.Environ()
}

// 获取环境变量值
func Env(key string) string {
    return os.Getenv(key)
}

// 设置环境变量
func Setenv(key, value string) error {
    return os.Setenv(key, value)
}

// 退出程序
func Exit(code int) {
    os.Exit(code)
}