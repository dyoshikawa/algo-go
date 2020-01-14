package p269

import (
	"github.com/dyoshikawa/algo/util"
	"testing"
)

var n = 4
var as = [][]int{
	{1, 2, 2, 4},
	{2, 1, 4},
	{3, 0},
	{4, 1, 3},
}
var expected = [][]int{
	{1, 0},
	{2, 1},
	{3, 2},
	{4, 1},
}

func TestInvoke(t *testing.T) {
	if res := Invoke(n, as); !util.EqualDoubleArrInt(expected, res) {
		t.Fatal(res)
	}
}
