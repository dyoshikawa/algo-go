package p73

func Invoke(n int, numbers []int) int {
	prds := make([]int, 0)
	for i := 0; i < n; i++ {
		newNumbers := make([]int, n)
		copy(newNumbers, numbers)
		newNumbers[i]++
		prd := 1
		for _, v := range newNumbers {
			prd *= v
		}
		prds = append(prds, prd)
	}
	max := 0
	for _, v := range prds {
		if v > max {
			max = v
		}
	}
	return max
}
