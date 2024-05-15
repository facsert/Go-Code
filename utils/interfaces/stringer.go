package interfaces

import (
	"fmt"
)

func Test() {
	fmt.Println(Node{"localhost", 8080}) // Node {localhost 8080}
}

// 定义 fmt 包的 Stringer 接口
// type Stringer interface { String() string }
// fmt.Println() 打印变量时, 会调用变量的 String() 方法

type Node struct {
	Host string
	Port int
}

func (n Node) String() string {
	return fmt.Sprintf("Node %s:%d", n.Host, n.Port)
}
