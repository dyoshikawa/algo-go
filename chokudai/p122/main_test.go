package p98

import "testing"

func TestInvoke0(t *testing.T) {
	n := 1
	east := 25
	west := 25
	south := 25
	north := 25
	if res := Invoke(n, east, west, south, north); res != 1.0 {
		t.Fatal(res)
	}
}

func TestInvoke1(t *testing.T) {
	n := 2
	east := 25
	west := 25
	south := 25
	north := 25
	if res := Invoke(n, east, west, south, north); res != 0.75 {
		t.Fatal(res)
	}
}

func TestInvoke2(t *testing.T) {
	n := 7
	east := 50
	west := 0
	south := 0
	north := 50
	if res := Invoke(n, east, west, south, north); res != 1.0 {
		t.Fatal(res)
	}
}

func TestInvoke3(t *testing.T) {
	n := 14
	east := 50
	west := 50
	south := 0
	north := 0
	if res := Invoke(n, east, west, south, north); res != 0.0001220703125 {
		t.Fatal(res)
	}
}
