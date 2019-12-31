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
