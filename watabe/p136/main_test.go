package p136

import "testing"

func TestInvoke(t *testing.T) {
	n := 5
	k := 3
	T := []int{8, 1, 7, 3, 9}
	if res := Invoke(n, k, T); res != 10 {
		t.Fatal(res)
	}
}

func TestInvokeAlt(t *testing.T) {
	n := 5
	k := 3
	T := []int{8, 1, 7, 3, 9}
	if res := InvokeAlt(n, k, T); res != 10 {
		t.Fatal(res)
	}
}
