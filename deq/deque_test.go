package deq

import (
	"time"
	"math/rand"
	"reflect"
	"testing"
)


func TestDeq_PushPop(t *testing.T) {
	d := NewDeq[int](nil, 0)

	d.Push(1)
	d.Push(2)
	d.Push(3)

	if v := d.Pop(); v != 3 {
		t.Fatalf("expected 3, got %d", v)
	}
	if v := d.Pop(); v != 2 {
		t.Fatalf("expected 2, got %d", v)
	}
	if v := d.Pop(); v != 1 {
		t.Fatalf("expected 1, got %d", v)
	}
}

func TestDeq_PushFrontPopFront(t *testing.T) {
	d := NewDeq[int](nil, 0)

	d.PushFront(1)
	d.PushFront(2)
	d.PushFront(3)

	if v := d.PopFront(); v != 3 {
		t.Fatalf("expected 3, got %d", v)
	}
	if v := d.PopFront(); v != 2 {
		t.Fatalf("expected 2, got %d", v)
	}
	if v := d.PopFront(); v != 1 {
		t.Fatalf("expected 1, got %d", v)
	}
}

func TestDeq_Mixed(t *testing.T) {
	d := NewDeq[int](nil, 0)

	d.Push(1)
	d.Push(2)
	d.PushFront(0)
	d.Push(3)

	if v := d.PopFront(); v != 0 {
		t.Fatal(v)
	}
	if v := d.Pop(); v != 3 {
		t.Fatal(v)
	}
	if v := d.PopFront(); v != 1 {
		t.Fatal(v)
	}
	if v := d.Pop(); v != 2 {
		t.Fatal(v)
	}
}

func TestDeq_Empty(t *testing.T) {
	d := NewDeq[int](nil, 0)

	if !d.IsEmpty() {
		t.Fatal("should be empty")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()

	d.Pop()
}

func TestDeq_SingleElement(t *testing.T) {
	d := NewDeq[int](nil, 0)

	d.Push(42)

	v, ok := d.Front()
	if ok != nil || v != 42 {
		t.Fatal(v, ok)
	}

	v, ok = d.Back()
	if ok != nil || v != 42 {
		t.Fatal(v, ok)
	}

	if d.Pop() != 42 {
		t.Fatal("wrong pop")
	}

	if !d.IsEmpty() {
		t.Fatal("should be empty")
	}
}

func TestDeq_RingBehavior(t *testing.T) {
	d := NewDeq[int](nil, 4)

	// 填满并弹出部分，制造环
	for i := range 8 {
		d.Push(i)
	}
	for range 5 {
		d.PopFront()
	}

	// 再插入，触发环绕
	for i := 100; i < 105; i++ {
		d.Push(i)
	}

	// 校验顺序
	var result []int
	for v := range d.All() {
		result = append(result, v)
	}

	expected := []int{5,6,7,100,101,102,103,104}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestDeq_Grow(t *testing.T) {
	d := NewDeq[int](nil, 2)

	n := 1000
	for i := range n {
		d.Push(i)
	}

	for i := range n {
		if v := d.PopFront(); v != i {
			t.Fatalf("expected %d, got %d", i, v)
		}
	}
}

func TestDeq_All(t *testing.T) {
	d := NewDeq[int](nil, 0)

	for i := range 10 {
		d.Push(i)
	}

	var result []int
	for v := range d.All() {
		result = append(result, v)
	}

	if len(result) != 10 {
		t.Fatal("wrong length")
	}

	for i := range 10 {
		if result[i] != i {
			t.Fatalf("expected %d got %d", i, result[i])
		}
	}
}

func BenchmarkDeq_PushPop(b *testing.B) {
	d := NewDeq[int](nil, 0)

	for i := 0; b.Loop(); i++ {
		d.Push(i)
		d.Pop()
	}
}

func TestDeq_Random(t *testing.T) {
	d := NewDeq[int](nil, 0)
	var ref []int

	rand.Seed(time.Now().UnixNano())

	for range 10000 {
		op := rand.Intn(4)

		switch op {
		case 0: // Push
			v := rand.Intn(1000)
			d.Push(v)
			ref = append(ref, v)

		case 1: // PushFront
			v := rand.Intn(1000)
			d.PushFront(v)
			ref = append([]int{v}, ref...)

		case 2: // Pop
			if len(ref) == 0 {
				continue
			}
			v1 := d.Pop()
			v2 := ref[len(ref)-1]
			ref = ref[:len(ref)-1]

			if v1 != v2 {
				t.Fatalf("pop mismatch %d vs %d", v1, v2)
			}

		case 3: // PopFront
			if len(ref) == 0 {
				continue
			}
			v1 := d.PopFront()
			v2 := ref[0]
			ref = ref[1:]

			if v1 != v2 {
				t.Fatalf("popFront mismatch %d vs %d", v1, v2)
			}
		}
	}
}