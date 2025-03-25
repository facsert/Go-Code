package set


type Set[T comparable] struct {
	elems map[T]struct{}
}

func NewSet[T comparable](s ...T) *Set[T] {
	set := &Set[T]{elems: make(map[T]struct{}, len(s))}
	for _, elem := range s {
		set.Add(elem)
	}
	return set
}

// 添加数据
func (s *Set[T]) Add(elem T) {
	s.elems[elem] = struct{}{}
}

// 删除数据
func (s *Set[T]) Remove(elem T) {
	delete(s.elems, elem)
}

// 判断是否存在
func (s *Set[T]) Contains(elem T) bool {
	_, ok := s.elems[elem]
	return ok
}

// 元素数量
func (s *Set[T]) Size() int {
	return len(s.elems)
}

// 是否为空
func (s *Set[T]) IsEmpty() bool {
	return len(s.elems) == 0
}

// 清空
func (s *Set[T]) Clear() {
	s.elems = make(map[T]struct{})
}

// 转为切片
func (s *Set[T]) Values() []T {
	values := make([]T, 0, len(s.elems))
	for elem := range s.elems {
		values = append(values, elem)
	}
	return values
}

// range 遍历所有元素
func (s *Set[T]) All() func(yield func(T) bool) {
    return func(yield func(T) bool) {
		for elem := range s.elems {
			if !yield(elem) {
				return
			}
		}
	}
}


// 并集, set 合并
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	union := &Set[T]{elems: make(map[T]struct{}, len(s.elems) + len(other.elems))}
	for elem := range s.elems {
		union.Add(elem)
	}
	for elem := range other.elems {
		union.Add(elem)
	}
	return union
}

// 交集, 获取两个 set 共有的元素
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	intersection := &Set[T]{elems: make(map[T]struct{}, len(s.elems))}
	if len(s.elems) > len(other.elems) {
		s, other = other, s
	}
	for elem := range s.elems {
		if other.Contains(elem) {
			intersection.Add(elem)
		}
	}
	return intersection
}

// 差集, 获取 s 中不在 other 中的元素
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	difference := &Set[T]{elems: make(map[T]struct{}, len(s.elems))}
	for elem := range s.elems {
		if !other.Contains(elem) {
			difference.Add(elem)
		}
	}
	return difference
}