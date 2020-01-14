package p287

import (
	"github.com/dyoshikawa/algo/util/collection/util_bool"
	"testing"
)

func TestInvoke(t *testing.T) {
	n = 10
	m = 9
	q = 3
	fs = [][]int{
		{0, 1},
		{0, 2},
		{3, 4},
		{5, 7},
		{5, 6},
		{6, 7},
		{6, 8},
		{7, 8},
		{8, 9},
	}
	qs = [][]int{
		{0, 1},
		{5, 9},
		{1, 3},
	}
	expected := []bool{
		true,
		true,
		false,
	}
	if res := Invoke(); !util_bool.EqualArr(res, expected) {
		t.Fatal(res)
	}
}
