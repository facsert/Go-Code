package deq

import (
	"fmt"
	"iter"
	"strings"
)


var MIN_SIZE = 2 << 4

// 环形数组构建双端队列(双指针)
// 双端操作本质是读取或写入 head 和 tail 索引的值
// prev 和 next 计算 head 和 tail 偏移
// 切片空间占满自动增长长度
type Deq[T any] struct {
	queue []T  // 存储数值切片
	head  int  // 指向第一个有效值
	tail  int  // 指向下一个可插入位置
	count int  // 当前有效值数量 queue[head:tail]
}

// 创建双端队列
// size 为 2 的倍数便于后续位操作
func NewDeq[T any](queue []T, size int) *Deq[T] {
	curr := max(len(queue), size)
	s := MIN_SIZE
	for s < curr {
		s <<= 1
	}
	deq := make([]T, s)
	copy(deq, queue)

	return &Deq[T]{
		queue: deq,
		head:  0,
		tail:  len(queue),
		count: len(queue),
	}
}

func (d *Deq[T])Len() int {
	if d.queue == nil {
        return 0
	}
	return d.count
}

func (d *Deq[T])Cap() int {
	if d.queue == nil {
        return 0
	}
    return len(d.queue)
}

func (d *Deq[T]) IsEmpty() bool {
	return d.count == 0
}

func (d *Deq[T]) Push(elem T) {
	d.grow()
	d.queue[d.tail] = elem
	d.tail = d.next(d.tail)
	d.count++
}

func (d *Deq[T]) PushFront(elem T) {
	d.grow()
	d.head = d.prev(d.head)
	d.queue[d.head] = elem
	d.count++
}

func (d *Deq[T]) Pop() T {
	if d.count == 0 {
		panic("pop with empty queue")
	}

	d.tail = d.prev(d.tail)
	elem := d.queue[d.tail]
	d.count--
	return elem
}

func (d *Deq[T]) PopFront() T {
	if d.count == 0 {
		panic("pop with empty queue")
	}
	elem := d.queue[d.head]
	d.head = d.next(d.head)
	d.count--
	return elem
}

func (d *Deq[T]) Front() (T, error) {
	if d.count == 0 {
		var zero T
		return zero, fmt.Errorf("empty queue")
	}
	return d.queue[d.head], nil
}

func (d *Deq[T]) Back() (T, error) {
	if d.count == 0 {
		var zero T
		return zero, fmt.Errorf("empty queue")
	}
	return d.queue[d.prev(d.tail)], nil
}

// 提供队列遍历功能
func (d *Deq[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		if d.count == 0 {
			return
		}
		for i := 0; i < d.count; i++ {
			pos := (d.head + i) & (len(d.queue) - 1)
			if !yield(d.queue[pos]) {
				break
			}
		}
	}
}

// 获取队列下一个数据在切片中的索引
// 2^n - 1 二进制除首位, 其余为 1, 可用于取余
// x % 2^n == x & (2^n - 1)
func (d *Deq[T]) next(i int) int {
	return (i + 1) & (len(d.queue) - 1)
}

func (d *Deq[T]) prev(i int) int {
	return (i - 1) & (len(d.queue) - 1)
}

// fmt 包打印
func (d *Deq[T]) String() string {
	if d.count == 0 {
		return "[]"
	}

	var s strings.Builder
	s.WriteByte('[')
	for i := 0; i < d.count; i++{
		if i > 0 {
            s.WriteByte(' ')
		}
		pos := (d.head + i) & (len(d.queue) - 1)
		fmt.Fprint(&s, d.queue[pos])
	}
	s.WriteByte(']')
	return s.String()
}

// 当有效数据数量达到切片数组大小
// 数组自动扩充为原有一倍
// 重新组织队列顺序
func (d *Deq[T]) grow() {
	if d.count < len(d.queue) {
		return
	}
	q := make([]T, len(d.queue)<<1)
	if d.tail > d.head {
		copy(q, d.queue[d.head:d.tail])
	} else {
		n := copy(q, d.queue[d.head:])
		copy(q[n:], d.queue[:d.tail])
	}
	d.head = 0
	d.tail = d.count
	d.queue = q
}
