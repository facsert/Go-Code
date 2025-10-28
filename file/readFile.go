/*
 * @Author: facsert
 * @Date: 2023-08-06 17:23:38
 * @LastEditTime: 2023-08-06 22:00:56
 * @LastEditors: facsert
 * @Description:
 */
package file

import (
	"bufio"
	"fmt"
	"os"
)

type File struct {
	FileName string
}

// 读取文件全部内容
func (f *File) Read() (string, error) {

	content, err := os.ReadFile(f.FileName)
	if err != nil {
		return "", fmt.Errorf("read file error: %w", err)
	}
	return string(content), nil
}

// 逐行读取文件
func (f *File) ReadLine() ([]string, error) {
	fs, err := os.OpenFile(f.FileName, os.O_RDONLY, 0666)
	if err != nil {
        return []string{}, fmt.Errorf("open file error: %w", err)
	}
	defer fs.Close()

	lines := make([]string, 0, 100)
	scanner := bufio.NewScanner(fs)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return []string{}, fmt.Errorf("scan error: %w", err)
	}
	return lines, nil
}

// 覆盖写入
func (f *File) Cover(s string) {
	if err := os.WriteFile(f.FileName, []byte(s), 0666); err != nil {
		panic(fmt.Sprintf("error: %v\n", err))
	}
}

// 追加写入
func (f *File) Append(s string) {
	file, err := os.OpenFile(f.FileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(fmt.Sprintf("error: %v\n", err))
	}
	defer file.Close()

	if _, err := file.Write([]byte(s)); err != nil {
		panic(fmt.Sprintf("write file failed: %v\n", err))
	}
}
