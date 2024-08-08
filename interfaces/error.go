package interfaces

import (
	"fmt"
	"time"
)

// type error interface { Error() string }
// 类似于 Stringer 接口, 只要实现 Error() 方法, 就可以作为 error 类型

type TimeoutErr struct {
	Time time.Time
	Desc string
}

func (e *TimeoutErr) Error() string {
	return fmt.Sprintf("TimeoutErr: %s at %s", e.Desc, e.Time)
}

func TestErr() {
	err := func() error {
        fmt.Println("spend 3 hours")
		return &TimeoutErr{
			Time: time.Now(),
			Desc: "timeout",
		}
	}()
	fmt.Println(err)
	// TimeoutErr: timeout at 2024-05-15 21:16:32.778150669 +0800 CST m=+0.000048819
}