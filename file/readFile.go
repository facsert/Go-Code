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
func (f *File) Read() string {
	content, err := os.ReadFile(f.FileName)
	if err != nil {
		panic(fmt.Sprintf("error: %s\n", err))
	}
	return string(content)
}

// 逐行读取文件
func (f *File) ReadLine() []string {
	file, err := os.OpenFile(f.FileName, os.O_RDONLY, 0666)
	if err != nil {
		panic(fmt.Sprintf("error: %v\n", err))
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
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
