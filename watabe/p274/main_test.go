package p269

import (
	"github.com/dyoshikawa/algo/util"
	"testing"
)

var n = 6
var as = [][]int{
	{1, 2, 2, 3},
	{2, 2, 3, 4},
	{3, 1, 5},
	{4, 1, 6},
	{5, 1, 6},
	{6, 0},
}
var expected = [][]int{
	{1, 12},
	{2, 11},
	{3, 8},
	{9, 10},
	{4, 7},
	{5, 6},
}

func TestInvokeRecursion(t *testing.T) {
	if res := InvokeRecursion(n, as); !util.EqualDoubleArrInt(expected, res) {
		t.Fatal(res)
	}
}
