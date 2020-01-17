package p234

import (
	"github.com/dyoshikawa/algo/util/collection/util_string"
	"testing"
)

func TestInvoke(t *testing.T) {
	n := 5
	as := []int{
		7,
		8,
		1,
		2,
		3,
	}
	expected := []string{
		"node 1: key = 7, left key = 8, right key = 1, ",
		"node 2: key = 8, parent key = 7, left key = 2, right key = 3, ",
		"node 3: key = 1, parent key = 7, ",
		"node 4: key = 2, parent key = 8, ",
		"node 5: key = 3, parent key = 8, ",
	}
	if res := Invoke(n, as); !util_string.EqualArr(res, expected) {
		t.Fatal(res)
	}
}
