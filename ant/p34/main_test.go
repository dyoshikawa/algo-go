package p34

import "testing"

func TestCase1(t *testing.T) {
	n := 4
	a := []int{1, 2, 4, 7}
	k := 13
	if res := Invoke(n, a, k); res != "Yes" {
		t.Fatal(res)
	}
}
