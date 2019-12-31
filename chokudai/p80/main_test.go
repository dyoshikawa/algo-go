package p80

import "testing"

func equalArr(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestInvoke1(t *testing.T) {
	if res := Invoke(10); !equalArr(res, []int{3, 9}) {
		t.Fatal(res)
	}
}

func TestInvoke2(t *testing.T) {
	if res := Invoke(3); !equalArr(res, []int{2}) {
		t.Fatal(res)
	}
}

func TestInvoke3(t *testing.T) {
	if res := Invoke(9); !equalArr(res, []int{2, 4, 8}) {
		t.Fatal(res)
	}
}

func TestInvoke4(t *testing.T) {
	if res := Invoke(26); !equalArr(res, []int{5, 25}) {
		t.Fatal(res)
	}
}

func TestInvoke5(t *testing.T) {
	if res := Invoke(30); !equalArr(res, []int{29}) {
		t.Fatal(res)
	}
}
