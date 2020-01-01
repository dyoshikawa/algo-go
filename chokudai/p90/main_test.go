package p90

import "testing"

func TestInvoke1(t *testing.T) {
	if res := Invoke("abab"); res != 5 {
		t.Fatal(res)
	}
}

func TestInvoke2(t *testing.T) {
	if res := Invoke("abacaba"); res != 7 {
		t.Fatal(res)
	}
}

func TestInvoke3(t *testing.T) {
	if res := Invoke("qwerty"); res != 11 {
		t.Fatal(res)
	}
}

func TestInvoke4(t *testing.T) {
	if res := Invoke("abdfhdyrbdbsdfghjk11kjhgfds"); res != 38 {
		t.Fatal(res)
	}
}
