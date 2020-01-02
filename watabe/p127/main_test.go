package p127

import "testing"

func equalArr(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestInvoke(t *testing.T) {
	n := 6
	strs := []string{
		"insert AAA",
		"insert AAC",
		"find AAA",
		"find CCC",
		"insert CCC",
		"find CCC",
	}
	if res := Invoke(n, strs); !equalArr(res, []string{"yes", "no", "yes"}) {
		t.Fatal(res)
	}
}
