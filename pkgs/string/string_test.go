package str

import (
	"testing"
	"reflect"
)

func TestContain(t *testing.T) {
    got := Contains("abc", "")
	if got!= true {
		t.Errorf("Contains() = %v, want %v", got, true)
	} else {
		t.Log("Contains() = ", got)
	}
}

func TestLower(t *testing.T) {
    got := Lower("ABC")
	if got != "abc" {
		t.Errorf("Lower() = %v, want %v", got, "abc")
	} else {
		t.Log("Lower() = ", got)
	}
}

func TestUpper(t *testing.T) {
    got := Upper("abc")
	if got!= "ABC" {
		t.Errorf("Upper() = %v, want %v", got, "ABC")
	} else {
		t.Log("Upper() = ", got)
	}
}

func TestReverse(t *testing.T) {
    got := Reverse("abc")
	if got!= "cba" {
		t.Errorf("Reverse() = %v, want %v", got, "cba")
	} else {
		t.Log("Reverse() = ", got)
	}
}

func TestSplit(t *testing.T) {
    got := Split("a,b,c", ",")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Split() = %v, want %v", got, want)
	} else {
		t.Log("Split() = ", got)
	}
}

func TestJoin(t *testing.T) {
    got := Join([]string{"a", "b", "c"}, "-")
	if got!= "a-b-c" {
		t.Errorf("Join() = %v, want %v", got, "a,b,c")
	} else {
		t.Log("Join() = ", got)
	}
}

func TestIndex(t *testing.T) {
    got := Index("abc", "b")
	if got!= 1 {
		t.Errorf("Index() = %v, want %v", got, 1)
	} else {
		t.Log("Index() = ", got)
	}
}

func TestReplace(t *testing.T) {
    got := Replace("abc", "b", "d", 2)
	if got!= "adc" {
		t.Errorf("Replace() = %v, want %v", got, "adc")
	} else {
		t.Log("Replace() = ", got)
	}
}

func TestRepeat(t *testing.T) {
    got := Repeat("abc", 3)
	if got!= "abcabcabc" {
		t.Errorf("Repeat() = %v, want %v", got, "abcabcabc")
	} else {
		t.Log("Repeat() = ", got)
	}
}

func TestStartsWith(t *testing.T) {
    got := StartsWith("abc", "a")
	if got!= true {
		t.Errorf("StartsWith() = %v, want %v", got, true)
	} else {
		t.Log("StartsWith() = ", got)
	}
}

func TestEndsWith(t *testing.T) {
    got := EndsWith("abc", "c")
	if got!= true {
		t.Errorf("EndsWith() = %v, want %v", got, true)
	} else {
		t.Log("EndsWith() = ", got)
	}
}
