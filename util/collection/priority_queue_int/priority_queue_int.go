package priority_queue_int

import "github.com/dyoshikawa/algo/util/collection/util_int"

type PriorityQueue struct {
	Data []int
}

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

func (pq *PriorityQueue) Push(key int) {
	pq.Data = append(pq.Data, key)
	i := len(pq.Data)
	for i > 0 && pq.Data[parent(i)] < pq.Data[i] {
		util_int.SwapArr(pq.Data, i, parent(i))
		i = parent(i)
	}
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

func (pq *PriorityQueue) Pop() int {
	max := pq.Data[0]
	lastI := len(pq.Data) - 1
	pq.Data[0] = pq.Data[lastI]
	util_int.UnsetArr(pq.Data, lastI)
	maxHeapify(pq.Data, 0)
	return max
}
