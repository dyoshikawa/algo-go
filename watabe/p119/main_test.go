package p119

import "testing"

func TestInvoke(t *testing.T) {
	n := 5
	S := []int{1, 2, 3, 4, 5}
	q := 3
	T := []int{3, 4, 1}
	if res := Invoke(n, S, q, T); res != 3 {
		t.Fatal(res)
	}
}
