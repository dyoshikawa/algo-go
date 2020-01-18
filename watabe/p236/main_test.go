package p236

import (
	"github.com/dyoshikawa/algo/util/collection/util_int"
	"testing"
)

func TestInvoke(t *testing.T) {
	n := 10
	as := []int{
		4, 1, 3, 2, 16, 9, 10, 14, 8, 7,
	}
	expected := []int{
		16, 14, 10, 8, 7, 9, 3, 2, 4, 1,
	}
	if res := Invoke(n, as); !util_int.EqualArr(res, expected) {
		t.Fatal(res)
	}
}
