package p240

import (
	"github.com/dyoshikawa/algo/util/collection/priority_queue_int"
	"github.com/dyoshikawa/algo/util/collection/util_int"
	"strconv"
	"strings"
)

var pq []int

func parent(i int) int {
	increasedI := i + 1
	isLeft := increasedI%2 == 0
	var increasedParentI int
	if isLeft {
		increasedParentI = increasedI / 2
	} else {
		increasedParentI = (increasedI - 1) / 2
	}
	return increasedParentI - 1
}

func heapIncreaseKey(i int, key int) {
	if key < pq[i] {
		panic("ERROR: 新しいキーは現在のキーより小さい")
	}
	pq[i] = key
	for i > 0 && pq[parent(i)] < pq[i] {
		util_int.SwapArr(pq, i, parent(i))
		i = parent(i)
	}
}

func insert(key int) {
	pq = append(pq, 0)
	heapIncreaseKey(len(pq)-1, key)
}

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

func extract() int {
	max := pq[0]
	lastI := len(pq) - 1
	pq[0] = pq[lastI]
	util_int.UnsetArr(pq, lastI)
	maxHeapify(pq, 0)
	return max
}

func Invoke(as []string) []int {
	pq = make([]int, 0)
	res := make([]int, 0)
	for _, s := range as {
		ss := strings.Split(s, " ")
		cmd := ss[0]
		switch cmd {
		case "insert":
			key, _ := strconv.Atoi(ss[1])
			insert(key)
		case "extract":
			res = append(res, extract())
		case "end":
			break
		}
	}
	return res
}

func InvokeAlt(as []string) []int {
	pq := priority_queue_int.NewPriorityQueue()
	res := make([]int, 0)
	for _, s := range as {
		ss := strings.Split(s, " ")
		cmd := ss[0]
		switch cmd {
		case "insert":
			key, _ := strconv.Atoi(ss[1])
			pq.Push(key)
		case "extract":
			res = append(res, pq.Pop())
		case "end":
			break
		}
	}
	return res
}
