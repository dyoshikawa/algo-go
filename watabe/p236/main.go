package p236

import "github.com/dyoshikawa/algo/util/collection/util_int"

func maxHeapify(as []int, i int) {
	increasedI := i + 1
	l := increasedI*2 - 1
	r := increasedI*2 + 1 - 1
	h := len(as)
	largestI := i
	if l < h && as[l] > as[i] {
		largestI = l
	}
	if r < h && as[r] > as[largestI] {
		largestI = r
	}
	if largestI != i {
		util_int.SwapArr(as, i, largestI)
		maxHeapify(as, largestI)
	}
}

func buildMaxHeap(as []int) []int {
	n := len(as)
	clonedAs := make([]int, n)
	copy(clonedAs, as)
	for i := n/2 - 1; i >= 0; i-- {
		maxHeapify(clonedAs, i)
	}
	return clonedAs
}

func Invoke(_ int, as []int) []int {
	return buildMaxHeap(as)
}
