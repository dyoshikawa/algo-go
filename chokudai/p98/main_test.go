package p98

import "testing"

func TestInvoke1(t *testing.T) {
	friends := []string{
		"NNN",
		"NNN",
		"NNN",
	}
	if res := Invoke(friends); res != 0 {
		t.Fatal(res)
	}
}

func TestInvoke2(t *testing.T) {
	friends := []string{
		"NYY",
		"YNY",
		"YYN",
	}
	if res := Invoke(friends); res != 2 {
		t.Fatal(res)
	}
}

func TestInvoke3(t *testing.T) {
	friends := []string{
		"NYNNN",
		"YNYNN",
		"NYNYN",
		"NNYNY",
		"NNNYN",
	}
	if res := Invoke(friends); res != 4 {
		t.Fatal(res)
	}
}

func TestInvoke4(t *testing.T) {
	friends := []string{
		"NNNNYNNNNN",
		"NNNNYNYYNN",
		"NNNYYYNNNN",
		"NNYNNNNNNN",
		"YYYNNNNNNN",
		"NNYNNNNNYN",
		"NYNNNNNYNN",
		"NYNNNNYNNN",
		"NNNNNYNNNN",
		"NNNNYNNNNN",
	}
	if res := Invoke(friends); res != 8 {
		t.Fatal(res)
	}
}
