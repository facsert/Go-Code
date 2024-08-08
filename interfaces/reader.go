package interfaces

import (
	
	"fmt"
	_ "fmt"
	"io"
)

// io 包 type Reader interface { Read(p []byte) (n int, err error) }
// 通过数据源创建 Reader 对象, 使用 Read 方法从数据源读取内容写入 p 切片中
// 每次读取长度为 p 的长度, 返回读取的字节数和错误信息
// 当数据流结束, 返回 io.EOF 错误

type StringReader struct {
	index   int
    content string
}

func (s *StringReader) Read(p []byte) (n int, err error) {
	if s.index >= len(s.content) {
		return 0, io.EOF
	}
	n = copy(p, s.content[s.index:])
	s.index += n
    return n, nil
}

func TestReader() {
    reader := &StringReader{index: 0,content: "123456789"}
	b := make([]byte, 3)
	for {
		_, err := reader.Read(b)
		if err == io.EOF {
			fmt.Println("read over")
			break
		}
		fmt.Println(string(b))
	}
}

// 123
// 456
// 789
// read over