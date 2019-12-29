package p73

import "testing"

func TestCase1(t *testing.T) {
	if res := Invoke(3, []int{1, 2, 3}); res != 12 {
		t.Fatal(res)
	}
}

func TestCase2(t *testing.T) {
	if res := Invoke(6, []int{1, 3, 2, 1, 1, 3}); res != 36 {
		t.Fatal(res)
	}
}

func TestCase3(t *testing.T) {
	if res := Invoke(6, []int{1000, 999, 998, 997, 996, 995}); res != 986074810223904000 {
		t.Fatal(res)
	}
}

func TestCase4(t *testing.T) {
	if res := Invoke(4, []int{1, 1, 1, 1}); res != 2 {
		t.Fatal(res)
	}
}
